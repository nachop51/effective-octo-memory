import Elysia from 'elysia'
import { UnauthorizedError } from './errors'
import jwt, { JWTPayloadSpec } from '@elysiajs/jwt'

type AuthToken =
  | (JWTPayloadSpec & {
      user: {
        id: string
        email: string
      }
    })
  | false

export const AuthMiddleware = new Elysia({ name: 'Auth.Middleware' })
  .use(
    jwt({
      name: 'jwt',
      secret: Bun.env.JWT_SECRET || 'supersecret',
    })
  )
  .derive(
    { as: 'global' },
    async ({ headers: { 'api-key': apiKey }, cookie: { auth }, jwt }) => {
      if (!auth.value) {
        throw new UnauthorizedError('Authentication cookie is missing')
      }

      const token = (await jwt.verify(auth.value as string)) as AuthToken

      if (!token) {
        throw new UnauthorizedError('Invalid authentication token')
      }

      const { user } = token

      console.log(user)

      return {
        token: {
          user,
        },
      }
    }
  )

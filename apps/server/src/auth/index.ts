import jwt from '@elysiajs/jwt'
import { User } from 'db'
import Elysia, { t } from 'elysia'
import { BadRequestError } from '../common/errors'
import { ConflictError, ForbiddenError, UnauthorizedError } from './errors'
import { AuthMiddleware } from './middleware'
import { AuthService } from './service'
import { PASSWORD_MIN_LENGTH } from '../utils/consts'

const createTokenPayload = (user: Pick<User, 'id' | 'email'>) => ({
  user: {
    id: user.id,
    email: user.email,
  },
  exp: Math.floor(Date.now() / 1000) + 30 * 24 * 60 * 60, // 30 days
  sub: user.id,
})

const createCookieOptions = (token: string) => ({
  value: token,
  expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
  httpOnly: true,
  path: '/',
  secure: Bun.env.NODE_ENV === 'production',
  sameSite: 'lax' as const,
})

export const auth = new Elysia({ prefix: '/auth' })
  .use(
    jwt({
      name: 'jwt',
      secret: Bun.env.JWT_SECRET || 'supersecret',
    })
  )
  .error({
    UnauthorizedError,
    ForbiddenError,
    ConflictError,
  })
  .onError(({ code, set }) => {
    switch (code) {
      case 'UnauthorizedError':
        set.status = 401
        return { message: 'Unauthorized' }
      case 'ConflictError':
        set.status = 409
        return { message: 'Conflict' }
      case 'ForbiddenError':
        set.status = 403
        return { message: 'Forbidden' }
    }
  })
  .post(
    '/signin',
    async ({ body, jwt, cookie: { auth } }) => {
      const user = await AuthService.signIn(body)

      const token = await jwt.sign(createTokenPayload(user))

      auth.set(createCookieOptions(token))

      return { token }
    },
    {
      body: t.Object({
        email: t.String({ format: 'email' }),
        password: t.String({
          minLength: PASSWORD_MIN_LENGTH,
        }),
      }),
    }
  )
  .post(
    '/signup',
    async ({ body, jwt, cookie: { auth } }) => {
      if (body.password !== body.confirmPassword) {
        throw new BadRequestError('Passwords do not match')
      }

      const user = await AuthService.signUp(body)

      const token = await jwt.sign(createTokenPayload(user))

      auth.set(createCookieOptions(token))

      return { token }
    },
    {
      body: t.Object({
        email: t.String({ format: 'email' }),
        password: t.String({
          minLength: PASSWORD_MIN_LENGTH,
        }),
        confirmPassword: t.String({
          minLength: PASSWORD_MIN_LENGTH,
        }),
      }),
    }
  )
  .use(AuthMiddleware)
  .get('/check', ({ token }) => token.user)
  .delete('/signout', ({ cookie: { auth } }) => {
    auth.set({
      value: '',
      expires: new Date(0),
      httpOnly: true,
      path: '/',
      secure: Bun.env.NODE_ENV === 'production',
      sameSite: 'lax',
    })

    return { message: 'Signed out successfully' }
  })

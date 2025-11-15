import Elysia, { t } from 'elysia'
import { AuthService } from './service'
import { ConflictError, ForbiddenError, UnauthorizedError } from './errors'

export const auth = new Elysia({ prefix: '/auth' })
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
  .get('/check', 'hello!')
  .post(
    '/signin',
    async ({ body }) => {
      console.log({ body })

      const user = await AuthService.signIn(body)

      console.log({ user })

      return user
    },
    {
      body: t.Object({
        email: t.String({ format: 'email' }),
        password: t.String(),
      }),
    }
  )
  .post(
    '/signup',
    async ({ body }) => {
      const user = await AuthService.signUp(body)

      return user
    },
    {
      body: t.Object({
        email: t.String({ format: 'email' }),
        password: t.String(),
      }),
    }
  )

import { Elysia, InternalServerError, NotFoundError } from 'elysia'
import { auth } from './auth'
import cors from '@elysiajs/cors'
import openapi from '@elysiajs/openapi'
import jwt from '@elysiajs/jwt'
import { BadRequestError } from './common/errors'

const app = new Elysia()
  .use(cors())
  .use(
    jwt({
      name: 'jwt',
      secret: Bun.env.JWT_SECRET || 'supersecret',
    })
  )
  .use(openapi())
  .error({
    NotFoundError,
    BadRequestError,
    InternalServerError,
  })
  .onError(({ code, set }) => {
    switch (code) {
      case 'NotFoundError':
        set.status = 404
        return { message: 'Not Found' }
      case 'BadRequestError':
        set.status = 400
        return { message: 'Bad Request' }
      case 'InternalServerError':
        set.status = 500
        return { message: 'Internal Server Error' }
    }
  })
  .get('/', () => 'Hello Elysia')
  .use(auth)
  .listen(3000)

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}!`
)

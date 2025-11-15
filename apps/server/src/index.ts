import { Elysia } from 'elysia'
import cors from '@elysiajs/cors'
import openapi from '@elysiajs/openapi'
import jwt from '@elysiajs/jwt'
import { auth } from './auth'
import { accounts } from './accounts'

const app = new Elysia()
  .use(cors())
  .use(
    jwt({
      name: 'jwt',
      secret: Bun.env.JWT_SECRET || 'supersecret',
    })
  )
  .use(openapi())
  // .error({
  //   NotFoundError,
  //   BadRequestError,
  //   InternalServerError,
  // })
  // .onError(({ code, set, error }) => {
  //   console.log('is here?', code, error)
  //
  //   switch (code) {
  //     case 'NotFoundError':
  //       set.status = 404
  //       return { message: error.message }
  //     case 'BadRequestError':
  //       set.status = 400
  //       return { message: error.message }
  //     case 'InternalServerError':
  //       set.status = 500
  //       return { message: 'Internal Server Error' }
  //   }
  // })
  .get('/', () => 'Hello Elysia')
  .use(auth)
  .use(accounts)
  .listen(3000)

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}!`
)

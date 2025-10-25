import { Elysia } from 'elysia'
import cors from '@elysiajs/cors'
import openapi from '@elysiajs/openapi'
import { auth } from './auth'
import { accounts } from './accounts'
import { AuthMiddleware } from './auth/middleware'
import { transactions } from './transactions'

const app = new Elysia()
  .use(cors())
  .use(openapi())
  // Auth endpoints shouldn't be protected
  .use(auth)
  .use(AuthMiddleware)
  .get('/', () => 'This endpoint is protected!')
  .use(accounts)
  .use(transactions)
  .listen(3000)

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}!`
)

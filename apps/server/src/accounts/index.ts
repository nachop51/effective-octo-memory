import Elysia, { t } from 'elysia'
import { AccountService } from './service'
import { AuthMiddleware } from '../auth/middleware'
import { AccountCurrency } from 'db'

export const accounts = new Elysia({ prefix: '/accounts' })
  .use(AuthMiddleware)
  .get('/', ({ token: { user } }) => {
    const accounts = AccountService.getUserAccounts({
      userId: user.id,
    })

    return accounts
  })
  .get('/:id', ({ params, token: { user } }) => {
    const account = AccountService.getAccountById({
      accountId: params.id,
      userId: user.id,
    })

    return account
  })
  .post(
    '/',
    async ({
      body: { name, number = null, currency = 'MUR' },
      token: { user },
    }) => {
      console.log({ name, number, currency })

      const account = await AccountService.createAccount({
        userId: user.id,
        number,
        name,
      })

      return account
    },
    {
      body: t.Object({
        name: t.String({ minLength: 1 }),
        number: t.Optional(t.String()),
        currency: t.Optional(t.Enum(AccountCurrency)),
      }),
    }
  )
  .patch('/:id', ({ params }) => `Update account with ID: ${params.id}`)
  .delete('/:id', ({ params }) => `Delete account with ID: ${params.id}`)

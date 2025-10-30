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
      body: {
        name,
        accountNumber = null,
        initialBalance = null,
        currency = 'MUR',
      },
      token: { user },
    }) => {
      const account = await AccountService.createAccount({
        userId: user.id,
        name,
        number: accountNumber,
        currency: currency as AccountCurrency,
        initialBalance: initialBalance?.toString() || '0',
      })

      return account
    },
    {
      body: t.Object({
        name: t.String({ minLength: 1 }),
        accountNumber: t.Optional(t.String()),
        initialBalance: t.Optional(t.Number()),
        currency: t.Optional(t.Enum(AccountCurrency)),
      }),
    }
  )
  .patch('/:id', ({ params }) => `Update account with ID: ${params.id}`)
  .delete('/:id', ({ params }) => `Delete account with ID: ${params.id}`)

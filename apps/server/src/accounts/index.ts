import Elysia, { t } from 'elysia'
import { AccountService } from './service'

export const accounts = new Elysia({ prefix: '/accounts' })
  .get('/', () => {
    const accounts = AccountService.getUserAccounts({
      userId: '1',
    })

    return accounts
  })
  .get(
    '/:id',
    ({ params }) => {
      const account = AccountService.getAccountById({
        accountId: params.id,
        userId: '1',
      })

      return account
    },
    {
      params: t.Object({
        id: t.String(),
      }),
    }
  )
  .post(
    '/',
    async ({ body: { name, number = null, currency = 'MUR' } }) => {
      console.log({ name, number, currency })

      const account = await AccountService.createAccount({
        userId: '1',
        number,
        name,
      })

      return account
    },
    {
      body: t.Object({
        name: t.String(),
        number: t.Optional(t.String()),
        currency: t.Optional(
          t.Enum({
            USD: 'USD',
            EUR: 'EUR',
            GBP: 'GBP',
            MUR: 'MUR',
          })
        ),
      }),
    }
  )
  .patch('/:id', ({ params }) => `Update account with ID: ${params.id}`)

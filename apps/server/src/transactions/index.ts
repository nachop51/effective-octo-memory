import Elysia, { t } from 'elysia'
import { AuthMiddleware } from '../auth/middleware'
import { TransactionsService } from './service'

export const transactions = new Elysia({ prefix: '/transactions/:accountId' })
  .use(AuthMiddleware)
  .get(
    '/',
    ({ params: { accountId }, token: { user }, query: { limit } }) => {
      const transactions = TransactionsService.getTransactionHistory({
        accountId: accountId,
        userId: user.id,
        limit: limit,
      })

      return transactions
    },
    {
      query: t.Object({
        limit: t.Number({ default: 20 }),
      }),
    }
  )
  .post(
    '/receive',
    async ({
      params: { accountId: toAccountId },
      token: { user },
      body: { amount, description },
    }) => {
      const updatedAccount = await TransactionsService.receiveFunds({
        userId: user.id,
        toAccountId: toAccountId,
        amount: amount,
        description: description,
      })

      return updatedAccount
    },
    {
      body: t.Object({
        amount: t.Number({ minimum: 1, maximum: 1e12 }),
        description: t.Optional(t.String()),
      }),
    }
  )
  .post('/transfer', () => 'transfer')

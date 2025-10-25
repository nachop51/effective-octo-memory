import { type Account, type User, accounts, db, transactions } from 'db'
import { and, eq, sql } from 'drizzle-orm'
import { NotFoundError } from 'elysia'

export abstract class TransactionsService {
  static async getTransactionHistory({
    accountId,
    userId,
    limit,
  }: {
    accountId: Account['id']
    userId: User['id']
    limit: number
  }) {
    const account = await db.query.accounts.findFirst({
      where: (a, { eq, and }) => and(eq(a.id, accountId), eq(a.userId, userId)),
    })

    if (!account || account.userId !== userId) {
      throw new NotFoundError('Account not found')
    }

    const transactions = await db.query.transactions.findMany({
      where: (t, { eq, and }) => and(eq(t.accountId, accountId)),
      limit: limit,
      orderBy: (t, { desc }) => [desc(t.createdAt)],
    })

    return transactions
  }

  static async receiveFunds({
    userId,
    toAccountId,
    amount,
    description = null,
  }: {
    userId: User['id']
    toAccountId: Account['id']
    amount: number
    description?: string | null
  }) {
    const account = await db.transaction(async (tx) => {
      const [acc] = await tx
        .update(accounts)
        .set({
          balance: sql`${accounts.balance} + ${amount}`,
        })
        .where(and(eq(accounts.id, toAccountId), eq(accounts.userId, userId)))
        .returning({
          newBalance: accounts.balance,
        })

      if (!acc) {
        throw new NotFoundError('Account not found')
      }

      await tx.insert(transactions).values({
        accountId: toAccountId,
        amount: amount.toString(),
        type: 'income',
        description,
      })

      return acc
    })

    return account
  }

  // static async transferFunds({
  //   userId,
  //   fromAccountId,
  //   toAccountId,
  //   amount,
  // }: {
  //   userId: User['id']
  //   fromAccountId: Account['id']
  //   toAccountId: Account['id']
  //   amount: number
  // }) {}
}

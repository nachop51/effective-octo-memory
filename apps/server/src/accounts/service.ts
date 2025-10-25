import { Account, accounts, db, User } from 'db'
import { eq, and } from 'drizzle-orm'
import { NotFoundError } from 'elysia'

export abstract class AccountService {
  static async getUserAccounts({ userId }: { userId: User['id'] }) {
    const accounts = await db.query.accounts.findMany({
      where: (t, { eq }) => eq(t.userId, userId),
    })

    return accounts
  }

  static async getAccountById({
    accountId,
    userId,
  }: {
    accountId: Account['id']
    userId: User['id']
  }) {
    const account = await db.query.accounts.findFirst({
      where: (t, { eq, and }) => and(eq(t.id, accountId), eq(t.userId, userId)),
    })

    if (!account) {
      throw new NotFoundError('Account not found')
    }

    return account
  }

  static async createAccount({
    userId,
    name,
    number,
  }: {
    userId: User['id']
    name: Account['name']
    number: Account['number']
  }) {
    const [account] = await db
      .insert(accounts)
      .values({
        userId,
        name,
        number,
      })
      .returning()

    return account
  }

  static async updateAccount({
    userId,
    accountId,
    number,
    name,
  }: {
    userId: User['id']
    accountId: Account['id']
    number: Account['number']
    name: Account['name']
  }) {
    const [account] = await db
      .update(accounts)
      .set({ name, number })
      .where(and(eq(accounts.userId, userId), eq(accounts.id, accountId)))
      .returning()

    if (!account) {
      throw new NotFoundError('Account not found')
    }

    return account
  }
}

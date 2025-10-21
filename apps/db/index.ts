import { db } from './src'
import { users, accounts, transactions } from './src/schema'

export type User = typeof users.$inferSelect
export type Account = typeof accounts.$inferSelect
export type Transaction = typeof transactions.$inferSelect

export { db, users, accounts, transactions }

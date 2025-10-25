export { db } from './src'
import { users, accounts, transactions, categories } from './src/schema'

export type User = typeof users.$inferSelect
export type Account = typeof accounts.$inferSelect
export type Transaction = typeof transactions.$inferSelect
export type Category = typeof categories.$inferSelect

export { accounts, users, transactions, categories } from './src/schema'

export {
  accountCurrencies,
  AccountCurrency,
  DEFAULT_CURRENCY,
} from './src/consts'

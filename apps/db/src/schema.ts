import { numeric, pgTable, text, check } from 'drizzle-orm/pg-core'
import { createdAt, cuid, updatedAt } from './custom-types'
import { sql } from 'drizzle-orm'

export const users = pgTable('users', {
  id: cuid(),
  email: text().notNull().unique(),
  password: text().notNull(),
  role: text({ enum: ['user', 'admin'] })
    .default('user')
    .notNull(),
  createdAt: createdAt(),
  updatedAt: updatedAt(),
})

export const accounts = pgTable('accounts', {
  id: cuid(),
  name: text().notNull(),
  userId: text()
    .notNull()
    .references(() => users.id),
  createdAt: createdAt(),
  updatedAt: updatedAt(),
})

export const transactions = pgTable(
  'transactions',
  {
    id: cuid(),
    accountId: text()
      .notNull()
      .references(() => accounts.id),
    amount: numeric({ precision: 12, scale: 2 }).notNull(),
    type: text({ enum: ['income', 'expense', 'transfer'] }).notNull(),
    description: text().notNull(),
    createdAt: createdAt(),
    updatedAt: updatedAt(),
  },
  (t) => [check('amount_positive', sql`${t.amount} >= 1`)]
)

import { numeric, pgTable, text, check, varchar } from 'drizzle-orm/pg-core'
import { createdAt, cuid, updatedAt } from './custom-types'
import { relations, sql } from 'drizzle-orm'

export const users = pgTable('users', {
  id: cuid(),
  email: varchar().notNull().unique(),
  password: text().notNull(),
  role: varchar({ enum: ['user', 'admin'] })
    .default('user')
    .notNull(),
  createdAt: createdAt(),
  updatedAt: updatedAt(),
})

export const accounts = pgTable('accounts', {
  id: cuid(),
  number: varchar(),
  name: varchar().notNull(),
  userId: varchar()
    .notNull()
    .references(() => users.id),
  balance: numeric({ precision: 12, scale: 2 })
    .default(sql`0`)
    .notNull(),
  currency: varchar({ length: 3 }).default('USD').notNull(),
  createdAt: createdAt(),
  updatedAt: updatedAt(),
})

export const transactions = pgTable(
  'transactions',
  {
    id: cuid(),
    accountId: varchar()
      .notNull()
      .references(() => accounts.id),
    amount: numeric({ precision: 12, scale: 2 }).notNull(),
    type: varchar({ enum: ['income', 'expense', 'transfer'] }).notNull(),
    description: varchar().notNull(),
    createdAt: createdAt(),
    updatedAt: updatedAt(),
    categoryId: varchar().references(() => categories.id),
  },
  (t) => [check('amount_positive', sql`${t.amount} >= 1`)]
)

export const categories = pgTable('categories', {
  id: cuid(),
  name: varchar().notNull(),
})

// --------------------------------------------
// Relations
// --------------------------------------------

export const accountsRelations = relations(accounts, ({ one, many }) => ({
  user: one(users, {
    fields: [accounts.userId],
    references: [users.id],
  }),
  transactions: many(transactions),
}))

export const usersRelations = relations(users, ({ many }) => ({
  accounts: many(accounts),
}))

export const transactionsRelations = relations(transactions, ({ one }) => ({
  account: one(accounts, {
    fields: [transactions.accountId],
    references: [accounts.id],
  }),
  category: one(categories, {
    fields: [transactions.categoryId],
    references: [categories.id],
  }),
}))

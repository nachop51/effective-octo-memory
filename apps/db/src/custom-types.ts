import { init } from '@paralleldrive/cuid2'
import { text, timestamp, varchar } from 'drizzle-orm/pg-core'

const createId = init({ length: 20, fingerprint: 'apply-hbtn-mur-fingerprint' })

export const cuid = (name: string = 'id') =>
  varchar(name)
    .primaryKey()
    .notNull()
    .$defaultFn(() => createId())

export const createdAt = () => timestamp().defaultNow().notNull()

export const updatedAt = () =>
  timestamp()
    .defaultNow()
    .notNull()
    .$onUpdateFn(() => new Date())

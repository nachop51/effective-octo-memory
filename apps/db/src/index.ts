import { drizzle } from 'drizzle-orm/node-postgres'
import * as schema from './schema'

const conn = Bun.env.DATABASE_URL

if (!conn) {
  throw new Error(
    'DATABASE_URL is not set, probably something wrong about how the script is being executed'
  )
}

export const db = drizzle(conn, { schema })

db.$client
  .query('SELECT 1')
  .then(() => {
    console.log('Database connected successfully')
  })
  .catch((err) => {
    console.error('Failed to connect to the database:', err)
  })

import { db, users } from 'db'

const user = await db.insert(users).values({
  email: 'John.doe@gmail.com',
  password: await Bun.password.hash('securepassword123'),
  id: '1',
})

console.log({ user })

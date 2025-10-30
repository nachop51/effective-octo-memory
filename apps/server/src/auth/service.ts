import { db, users } from 'db'
import { InternalServerError } from 'elysia'
import { ConflictError, UnauthorizedError } from './errors'

export abstract class AuthService {
  static async signIn({
    email,
    password,
  }: {
    email: string
    password: string
  }) {
    const user = await db.query.users.findFirst({
      where: (t, { eq }) => eq(t.email, email),
      columns: {
        id: true,
        email: true,
        password: true,
      },
    })

    if (!user) {
      throw new UnauthorizedError('Invalid credentials')
    }

    if (!(await Bun.password.verify(password, user.password))) {
      throw new UnauthorizedError('Invalid credentials')
    }

    return user
  }

  static async signUp({
    email,
    password,
  }: {
    email: string
    password: string
  }) {
    const user = await db.query.users.findFirst({
      where: (t, { eq }) => eq(t.email, email),
      columns: {
        id: true,
      },
    })

    if (user) {
      throw new ConflictError('User already exists')
    }

    const hashedPassword = await Bun.password.hash(password)

    const [newUser] = await db
      .insert(users)
      .values({
        email,
        password: hashedPassword,
      })
      .returning()

    if (!newUser) {
      throw new InternalServerError('Something went wrong when creating user')
    }

    return newUser
  }
}

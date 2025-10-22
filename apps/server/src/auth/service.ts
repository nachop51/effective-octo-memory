import { db, User, users } from 'db'
import { NotFoundError } from 'elysia'
import { ConflictError } from './errors'

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
        password: true,
      },
    })

    if (!user) {
      throw new NotFoundError('Invalid credentials')
    }

    if (!(await Bun.password.verify(password, user.password))) {
      throw new NotFoundError('Invalid credentials')
    }

    return {
      email,
      token: 'fake-jwt-token',
    }
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
    })

    if (user) {
      throw new ConflictError('User already exists')
    }

    const hashedPassword = await Bun.password.hash(password)

    const newUser = await db
      .insert(users)
      .values({
        email,
        password: hashedPassword,
      })
      .returning()

    return newUser[0]
  }

  static async generateToken(user: User) {
    console.log({ user })
  }
}

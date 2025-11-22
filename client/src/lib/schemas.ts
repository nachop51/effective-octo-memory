import z from 'zod'

export const loginSchema = z.object({
  email: z.email('Please enter a valid email address'),
  password: z
    .string()
    .min(8, 'The password needs to be at least 8 characters long'),
})

export type LoginFormData = z.infer<typeof loginSchema>

export const signupSchema = z
  .object({
    email: z.email('Please enter a valid email address'),
    password: z
      .string()
      .min(8, 'The password needs to be at least 8 characters long'),
    confirmPassword: z
      .string()
      .min(8, 'The password needs to be at least 8 characters long'),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: 'Passwords do not match',
  })

export type SignupFormData = z.infer<typeof signupSchema>

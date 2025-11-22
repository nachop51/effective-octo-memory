import { GalleryVerticalEnd } from 'lucide-react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import {
  Field,
  FieldDescription,
  FieldLabel,
  FieldGroup,
} from '@/components/ui/field'
import { Input } from '@/components/ui/input'
import { Link, useLocation } from 'wouter'
import { useState } from 'react'
import api from '@/lib/api/effective'
import { useAuthStore } from '@/lib/stores/auth'
import { loginSchema, type LoginFormData } from '@/lib/schemas'
import type { AuthResponse } from '@/lib/types'

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<'div'>) {
  const [location, navigate] = useLocation()
  const [isLoading, setIsLoading] = useState(false)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const login = useAuthStore((state) => state.login)

  const onSubmit = async (data: LoginFormData) => {
    setIsLoading(true)
    setErrorMessage(null)

    try {
      const response = await api.post('auth/signin', {
        json: {
          email: data.email,
          password: data.password,
        },
      })

      if (response.status === 401) {
      } else if (response.status === 200) {
        const data = await response.json() as AuthResponse
        await login(data.cookie)
        navigate('/')
      }
    } catch (error) {
      console.error('Login error:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleFormSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    handleSubmit(onSubmit)(e)
  }

  return (
    <div className={cn('flex flex-col gap-6', className)} {...props}>
      <div onSubmit={handleFormSubmit}>
        <FieldGroup>
          <div className="flex flex-col items-center gap-2 text-center">
            <a
              href="#"
              className="flex flex-col items-center gap-2 font-medium"
            >
              <div className="flex size-8 items-center justify-center rounded-md">
                <GalleryVerticalEnd className="size-6" />
              </div>
              <span className="sr-only">Effect.</span>
            </a>
            <h1 className="text-xl font-bold">
              Welcome to Effective Octo Memory
            </h1>
            <FieldDescription>
              Don&apos;t have an account? <Link href="/signup">Sign up</Link>
            </FieldDescription>
          </div>

          {errorMessage && (
            <div className="rounded-md bg-red-50 p-3 text-sm text-red-800 border border-red-200">
              {errorMessage}
            </div>
          )}

          <Field>
            <FieldLabel htmlFor="email">Email</FieldLabel>
            <Input
              id="email"
              type="email"
              placeholder="e@example.com"
              {...register('email')}
              aria-invalid={errors.email ? 'true' : 'false'}
            />
            {errors.email && (
              <p className="text-sm text-red-600 mt-1">
                {errors.email.message}
              </p>
            )}
          </Field>

          <Field>
            <FieldLabel htmlFor="password">Password</FieldLabel>
            <Input
              id="password"
              type="password"
              placeholder="********"
              {...register('password')}
              aria-invalid={errors.password ? 'true' : 'false'}
            />
            {errors.password && (
              <p className="text-sm text-red-600 mt-1">
                {errors.password.message}
              </p>
            )}
          </Field>

          <Field>
            <Button
              type="button"
              onClick={handleSubmit(onSubmit)}
              disabled={isLoading}
              className="w-full"
            >
              {isLoading ? 'Initializing session...' : 'Login'}
            </Button>
          </Field>

          <div className="text-center">
            <a
              href="#"
              className="text-sm text-muted-foreground hover:text-primary"
            >
              ¿Olvidaste tu contraseña?
            </a>
          </div>
        </FieldGroup>
      </div>
    </div>
  )
}

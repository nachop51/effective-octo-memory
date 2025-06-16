import { API_URL } from '$lib/constants'
import type { ErrorResponse, User } from '$lib/types'
import { handlePromise, type HandlePromiseResult } from '$lib/utils/fns'
import paths from '$lib/utils/paths'
import ky, { HTTPError } from 'ky'
import { goto } from '$app/navigation'

type LoginResponse = {
	user: User
	access_token: string
}

export const effectiveBack = ky.create({
	prefixUrl: API_URL,
	credentials: 'include',
	headers: {
		'Content-Type': 'application/json'
	},
	hooks: {
		beforeError: [
			async (err) => {
				if (err.response.status === 401) {
					goto(paths.login())
				}
				return err
			}
		]
	}
})

export async function apiHealthCheck(): Promise<boolean> {
	try {
		const res = await fetch(API_URL + '/health')

		if (!res.ok) {
			throw new Error('Health check failed')
		}

		console.log('Health check passed')

		return true
	} catch (error) {
		console.error('Health check error:', error)
		return false
	}
}

export async function logIn({
	email,
	password
}: {
	email: string
	password: string
}): HandlePromiseResult<LoginResponse> {
	const [res, err] = await handlePromise(
		effectiveBack.post<LoginResponse>(paths.api.login(), {
			json: { email, password }
		})
	)

	if (err) {
		console.error('Login failed:', err)
		return [null, err]
	}

	const data = await res.json()

	return [data, null]
}

export async function signUp(
	user: Omit<User, 'id'> & { confirmPassword: User['password'] }
): HandlePromiseResult<LoginResponse, ErrorResponse> {
	const [res, err] = await handlePromise(
		effectiveBack.post(paths.api.signup(), {
			json: user
		})
	)

	if (err) {
		if (err instanceof HTTPError) {
			const errorData = await err.response.json<ErrorResponse>()

			return [null, errorData]
		}

		return [null, { message: 'Unknown error occurred', code: 500 }]
	}

	const data = await res.json<LoginResponse>()

	return [data, null]
}

export async function logOut(): HandlePromiseResult<null> {
	const [, err] = await handlePromise(effectiveBack.post<''>(paths.api.logout()))

	if (err) {
		console.error('Logout failed:', err)
		return [null, err]
	}

	return [null, null]
}

export async function checkUser(): HandlePromiseResult<User> {
	const [res, err] = await handlePromise(effectiveBack.get(paths.api.checkUser()))

	if (err) {
		return [null, err]
	}

	const user = await res.json<User>()

	return [user, null]
}

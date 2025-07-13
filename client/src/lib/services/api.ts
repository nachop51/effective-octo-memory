import { API_URL } from '$lib/constants'
import type { User } from '$lib/types'
import { handlePromise } from '$lib/utils/fns'
import paths from '$lib/utils/paths'

export async function healthCheck(): Promise<boolean> {
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
}): Promise<User | null> {
	const res = await handlePromise(
		fetch(paths.api.login(), {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ email, password })
		})
	)

	console.log('Login response:', res)

	if (!res.ok) {
		console.error('Login failed:', res.error)
		return null
	}

	return await res.data.json()
}

export async function signUp({
	email,
	password,
	confirmPassword
}: {
	email: string
	password: string
	confirmPassword: string
}): Promise<User | null> {
	const result = await handlePromise(
		fetch(paths.api.signup(), {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ email, password, confirmPassword })
		})
	)

	if (!result.ok) {
		return null
	}

	const user = await result.data.json()
	if (!user) {
		console.error('Sign up failed: No user returned')
		return null
	}
	return user
}

export async function checkUser(): Promise<User | null> {
	const result = await handlePromise(
		fetch(paths.api.checkUser(), {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		})
	)

	if (!result.ok) {
		console.error('Check login failed:', result.error)
		return null
	}

	const user = await result.data.json()
	if (!user) {
		console.error('Check login failed: No user returned')
		return null
	}

	return user
}

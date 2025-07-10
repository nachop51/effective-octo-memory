import { API_URL } from '$lib/constants'
import type { User } from '$lib/types'

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

export async function logIn(email: string, password: string): Promise<User | null> {
	try {
		const res = await fetch(API_URL + '/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ email, password })
		})

		if (!res.ok) {
			throw new Error('Login failed')
		}

		const user = await res.json()
		return user
	} catch (error) {
		console.error('Login error:', error)
		return null
	}
}

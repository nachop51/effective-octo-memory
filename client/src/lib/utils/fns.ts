import { AUTH_COOKIE_NAME } from '$env/static/private'
import { DEV_MODE } from '$lib/constants'
import type { Cookies } from '@sveltejs/kit'

export type HandlePromiseResult<T, K = Error> = Promise<[T, null] | [null, K]>

export async function handlePromise<T>(promise: Promise<T>): HandlePromiseResult<T> {
	try {
		const result = await promise

		return [result, null]
	} catch (error) {
		return [null, error instanceof Error ? error : new Error('Unknown error')]
	}
}

export function setAuthCookie(cookies: Cookies, value: string) {
	cookies.set(AUTH_COOKIE_NAME, value, {
		path: '/',
		httpOnly: true,
		secure: !DEV_MODE,
		sameSite: 'strict',
		maxAge: 60 * 60 * 24 * 30
	})
}

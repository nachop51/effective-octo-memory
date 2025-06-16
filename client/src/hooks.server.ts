import { AUTH_COOKIE_NAME, JWT_SECRET_KEY } from '$env/static/private'
import { UNPROTECTED_ROUTES } from '$lib/constants'
import type { User } from '$lib/types'
import paths from '$lib/utils/paths'
import { redirect } from '@sveltejs/kit'
import jwt, { type JwtPayload } from 'jsonwebtoken'

export async function handle({ event, resolve }) {
	const pathname = event.url.pathname
	console.log('Hook running for:', pathname)

	const token = event.cookies.get(AUTH_COOKIE_NAME)

	if (!token) {
		if (!UNPROTECTED_ROUTES.has(pathname)) {
			throw redirect(302, '/login')
		}
		return resolve(event)
	}

	try {
		const decoded = jwt.verify(token, JWT_SECRET_KEY) as JwtPayload & { user: User }
		event.locals.user = decoded.user
	} catch (err) {
		console.error('Invalid token:', err)
		event.cookies.delete(AUTH_COOKIE_NAME, { path: '/' })
		return resolve(event)
	}

	if (UNPROTECTED_ROUTES.has(pathname)) {
		throw redirect(302, paths.home())
	}

	return resolve(event)
}

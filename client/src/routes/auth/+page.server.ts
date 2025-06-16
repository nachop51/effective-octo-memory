import { logIn } from '$lib/services/api'
import { fail, redirect } from '@sveltejs/kit'
import paths from '$lib/utils/paths'
import { AUTH_COOKIE_NAME } from '$env/static/private'
import { DEV_MODE } from '$lib/constants'
import { setAuthCookie } from '$lib/utils/fns.js'

export const actions = {
	logIn: async ({ cookies, request }) => {
		const formData = await request.formData()
		const email = formData.get('email') as string
		const password = formData.get('password') as string

		if (!email || !password) {
			return { error: 'Email and password are required.' }
		}

		const [data, err] = await logIn({ email, password })

		if (err != null) {
			return fail(400, {
				error: 'Invalid credentials'
			})
		}

		setAuthCookie(cookies, data.access_token)

		throw redirect(302, paths.home())
	},
	logOut: async ({ cookies }) => {
		cookies.delete(AUTH_COOKIE_NAME, {
			path: '/',
			httpOnly: true,
			secure: !DEV_MODE,
			sameSite: 'strict'
		})

		throw redirect(302, paths.login())
	}
}

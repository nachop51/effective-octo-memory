import { signUp } from '$lib/services/api'
import { setAuthCookie } from '$lib/utils/fns.js'
import paths from '$lib/utils/paths.js'
import { fail, redirect } from '@sveltejs/kit'

export const actions = {
	default: async ({ cookies, request }) => {
		const formData = await request.formData()
		const firstName = formData.get('firstName') as string
		const lastName = formData.get('lastName') as string
		const email = formData.get('email') as string
		const password = formData.get('password') as string
		const confirmPassword = formData.get('confirmPassword') as string

		if (!firstName || !lastName || !email || !password || !confirmPassword) {
			return fail(400, {
				error: 'All fields are required.'
			})
		}

		const [data, err] = await signUp({
			firstName,
			lastName,
			email,
			password,
			confirmPassword
		})

		console.log({ data, err })

		if (err != null) {
			return fail(err.code, {
				error: err.message
			})
		}

		console.log('Sign up successful, setting auth cookie:', { data })

		setAuthCookie(cookies, data.access_token)

		throw redirect(302, paths.home())
	}
}

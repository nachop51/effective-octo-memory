import { goto } from '$app/navigation'
import { page } from '$app/state'
import {
	logIn as logInService,
	checkUser as checkUserService,
	signUp as signUpService
} from '$lib/services/api'
import type { User } from '$lib/types.d.ts'
import { getContext, onMount, setContext } from 'svelte'

interface IUserStore {
	user: User | null
	isLoading: boolean
	logIn: ({ email, password }: Pick<User, 'email' | 'password'>) => Promise<User | null>
	signUp: ({
		email,
		password,
		confirmPassword
	}: Pick<User, 'email' | 'password'> & { confirmPassword: string }) => Promise<User | null>
	checkLogIn: () => Promise<User | null>
	logout: () => Promise<void>
}

class UserStore implements IUserStore {
	user: User | null = $state(null)
	isLoading: boolean = $state(false)

	constructor() {
		onMount(async () => {
			const res = await this.checkLogIn()

			if (res) {
				this.user = res
			}

			if (!res && page.url.pathname !== '/login') {
				goto('/login')
			}
		})
	}

	async logIn({ email, password }: Pick<User, 'email' | 'password'>): Promise<User | null> {
		this.isLoading = true

		const user = await logInService({
			email,
			password: password!
		})

		this.isLoading = false

		if (!user) {
			return null
		}

		this.user = user

		return user
	}

	async signUp({
		email,
		password,
		confirmPassword
	}: Pick<User, 'email' | 'password'> & { confirmPassword: string }): Promise<User | null> {
		this.isLoading = true

		const user = await signUpService({
			email,
			password: password!,
			confirmPassword
		})

		this.isLoading = false

		if (!user) {
			return null
		}

		this.user = user

		return user
	}

	async checkLogIn(): Promise<User | null> {
		this.isLoading = true

		const user = await checkUserService()

		this.isLoading = false

		return user
	}

	async logout(): Promise<void> {
		this.isLoading = true

		try {
			// TODO: Implement actual logout logic here
		} catch (error) {
			console.error('Logout failed:', error)
		} finally {
			this.isLoading = false
		}
	}
}

// -------------------------------------------------------------------

const USER_STORE_KEY = '$_user_store'

export function initUserStore() {
	return setContext(USER_STORE_KEY, new UserStore())
}

export function getUserStore() {
	const store = getContext<ReturnType<typeof initUserStore>>(USER_STORE_KEY)

	if (!store) {
		throw new Error('User store not initialized')
	}

	return store
}

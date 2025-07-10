import { goto } from '$app/navigation'
import { page } from '$app/state'
import type { User } from '$lib/types.d.ts'
import { getContext, onMount, setContext } from 'svelte'

export const userState = $state({
	user: null as User | null,
	isLoading: false
})

interface IUserStore {
	user: User | null
	isLoading: boolean
	logIn: (email: string, password: string) => Promise<User | null>
	checkLogIn: () => Promise<User | null>
	updateUser: (user: User) => Promise<User>
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

	async logIn(email: string, password: string): Promise<User | null> {
		this.isLoading = true

		try {
			// TODO: Implement actual login logic here
			await new Promise((resolve) => setTimeout(resolve, 1000))
		} catch (error) {
			console.error('Login failed:', error)
			return null
		} finally {
			this.isLoading = false
		}
	}

	async checkLogIn(): Promise<User | null> {
		this.isLoading = true

		try {
			// TODO: Implement actual check login logic here
		} catch (error) {
			console.error('Check login failed:', error)
			return null
		} finally {
			this.isLoading = false
		}
	}
	async updateUser(user: User): Promise<User | null> {
		this.isLoading = true

		try {
			// TODO: Implement actual update user logic here
		} catch (error) {
			console.error('Update user failed:', error)
			throw error
		} finally {
			this.isLoading = false
		}
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

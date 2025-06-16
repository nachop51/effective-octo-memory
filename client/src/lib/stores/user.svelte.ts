// import { goto } from '$app/navigation'
// import { page } from '$app/state'
// import {
// 	logIn as logInService,
// 	checkUser as checkUserService,
// 	signUp as signUpService,
// 	logOut as logoutService
// } from '$lib/services/api'
// import type { User } from '$lib/types.d.ts'
// import { getContext, onMount, setContext } from 'svelte'

// interface IUserStore {
// 	user: User | null
// 	isLoading: boolean
// 	logIn: ({ email, password }: Pick<User, 'email' | 'password'>) => Promise<User | null>
// 	signUp: (data: Omit<User, 'id'> & { confirmPassword: string }) => Promise<User | null>
// 	checkLogIn: () => Promise<User | null>
// 	logout: () => Promise<void>
// }

// class UserStore implements IUserStore {
// 	user: User | null = $state(null)
// 	isLoading: boolean = $state(true)

// 	constructor() {
// 		onMount(async () => {
// 			const res = await this.checkLogIn()

// 			if (res) {
// 				this.user = res
// 			}

// 			if (!res && page.url.pathname !== '/login') {
// 				goto('/login')
// 			}
// 		})
// 	}

// 	async logIn({ email, password }: Pick<User, 'email' | 'password'>): Promise<User | null> {
// 		const [user, err] = await logInService({
// 			email,
// 			password: password!
// 		})

// 		if (err) {
// 			return null
// 		}

// 		this.user = user

// 		return user
// 	}

// 	async signUp({
// 		firstName,
// 		lastName,
// 		email,
// 		password,
// 		confirmPassword
// 	}: Omit<User, 'id'> & { confirmPassword: string }): Promise<User | null> {
// 		const [user, err] = await signUpService({
// 			firstName,
// 			lastName,
// 			email,
// 			password: password!,
// 			confirmPassword
// 		})

// 		if (err) {
// 			return null
// 		}

// 		this.user = user

// 		return user
// 	}

// 	async checkLogIn(): Promise<User | null> {
// 		this.isLoading = true

// 		const [user, err] = await checkUserService()

// 		this.isLoading = false

// 		if (err) {
// 			// console.error('Check login failed:', res.error)
// 			return null
// 		}

// 		this.user = user

// 		return user
// 	}

// 	async logout(): Promise<void> {
// 		const [, err] = await logoutService()

// 		if (err) {
// 			console.error('Logout failed:', err)
// 			return
// 		}
// 		this.user = null
// 	}
// }

// // -------------------------------------------------------------------

export const USER_STORE_KEY = '$_user_store'

// export function initUserStore() {
// 	return setContext(USER_STORE_KEY, new UserStore())
// }

// export function getUserStore() {
// 	const store = getContext<ReturnType<typeof initUserStore>>(USER_STORE_KEY)

// 	if (!store) {
// 		throw new Error('User store not initialized')
// 	}

// 	return store
// }

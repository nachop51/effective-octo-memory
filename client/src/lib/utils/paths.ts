import { API_URL } from '$lib/constants'

const paths = {
	login: () => '/login',
	home: () => '/home',
	dashboard: () => '/dashboard',
	profile: () => '/profile',
	settings: () => '/settings',
	api: {
		login: () => API_URL + '/login',
		signup: () => API_URL + '/signup',
		checkUser: () => API_URL + '/check',
		logout: () => API_URL + '/logout',
		health: () => API_URL + '/health'
	}
}

export default paths

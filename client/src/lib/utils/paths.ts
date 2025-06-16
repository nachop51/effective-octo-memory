const paths = {
	home: () => '/',
	login: () => '/login',
	signup: () => '/signup',
	dashboard: () => '/dashboard',
	profile: () => '/profile',
	settings: () => '/settings',
	api: {
		login: () => 'login',
		signup: () => 'signup',
		checkUser: () => 'check',
		logout: () => 'logout',
		health: () => 'health'
	}
}

export default paths

export const DEV_MODE = import.meta.env.DEV

export const API_URL = DEV_MODE ? 'http://localhost:1234' : 'https://back:1234'

export const UNPROTECTED_ROUTES = new Set(['/login', '/signup', '/auth'])

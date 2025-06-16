export interface User {
	id: string
	firstName: string
	lastName: string
	email: string
	password?: string
}

export interface ErrorResponse {
	message: string
	code: number
	details?: Record<string, string>
}

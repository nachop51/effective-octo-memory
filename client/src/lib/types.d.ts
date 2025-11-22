export interface AuthResponse {
  cookie: string
  username?: string
}

export interface CheckAuthResponse {
  username: string
  email?: string
}

export interface ErrorResponse {
  error: string
  message?: string
}

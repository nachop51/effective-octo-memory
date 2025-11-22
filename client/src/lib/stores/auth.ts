import { create } from 'zustand'
import api from '../api/effective'
import type { CheckAuthResponse } from '../types'

interface AuthState {
  user: string | null
  isLoading: boolean
  isAuthenticated: boolean
  checkAuth: () => Promise<void>
  login: (token: string) => Promise<void>
  logout: () => void
}

export const useAuthStore = create<AuthState>()((set) => ({
  user: null,
  isLoading: true,
  isAuthenticated: false,

  checkAuth: async () => {
    set({ isLoading: true })
    try {
      const cookie = localStorage.getItem('auth')

      if (!cookie) {
        set({ isLoading: false, isAuthenticated: false, user: null })
        return
      }

      const data = await api
        .get('auth/check', {
          headers: { Authorization: `Bearer ${cookie}` },
        })
        .json() as CheckAuthResponse

      set({
        user: data.username,
        isAuthenticated: true,
        isLoading: false,
      })
    } catch (error) {
      localStorage.removeItem('auth')
      console.error(error)
      set({ isLoading: false, isAuthenticated: false, user: null })
    }
  },

  login: async (token: string) => {
    localStorage.setItem('auth', token)
    set({ isAuthenticated: true })
  },

  logout: () => {
    localStorage.removeItem('auth')
    set({ user: null, isAuthenticated: false })
  },
}))

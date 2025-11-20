import ky from 'ky'

const api = ky.create({
  prefixUrl: import.meta.env.VITE_API_URL,
  credentials: 'include', // 4 cookies
})

export default api

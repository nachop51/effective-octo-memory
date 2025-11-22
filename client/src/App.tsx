import { Route, Switch } from 'wouter'
import { LoginForm } from '@/components/login-form'
import { SignupForm } from '@/components/signup-form'
import { ProtectedRoute } from '@/components/protected-route'
import { useAuthStore } from '@/lib/stores/auth'
import { useEffect } from 'react'
import '@/App.css'

function App() {
  const checkAuth = useAuthStore((state) => state.checkAuth)
  const isLoading = useAuthStore((state) => state.isLoading)

  useEffect(() => {
    checkAuth()
  }, [checkAuth])

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <main>
      <Switch>
        <Route path="/">
          <ProtectedRoute>Home</ProtectedRoute>
        </Route>

        <Route path="/login">
          <LoginForm />
        </Route>

        <Route path="/signup">
          <SignupForm />
        </Route>

        <Route>404: No such page!</Route>
      </Switch>
    </main>
  )
}

export default App

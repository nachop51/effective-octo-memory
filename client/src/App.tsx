import { Route, Switch } from 'wouter'
import { LoginForm } from './components/login-form'
import './App.css'

function App() {
  return (
    <main>
      <Switch>
        <Route path="/">Home</Route>

        <Route path="/login">
          <LoginForm />
        </Route>

        <Route path="/signup">SignUp</Route>

        <Route>404: No such page!</Route>
      </Switch>
    </main>
  )
}

export default App

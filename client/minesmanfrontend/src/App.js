import { Fragment, useEffect } from 'react'
import "./App.css"
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import AuthState from './context/AuthState'
import Console from './Console'
import Navbar from './Nav'
import Login from './Login'
// sedoicnsdklvn
import { X } from "./Socket"

function App() {
  useEffect(() => {
    X();
   /*  return () => {
      // cleanup
    } */
  }, [])
  return (
    <AuthState>
      <Navbar />
      <Router>
        <Fragment>
          <Switch>
            <Route exact path='/' component={Login} />
            <Route exact path='/console' component={Console} />
          </Switch>
        </Fragment>
      </Router>
    </AuthState>
  );
}

export default App;

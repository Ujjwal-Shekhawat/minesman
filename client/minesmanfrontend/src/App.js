import { Fragment, useEffect } from 'react'
import "./App.css"
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import AuthState from './context/AuthState'
import Console from './Console'
import Navbar from './Nav'
import Login from './Login'

function App() {
  return (
    <AuthState>
      <Router>
        <Fragment>
          <Navbar />
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

import React, { useContext, Fragment } from 'react'
import authContext from "./context/authContext"

function Nav() {
    const authCtx = useContext(authContext)
    const { isAuth, logout } = authCtx

    const logoff = () => {
        logout()
    }

    const loggedin = (
        <div>
            <nav className="navbar navbar-expand-lg navbar-light bg-light" style={{ padding: '10px' }}>
                <a className="navbar-brand" >MINESMAN</a>
                <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>

                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item active">
                            <a className="nav-link" href="/">Login</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="/register">Register</a>
                        </li>
                        <li className="nav-item">
                            <button type="button" className="nav-link" onClick={logoff}>Logout</button>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>
    )

    const loggedout = (
        <div>
            <nav className="navbar navbar-expand-lg navbar-light bg-light" style={{ padding: '10px' }}>
                <a className="navbar-brand" >MINESMAN</a>
                <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>

                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item active">
                            <a className="nav-link" href="/">Login</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link" href="/register">Register</a>
                        </li>
                        {/* <li className="nav-item">
                            <button className="nav-link" onClick={logoff}>Logout</button>
                        </li> */}
                    </ul>
                </div>
            </nav>
        </div>
    )
    return (
        <Fragment>
            {(isAuth == true) ? loggedin : loggedout}
        </Fragment>
    )
}

export default Nav;
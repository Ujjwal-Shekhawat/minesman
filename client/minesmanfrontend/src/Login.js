import { Fragment, React, useContext, useEffect, useState } from 'react'
import authContext from './context/authContext';
import { Redirect } from 'react-router-dom'

export default function Login(props) {

    const [formData, setformData] = useState({
        username: '',
        password: '',
    });

    const { username, password } = formData;

    const authCtx = useContext(authContext)
    const { login, isAuth, authUser, loading } = authCtx

    useEffect(() => {
        authCtx.authUser()
        console.log("use effect called from login")
    }, [isAuth])

    const onChange = (e) => {
        setformData({ ...formData, [e.target.name]: e.target.value })
    }

    const onSubmit = (e) => {
        e.preventDefault();
        console.log(formData)
        if (username === '' || password === '') {
            alert("Fill out the feilds")
        } else {
            login({ username, password })
        }
    }

    const lform = (
        <div className="container">
            <form onSubmit={onSubmit}>
                <div className="mb-3">
                    <label htmlFor="userName" className="form-label">User Name</label>
                    <input type="text" className="form-control" id="userName" aria-describedby="emailHelp" name='username' value={username} onChange={onChange} />
                </div>
                <div className="mb-3">
                    <label htmlFor="exampleInputPassword1" className="form-label">Password</label>
                    <input type="password" className="form-control" id="exampleInputPassword1" name='password' value={password} onChange={onChange} />
                </div>
                <button type="submit" className="btn btn-primary">Submit</button>
            </form>
        </div>
    )

    const redirect = (
        <div>
            <Redirect to='/console' />
        </div>
    )

    return (
        <div>
            {console.log("im rerendering")}
            {(isAuth == true && !loading) ? redirect : (!loading) ? lform : <h1>Loading</h1>}
        </div>
    )
}

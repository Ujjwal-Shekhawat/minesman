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
    const { login, isAuth, authUser } = authCtx

    useEffect(() => {
        /* if (isAuth) {
            props.history.push('/console')
        } */
        authUser()
    }, [])

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

    if (isAuth == false || isAuth == null) {
        { console.log(isAuth) }
        return (
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
    }

    return (
        <div>
            {console.log(isAuth)}
            <Redirect to='/console' />
        </div>
    )
}

import { React, useReducer } from 'react'
import AuthContext from './authContext'
import authReducer from './authReducer'
import axios from 'axios'

const setToken = (token) => {
    if (token) {
        axios.defaults.headers.common['auth-token'] = token
    } else {
        delete axios.defaults.headers.common['auth-token']
    }
}

export default function AuthState(props) {
    const initState = {
        token: localStorage.getItem('token'),
        isAuth: false,
        username: null,
        error: null,
    }

    const [state, dispatch] = useReducer(authReducer, initState);

    const authUser = async () => {
        console.log("authUser was called")
        if (localStorage.getItem("token") != null || localStorage.getItem("token") != undefined) {
            setToken(localStorage.getItem("token"))
        }
        try {
            const result = axios.get('http://20.197.57.10:8080/console')
            dispatch({ type: 'user_authenticated', payload: result.data })
        } catch (error) {
            dispatch({ type: 'logout' })
        }
    }

    const login = async (formData) => {
        const config = { headers: { "Content-Type": "application/json" } }
        try {
            const result = await axios.post('http://20.197.57.10:8080/', formData, config)
            // console.log(result.data)
            dispatch({ type: 'login_success', payload: result.data })
            authUser();
        } catch (error) {
            dispatch({ type: 'logout' })
        }
    }

    const logout = () => {
        console.log("logging out")
        dispatch({ type: 'logout' })
    }

    return (
        <AuthContext.Provider
            value={{
                token: state.token,
                isAuth: state.isAuth,
                username: state.username,
                error: state.error,
                login,
                authUser,
                logout
            }}>
            {props.children}
        </AuthContext.Provider>
    )
}

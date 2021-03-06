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
        token: localStorage.getItem('token', null),
        isAuth: false,
        username: null,
        error: null,
        loading: false,
    }

    const [state, dispatch] = useReducer(authReducer, initState);

    const authUser = async () => {
        // console.log("authUser was called")
        if (localStorage.token) {
            setToken(localStorage.token)
        } else {
            setToken()
        }
        try {
            const result = await axios.get('https://20.193.246.52/console')
            dispatch({ type: 'user_authenticated', payload: result.data })
            console.log(result.data)
        } catch (error) {
            dispatch({ type: 'logout' })
            console.log("auth error", error)
        }
    }

    const login = async (formData) => {
        const config = { headers: { "Content-Type": "application/json" } }
        try {
            const result = await axios.post('https://20.193.246.52/', formData, config)
            // console.log(result.data)
            dispatch({ type: 'login_success', payload: result.data })
            // authUser();
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
                loading: state.loading,
                login,
                authUser,
                logout
            }}>
            {props.children}
        </AuthContext.Provider>
    )
}

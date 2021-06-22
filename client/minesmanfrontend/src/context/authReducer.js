export default (state, action) => {
    switch (action.type) {
        case 'login_success':
            localStorage.setItem('token', action.payload.token);
            return ({
                ...state,
                ...action.payload,
                isAuth: true,
            })
        case 'user_authenticated':
            return ({
                ...state,
                isAuth: true,
                username: action.payload
            })
        case 'logout':
            localStorage.removeItem('token')
            return ({
                ...state,
                token: null,
                isAuth: false,
                username: null
            })
        default:
            return state
    }
}
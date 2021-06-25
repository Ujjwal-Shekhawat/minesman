export default (state, action) => {
    switch (action.type) {
        case 'login_success':
            localStorage.setItem('token', action.payload.token);
            return ({
                ...state,
                ...action.payload,
                isAuth: true,
                loading: false,
            });
        case 'user_authenticated':
            return ({
                ...state,
                ...action.payload,
                username: action.payload,
                // isAuth: true,
                loading: false,
            });
        case 'logout':
            localStorage.removeItem('token')
            return ({
                ...state,
                token: null,
                isAuth: false,
                username: null,
                loading: false,
            });
        default:
            return state
    }
}
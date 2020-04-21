import { LOGIN_REQUESTED } from './types.js';
import { LOGIN_SUCCEEDED } from './types.js';

export const RequestLogin = () => {
    return {
        type: LOGIN_REQUESTED,
        payload: {
        },
    };
}

export const GrantLogin = (username) => {
    return {
        type: LOGIN_SUCCEEDED,
        payload: {
            isLogged: true,
            username: username
        },
    };
}
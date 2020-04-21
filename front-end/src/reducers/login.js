import { LOGIN_REQUESTED } from '../actions/types.js';
import { LOGIN_SUCCEEDED } from '../actions/types.js';

const initialState = {
    isLogged: false,
    username: null
}

export default function loginReducer(state = initialState, action) {
    switch (action.type) {
        case LOGIN_REQUESTED:
            return {
                initialState
            };
        case LOGIN_SUCCEEDED:
            return {
                isLogged: action.payload.isLogged,
                username: action.payload.username
            };
        default:
            return state;
    }
}
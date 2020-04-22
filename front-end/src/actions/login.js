import { LOGIN_REQUESTED } from './types.js';
import { LOGIN_SUCCEEDED } from './types.js';

export const requestLogin = () => {
  return {
    type: LOGIN_REQUESTED,
    payload: {
    },
  };
}

export const grantLogin = (username) => {
  return {
    type: LOGIN_SUCCEEDED,
    payload: {
      isLogged: true,
      username: username
    },
  };
}
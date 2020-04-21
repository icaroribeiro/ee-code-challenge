import { connectRouter } from 'connected-react-router';
import { combineReducers } from 'redux'

import login from './login.js';
import repository from './repository.js';

const rootReducer = (history) => combineReducers({
  loginReducer: login,
  repositoryReducer: repository,
  router: connectRouter(history),
});

export default rootReducer;
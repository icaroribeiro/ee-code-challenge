import thunk from 'redux-thunk';
import { routerMiddleware } from 'connected-react-router';
import { createStore, applyMiddleware } from 'redux';

import history from '../routes/history.js';
import rootReducer from '../reducers/index.js';

const middlewares = [
  thunk,
  routerMiddleware(history),
];

const store = createStore(
  rootReducer(history),
  applyMiddleware(...middlewares)
);

export default store;
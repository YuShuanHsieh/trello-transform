/* eslint-disable no-undef */
/* eslint-disable no-underscore-dangle */
import { createStore, compose, applyMiddleware } from 'redux';
import { createEpicMiddleware } from 'redux-observable';

import { transform } from './reducer';
import { transformEpic } from './epic';

const epicMiddleware = createEpicMiddleware();
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export function configureStore() {
  const store = createStore(
    transform,
    composeEnhancers(
      applyMiddleware(epicMiddleware),
    ),
  );

  epicMiddleware.run(transformEpic);

  return store;
}

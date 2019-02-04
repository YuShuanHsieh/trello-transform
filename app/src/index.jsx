import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import { configureStore } from './store';
import { App } from './App';
import './index.css';

const store = configureStore();

ReactDOM.render(
  <Provider store={store}>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.0/animate.min.css" />
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" />
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" />
    <App />
  </Provider>,
  // eslint-disable-next-line no-undef
  document.getElementById('root'),
);

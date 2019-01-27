import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import { menu } from './configuration';
import { Header } from './components/app/header/header';
import { Home } from './components/pages/home/home';
import { Result } from './components/pages/result/result';
import './App.css';

export class App extends Component {
  
  render() {
    return (
      <Router>
        <React.Fragment>
          <Header></Header>
          <Switch>
            <Route path={menu.home.path} component={Home} />
            <Route path={menu.result.path} component={Result} />
            <Route component={Home} />
          </Switch>
        </React.Fragment>
      </Router>
    );
  }
}

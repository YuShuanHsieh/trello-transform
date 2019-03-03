import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { MuiThemeProvider } from '@material-ui/core';

import { menu } from './configuration';
import { Header } from './components/app/header/header';
import { Home } from './components/pages/home/home';
import { Result } from './components/pages/result/result';
import { theme } from './configs/theme';
import style from './App.module.scss';

export function App() {
  return (
    <Router>
      <MuiThemeProvider theme={theme}>
        <div className={style.app}>
          <Header />
          <Switch>
            <Route path={menu.home.path} component={Home} />
            <Route path={menu.result.path} component={Result} />
            <Route component={Home} />
          </Switch>
        </div>
      </MuiThemeProvider>
    </Router>
  );
}

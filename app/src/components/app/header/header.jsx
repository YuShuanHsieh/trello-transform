import React from "react";
import { Home, Build, PieChart } from '@material-ui/icons';

import { menu } from '../../../configuration';
import style from './header.module.css';
import { LabelLink } from '../../shared/labelLink';

export class Header extends React.Component {
  render() {
    return (
      <div className={style.headerContainer}>
        <LabelLink path={menu.home.path} icon={Home} label={menu.home.label}></LabelLink>
        <LabelLink path={menu.result.path} icon={Build} label={menu.result.label}></LabelLink> 
        <LabelLink path={menu.analysis.path} icon={PieChart} label={menu.analysis.label}></LabelLink> 
      </div>
    )
  }
}
import React from "react";
import { NavLink } from 'react-router-dom';
import { IconLabel } from './iconLabel';
import style from './labelLink.module.css';

export class LabelLink extends React.Component {

  render() {
    const { path, component, ...rest } = this.props;

    return (
      <NavLink to={path} activeClassName={style.actived}>
        <IconLabel {...rest}></IconLabel>
      </NavLink>
    )
  }
}
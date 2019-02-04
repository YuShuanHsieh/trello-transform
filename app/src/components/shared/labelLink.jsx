import React from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';
import { IconLabel } from './iconLabel';
import style from './labelLink.module.css';

export function LabelLink({ path, component, ...rest }) {
  return (
    <NavLink to={path} activeClassName={style.actived}>
      <IconLabel {...rest} />
    </NavLink>
  );
}

LabelLink.propTypes = {
  path: PropTypes.string.isRequired,
  component: PropTypes.func.isRequired,
};

import React from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';
import { IconLabel } from './iconLabel';
import style from './labelLink.module.scss';

export function LabelLink({ path, ...rest }) {
  return (
    <NavLink to={path} activeClassName={style.actived}>
      <IconLabel {...rest} />
    </NavLink>
  );
}

LabelLink.propTypes = {
  path: PropTypes.string.isRequired,
};

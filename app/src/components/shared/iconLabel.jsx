import React from 'react';
import PropTypes from 'prop-types';
import style from './iconLabel.module.css';

export function IconLabel({ icon, label }) {
  const Icon = icon;

  return (
    <div className={style.iconLabelContainer}>
      <Icon />
      <span className={style.iconLabel}>{label}</span>
    </div>
  );
}
IconLabel.propTypes = {
  icon: PropTypes.func.isRequired,
  label: PropTypes.string.isRequired,
};

import React from "react";
import style from './iconLabel.module.css';

export class IconLabel extends React.Component {
  render() {
    const {icon, label} = this.props;
    const Icon = icon

    return (
      <div className={style.iconLabelContainer}>
        <Icon></Icon>
        <span className={style.iconLabel}>{label}</span>
      </div>
    )
  }
}
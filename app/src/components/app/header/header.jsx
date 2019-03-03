import React from 'react';
import { Home, Build } from '@material-ui/icons';

import { menu } from '../../../configuration';
import style from './header.module.scss';
import { LabelLink } from '../../shared/labelLink';

export function Header() {
  return (
    <div className={style.header}>
      <LabelLink path={menu.home.path} icon={Home} label={menu.home.label} />
      <LabelLink path={menu.result.path} icon={Build} label={menu.result.label} />
      {/* <LabelLink path={menu.analysis.path} icon={PieChart} label={menu.analysis.label} /> */}
    </div>
  );
}

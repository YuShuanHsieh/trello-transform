import React from 'react';

import { NavItem } from '../navLinks/navLinks';

import styles from './navMenu.module.scss';

interface NavMenuProps {
  items: NavItem[];
  target?: HTMLElement;
}

function NavMenu(props: NavMenuProps): JSX.Element {
  const left = props.target
    ? props.target.offsetLeft + (props.target.clientWidth / 2)
    : 0;
  return (
    <div className={styles.menu}>
      <div className={styles.menuArrow} style={{ left: `${left}px` }} />
      <ul>
        { props.items.map(item => (<li key={item.key}>{ item.displayName }</li>)) }
      </ul>
    </div>
  );
}

export default NavMenu;

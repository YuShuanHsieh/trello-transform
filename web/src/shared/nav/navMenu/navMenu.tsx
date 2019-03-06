import React from 'react';

import { NavItem } from '../navLinks/navLinks';

import styles from './navMenu.module.scss';

interface NavMenuProps {
  items: NavItem[];
  target?: HTMLElement;
}

function NavMenu(props: NavMenuProps): JSX.Element {
  return (
    <div className={styles.menu}>
      <ul>
        { props.items.map(item => (<li key={item.key}>{ item.displayName }</li>)) }
      </ul>
    </div>
  );
}

export default NavMenu;

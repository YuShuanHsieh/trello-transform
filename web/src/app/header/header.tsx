import React from 'react';

import NavLinks, { NavItem } from '../../shared/nav/navLinks/navLinks';
import styles from './header.module.scss';

const items: NavItem[] = [
  { key: 'home', displayName: 'Home' },
  { key: 'document', displayName: 'Document' },
  { key: 'blog', displayName: 'Blog' },
  { key: 'document-api', displayName: 'API', parent: 'document' },
  { key: 'document-libs', displayName: 'Library', parent: 'document' },
];

function Header(): JSX.Element {
  return (
    <header className={styles.header}>
      <h1 className={styles.headerLogo}>Trello Transform</h1>
      <NavLinks items={items} />
    </header>
  );
}

export default Header;

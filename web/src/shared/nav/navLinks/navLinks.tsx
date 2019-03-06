import React, { useMemo } from 'react';

import NavLink from './navLink';
import styles from './navLinks.module.scss';

export interface NavItem {
  key: string;
  displayName: string;
  route?: string;
  parent?: string;
}

export interface NavItemView extends NavItem {
  children: NavItemView[];
}

interface NavLinksProps {
  items: NavItem[];
  activeKey?: string;
}

function transformToView(items: NavItem[]): NavItemView[] {
  if (!items.length) return [];
  const maxTimes = 10;
  const result: NavItemView[] = [];
  const tempMap: Map<string, NavItemView> = new Map();
  let time = 0;
  let queue: NavItemView[] = items.map(item => ({ ...item, children: [] }));
  while (queue.length > 0 && time < maxTimes) {
    const temp: NavItemView[] = [];
    queue.forEach((item) => {
      if (!item.parent) {
        result.push(item);
      } else if (tempMap.has(item.parent)) {
        const parentView = tempMap.get(item.parent);
        if (parentView) parentView.children.push(item);
      } else {
        temp.push(item);
      }
      if (!tempMap.has(item.key)) {
        tempMap.set(item.key, item);
      }
    });
    queue = temp;
    time += 1;
  }
  if (time >= maxTimes) {
    console.error("Some parents are missing, please check items' key");
  }
  return result;
}

function NavLinks(props: NavLinksProps): JSX.Element {
  const views = useMemo(() => transformToView(props.items), [props.items]);
  return (
    <div className={styles.navLinks}>
      { views.map(item => (<NavLink key={item.key} item={item} />)) }
    </div>
  );
}

export default NavLinks;

import React, { ComponentType } from 'react';

import styles from './iconLink.module.scss';

interface IconLinkProps {
  icon?: ComponentType;
  text: string;
}

export default function IconLink(props: IconLinkProps): JSX.Element {
  const { icon: Icon, text } = props;
  return (
    <div className={styles.iconLink}>
      <span className={styles.iconLinkIcon}>{ Icon ? <Icon /> : null }</span>
      <span className={styles.iconLinkText}>{ text }</span>
    </div>
  );
}

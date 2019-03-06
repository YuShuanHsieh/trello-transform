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
      { Icon ? <span className={styles.iconLinkIcon}><Icon /></span> : null }
      <span className={styles.iconLinkText}>{ text }</span>
    </div>
  );
}

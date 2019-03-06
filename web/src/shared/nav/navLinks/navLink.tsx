import React, { useCallback, useState, SyntheticEvent } from 'react';
import { FaHeart } from 'react-icons/fa';

import IconLink from '../../iconLink/iconLink';
import { NavItemView } from './navLinks';
import NavMenu from '../navMenu/navMenu';

interface NavLinkProps {
  item: NavItemView;
}

function NavLink(props: NavLinkProps): JSX.Element {
  const [show, setShow] = useState(false);
  const [target, setTarget] = useState<HTMLElement|undefined>(undefined);
  const handleMouseEnter = useCallback((e: SyntheticEvent<HTMLDivElement>) => {
    if (props.item.children.length > 0) {
      setTarget(e.currentTarget);
      setShow(true);
    }
  }, []);

  const handleMouseLeave = useCallback(() => {
    setShow(false);
  }, []);

  return (
    <div
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
    >
      <IconLink text={props.item.displayName} icon={FaHeart} />
      { show ? (
        <div style={{ position: 'absolute', left: 0, width: '100%' }}>
          <NavMenu target={target} items={props.item.children} />
        </div>
      ) : null }
    </div>
  );
}

export default NavLink;

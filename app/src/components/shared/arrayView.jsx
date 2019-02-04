import React from 'react';
import PropTypes from 'prop-types';

import { ObjectView } from './objectView';

export function ArrayView({ array }) {
  const items = array.map((value, index) => {
    const id = `array-key-${index}`;
    return (
      <li key={id}>
        <ObjectView property={id} object={value} />
      </li>
    );
  });

  return (
    <ul>{items}</ul>
  );
}
ArrayView.propTypes = {
  array: PropTypes.arrayOf(PropTypes.object).isRequired,
};

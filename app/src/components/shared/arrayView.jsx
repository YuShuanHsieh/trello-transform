import React from 'react';

import { ObjectView } from './objectView';

export function ArrayView({array}) {
  const items = array.map((value, index) => {
    return (
      <li key={index}>
        <ObjectView property={index} object={value}></ObjectView>
      </li>
    )
  })

  return (
    <ul>{items}</ul>
  )
}
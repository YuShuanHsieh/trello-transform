import * as type from './types';

const initialState = {
  label: {},
  list: [],
  reference: []
}

export function transform(state = initialState, action) {
  switch(action.type) {
    case type.TRANSFORM_DATA_DONE:
      return {...state, ...action.payload};
    default:
      return state;
  }
} 
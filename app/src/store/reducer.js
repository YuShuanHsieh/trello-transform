import * as type from './types';

const initialState = {};

export function transform(state = initialState, action) {
  switch (action.type) {
    case type.TRANSFORM_DATA_DONE:
      return ({ ...state, ...action.payload });
    default:
      return state;
  }
}

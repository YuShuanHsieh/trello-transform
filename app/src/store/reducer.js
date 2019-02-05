import * as type from './types';

const initialState = {
  result: {},
  uploaded: false,
};

export function transform(state = initialState, action) {
  switch (action.type) {
    case type.TRANSFORM_DATA_DONE:
      return ({ ...state, result: { ...action.payload }, uploaded: true });
    default:
      return state;
  }
}

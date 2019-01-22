import * as type from './types';

export const transformData = (options= {}) => ({type: type.TRANSFORM_DATA, payload: options})

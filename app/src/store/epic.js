import { catchError, mergeMap, map } from 'rxjs/operators';
import { ofType } from 'redux-observable';

import { TRANSFORM_DATA, TRANSFORM_DATA_DONE } from './types';
import { AppService } from '../services/app.service';

export const transformEpic = action$ => action$.pipe(
  ofType(TRANSFORM_DATA),
  mergeMap(
    action => AppService.transform(action.payload).pipe(
      map(res => ({ type: TRANSFORM_DATA_DONE, payload: res.response })),
      catchError((err) => {
        console.debug(err);
        return {};
      }),
    ),
  ),
);

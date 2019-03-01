/* eslint-disable no-undef */
import { ajax } from 'rxjs/ajax';

export class AppService {
  static transform({ file, list }) {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('list', list);
    return ajax.post('/transform', formData);
  }

  static isChartAvailable(object) {
    return Object.keys(object).every(key => typeof object[key] === 'number');
  }
}

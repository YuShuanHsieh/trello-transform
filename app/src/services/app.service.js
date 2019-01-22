import { ajax } from 'rxjs/ajax';

export class AppService {
  static transform(payload) {
    return ajax.post("/transform", payload)
  }
}
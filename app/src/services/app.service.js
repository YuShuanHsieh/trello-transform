import { ajax } from 'rxjs/ajax';

export class AppService {
  static transform(file) {
    const formData = new FormData();
    formData.append('file', file);
    return ajax.post("/transform", formData)
  }
}
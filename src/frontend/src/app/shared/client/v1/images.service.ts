import { Injectable } from '@angular/core';
import { HttpHeaders, HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { throwError } from 'rxjs';

@Injectable()
export class ImagesService {
  headers = new HttpHeaders({'Content-type': 'application/json'});
  options = {'headers': this.headers};
  constructor(private http: HttpClient) { }
  listPage(page: any): Observable<any> {
    let params = new HttpParams();
    params = params.set('pageNo', page.currentPage + '');
    params = params.set('pageSize', page.pageSize + '');
    params = params.set('versionName', page.versionName + '');
    return this.http
      .get(`/api/v1/images`, {params: params})
      .catch(error => throwError(error));
  }
}

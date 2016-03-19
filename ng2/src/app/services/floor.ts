import {Injectable} from 'angular2/core';
import {Http} from 'angular2/http';

@Injectable()
export class Floor {
  constructor(public http: Http) {
  }

  getData() {
    console.log('Floor#getData(): Get Data');
    var res = this.http.get('http://192.168.0.101/board')
      .map(res => res.json());
    console.log(res);
    return res;

  }

}

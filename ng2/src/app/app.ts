/*
 * Angular 2 decorators and services
 */
import {Component} from 'angular2/core';
import {RouteConfig, Router} from 'angular2/router';

import {Floor} from './services/floor';
//import {Home} from './home/home';

/*
 * App Component
 * Top Level Component
 */
@Component({
  selector: 'app',
  pipes: [],
  providers: [
    Floor
  ],
  directives: [],
  styles: [require('./app.scss')],
  template: require('./app.html')
})
export class App {
  classMap: string = "";
  mode: string = ""; //visitor, homelord, killer
  tour: string = ""; //visitor, homelord

  constructor(public floor: Floor) {
    this.classMap = "";//{ 'class1': true, 'class2': false };
    var that = this;
    var i: number = 1;

    window.setInterval(() => {
      var data = floor.getData().subscribe(
        data => console.log,
        err => console.error,
        () => console.log('Authentication Complete')
        )
    }, 1000)
  }

  updateInRoom(i: number) {
    this.classMap = 'in-room' + i;
  }
}

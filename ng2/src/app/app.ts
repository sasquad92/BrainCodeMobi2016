/*
 * Angular 2 decorators and services
 */
import {Component} from 'angular2/core';
import {RouteConfig, Router} from 'angular2/router';

import {Home} from './home/home';

/*
 * App Component
 * Top Level Component
 */
@Component({
  selector: 'app',
  pipes: [],
  providers: [],
  directives: [],
  styles: [require('./app.scss')],
  template: require('./app.html')
})
/*@RouteConfig([
  { path: '/',      name: 'Index', component: Home, useAsDefault: true },
  { path: '/home',  name: 'Home',  component: Home },
  // Async load a component using Webpack's require with es6-promise-loader and webpack `require`
  { path: '/about', name: 'About', loader: () => require('es6-promise!./about/about')('About') },
])*/
export class App {
  /*angularclassLogo = 'assets/img/angularclass-avatar.png';
  name = 'Angular 2 Webpack Starter';
  url = 'https://twitter.com/AngularClass';*/
  classMap:string = "" //{ [key: string]: boolean; }

  constructor() {
    this.classMap = "";//{ 'class1': true, 'class2': false };

    var that = this;
    var i: number = 1;
    setInterval(() => {
      i++
      that.classMap = 'in-room' + i;
      if(i==11) i =0;
    }, 1000)
  }
}

/*
 * Please review the https://github.com/AngularClass/angular2-examples/ repo for
 * more angular app examples that you may copy/paste
 * (The examples may not be updated as quickly. Please open an issue on github for us to update it)
 * For help or questions please contact us at @AngularClass on twitter
 * or our chat on Slack at https://AngularClass.com/slack-join
 */

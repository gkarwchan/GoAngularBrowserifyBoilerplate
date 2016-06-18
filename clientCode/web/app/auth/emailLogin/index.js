'use strict';

require('angular');
require('angular-resource');
require('angular-ui-router');

module.exports =
   angular.module('app.auth.email',  ['ui.router', 'ngResource'])
   .config(require('./config'))
   .controller('signupCtrl', require('./signup-controller'))
   .controller('loginCtrl', require('./login-controller'))
   .controller('singedupCtrl', require('./signedup-controller'))
   .factory('User', require('./user-factory'))
   .directive('compareTo', require('./compareTo-directive'))
   .name;

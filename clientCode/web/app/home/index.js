'use strict';

require('angular');
require('angular-ui-router');

module.exports = angular.module('app.home', ['ui.router'])
  .config(require('./config'))
  .controller('HomeCtrl', require('./controller'))
  .name;

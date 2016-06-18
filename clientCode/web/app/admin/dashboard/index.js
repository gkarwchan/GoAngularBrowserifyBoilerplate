'use strict';

require('angular');
require('angular-resource');
require('angular-ui-router');


module.exports =
  angular.module('app.admin.dashboard', [
      'ui.router',
      'ngResource'
    ])
    .config(require('./config'))
    .name;

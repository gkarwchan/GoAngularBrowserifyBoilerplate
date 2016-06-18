'use strict';

require('angular');
require('angular-resource');
require('angular-ui-router');


module.exports =
  angular.module('app.admin', [
      'ui.router',
      require('./dashboard'),
      'ngResource'
    ])
    .name;

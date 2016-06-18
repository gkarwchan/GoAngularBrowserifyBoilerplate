'use strict';

require('angular');
require('angular-resource');
require('angular-ui-router');


module.exports =
  angular.module('app.auth', [
      'ui.router',
      require('./emailLogin'),
      require('./oauth'),
      'ngResource'
    ])
    .config(require('./config'))
    .run(require('./run'))
    .name;

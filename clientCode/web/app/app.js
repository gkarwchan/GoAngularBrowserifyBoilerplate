'use strict';

require('angular');
require('angular-touch');
require('angular-ui-router');
require('angular-bootstrap');

angular.module('app', [
      'ui.router',
      'ngTouch',
      'ui.bootstrap',
      require('./auth'),
      require('./home'),
      require('./common'),
      require('./admin')
    ])
    .config(require('./config'))
    .run(require('./run'))
    .controller('NavbarCtrl', require('./navbar-controller'));

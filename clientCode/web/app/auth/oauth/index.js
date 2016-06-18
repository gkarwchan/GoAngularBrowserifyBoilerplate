'use strict';

require('angular');
require('angular-resource');
require('angular-ui-router');

module.exports =
   angular.module('app.auth.oauth',  ['ui.router', 'ngResource'])
   .factory('auth_Providers', require('./provider-factory'))
   .constant('auth_popupConstants', require('./popup-constants'))
   .service('auth_oauthService', require('./oauth-service'))
   .service('popupService', require('./popup-service'))
   .factory('pubsub', require('./pubsub-factory'))
   .name;

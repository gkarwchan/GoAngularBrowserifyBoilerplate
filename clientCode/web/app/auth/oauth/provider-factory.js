/*jslint node: true */
'use strict';

module.exports = /*@ngInject*/
  function($resource) {
    return $resource('/auth/:providerId', {providerId: '@providerId'});
  };

'use strict';

module.exports = /*@ngInject*/
  function($urlRouterProvider) {
    $urlRouterProvider
      .otherwise('/');
  };

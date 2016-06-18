'use strict';

module.exports = /*@ngInject*/
  function($resource) {
    return $resource('/user/:userId', {userId: '@userId'}, {
    	'login' : { method: "POST", url: '/login'}
    });
  };

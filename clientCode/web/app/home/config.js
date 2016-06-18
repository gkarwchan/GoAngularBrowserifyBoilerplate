'use strict';

module.exports = /*@ngInject*/
  function ($stateProvider) {
    $stateProvider
      .state('home', {
        url: '/',
        parent: 'public',
        views: {
          '': {
            templateUrl: 'app/home/index.html',
            controller: 'HomeCtrl'
          },
          'toolbar' : {
            template: ''
          }  
        },
        data: {
          title: 'Acne Expert',
          displayName: 'Home'
        }
      });
  };

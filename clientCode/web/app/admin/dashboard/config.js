'use strict';

module.exports = /*@ngInject*/
  function ($stateProvider) {
    $stateProvider
      .state('user.admin.dashboard', {
        url: '/',
        views: {
          '': {
            templateUrl: 'app/admin/dashboard/index.html'
          },
          'toolbar' : {
            template: ''
          }  
        },
        onEnter: /*@ngInject*/ function($state) {
                console.log('admin enter');
                console.log($state.current.data);
            },
        data: {
          title: 'Acne Expert',
          displayName: 'Home'
        }
      });
  };

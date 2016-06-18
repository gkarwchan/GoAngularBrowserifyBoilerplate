'use strict';
module.exports = /*@ngInject*/
    function($stateProvider) {
        $stateProvider.state('public', {
            abstract: true,
            views: {
              '': {
                  template: '<ui-view/>'
              },
              'toolbar': {
                  template: ''
              }  
            },
            data: {
                authenticate: false
            }
        });
        // abstart logged in user (authenticated)
        $stateProvider.state('user', {
            abstract: true,
            views: {
                'toolbar': {
                    templateUrl: 'app/auth/navbarUser.html'//,
                    //controller: ''
                },
                '': {
                    template: '<ui-view/>'
                }
            },
            params: {
                user: null
            },
            onEnter: /*@ngInject*/ function($stateParams, $state) {
                console.log('enter user');
                console.log(arguments);
                console.log($stateParams);
                $state.current.data.user = $stateParams.user;
            },
            data: {
                authenticate: true,
                user: null
            }
        });
        // regular read only user
        $stateProvider.state('user.regular', {
            abstract: true,
            template: '<ui-view/>',
            data: {
                accessLevel: 'readonly'
            }
        });
        
        // editor user
        $stateProvider.state('user.editor', {
            url: '/editor',
            abstract: true,
            template: '<ui-view/>',
            data: {
                accessLevel: 'editor'
            }
        });
        
        // Admin routes
        $stateProvider.state('user.admin', {
            url: '/admin',
            abstract: true,
            template: '<ui-view/>',
            data: {
                accessLevel: 'admin'
            }
        });
  };

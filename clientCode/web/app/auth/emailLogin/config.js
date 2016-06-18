'use strict';

module.exports = /*@ngInject*/
    function($stateProvider) {
        $stateProvider
            .state('public.login', {
                url: '/login',
                templateUrl: 'app/auth/emailLogin/login.html',
                controller: 'loginCtrl',
                data : {
                  title: 'Login page',
                  displayName: 'Login'
                }
            })
            .state('public.signedup', {
              url: '/signedupSuccessful/:email',
              templateUrl: 'app/auth/emailLogin/signedup.html',
              controller: 'singedupCtrl'
            })
            .state('public.signup', {
              url: '/signup',
              templateUrl: 'app/auth/emailLogin/signup.html',
              controller: 'signupCtrl',
              data : {
                title: 'Signup page',
                displayName: 'Signup'
              }
            });
    };

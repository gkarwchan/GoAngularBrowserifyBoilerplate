'use strict';

module.exports = /*@ngInject*/
    function($log, $rootScope, $state) {
        $rootScope.$on('$stateChangeStart', function(event, toState, toParams, fromState, fromParams) {
        //   console.log(event);
        //   console.log(toState);
        //   console.log(toParams);
        });
    };

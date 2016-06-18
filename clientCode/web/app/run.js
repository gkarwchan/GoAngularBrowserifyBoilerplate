'use strict';

module.exports = /*@ngInject*/
    function($rootScope, $state, $stateParams) {
        // add references to $state and $stateParams to the $rootScope so we can access them from any scope
        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
    };

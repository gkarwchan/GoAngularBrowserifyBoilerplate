'use strict';

module.exports = /*@ngInject*/
   function($scope, $state, $stateParams) {
     $scope.email = $stateParams.email;
   };

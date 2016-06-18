'use strict';

module.exports = /*@ngInject*/
  function($scope, auth_Providers, $window, auth_oauthService, $log, $state) {
    $scope.providers = auth_Providers.query();
    $scope.authenticate = function(provider) {
      auth_oauthService.authenticate(provider.url, provider.name)
        .then(function(response) {
          $log.debug('yes....');
          $log.debug(response);
          $state.go('user.admin.dashboard', {user: response});
        })
        .catch(function(error) {
          $log.debug('oh error');
          $log.debug(error);
        });
      
    };
  };

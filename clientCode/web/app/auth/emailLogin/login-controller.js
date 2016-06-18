'use strict';

module.exports = /*@ngInject*/
  function($scope, User, ngToast) {

    var onSuccess = function(data) {

    };

    var onError = function(data) {
      ngToast.create({
             className: 'danger',
             content: 'Cannot create account for email : <strong>' + $scope.email + '</strong>. <br/> Error: ' + data.data
         });
    };

    $scope.login = function() {
      var usr = new User({provider: 'email'});
      usr.email = $scope.email;
      usr.password = $scope.password;
	    usr.$login(onSuccess, onError);
    };
  };

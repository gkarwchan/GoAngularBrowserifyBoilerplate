'use strict';


module.exports = /*@ngInject*/
   function($scope, User, ngToast, $state) {
     var onSuccess = function(data) {
       ngToast.create({
              className: 'success',
              content: 'A confirmation was sent to : <strong>' + $scope.email + '</strong>. Please follow up'
          });
        $state.go('public.signedup', {email: $scope.email});
     };

     var onError = function(data) {
       ngToast.create({
              className: 'danger',
              content: 'Cannot create account for email : <strong>' + $scope.email + '</strong>. <br/> Error: ' + data.data
          });
     };

     $scope.signup = function() {
       var usr = new User({provider: 'email'});
       usr.email = $scope.email;
       usr.password = $scope.password;
       usr.$save(onSuccess, onError);
     };
   };

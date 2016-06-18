'use strict';

module.exports = /*@ngInject*/
function($q, popupService, $window, $log, $rootScope) {
    var service = {
        deferred:null,
        loginSuccess: null,
        popupClosed: null,
        listen: function() {
            this.loginSuccess = angular.bind(this, this.handleLoginSuccess);
            this.popupClosed = angular.bind(this, this.handlePopupClosed);
            $window.addEventListener('message', this.loginSuccess, false);
            $rootScope.$on('window-closed', this.popupClosed);
        },
        handleLoginSuccess: function(e) {
            if (e.data.message && e.data.message === 'loginSuccess') {
                this.deferred.resolve(e.data.userLogged);
            }
        },
        handlePopupClosed: function() {
            this.deferred.reject('the popup was closed');
        },
        authenticate: function(url, provider) {
            this.deferred = $q.defer();
            this.listen();
            
            popupService.open(url, provider);
            
            
            return this.deferred.promise;
        }
    };

    return service;
};

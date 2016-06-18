'use strict';

function stringifyOptions(options) {
        var parts = [];
        angular.forEach(options, function(value, key) {
            parts.push(key + '=' + value);
        });
        return parts.join(',');
    }

    function prepareOptions (options, $window) {
        var width = options.width || 500;
        var height = options.height || 500;
        return angular.extend({
            toolbar: 0,
            menubar: 0,
            location: 1,
            status: 1,
            directories: 0,
            scrollbars: 1,
            resizable: 0,
            width: width,
            height: height,
            left: $window.screenX + (($window.outerWidth - width) / 2),
            top: $window.screenY + (($window.outerHeight - height) / 2.5)
        }, options);
    }
    

module.exports = /*@ngInject*/
function($log, auth_popupConstants, $window, $interval, $rootScope) {

    var service = {
        window: null,
        interval: null,
        cancelWatcher: function(e) {
            if (e.data.message && e.data.message === 'loginSuccess') {
                $interval.cancel(service.interval);
            }
        }, 
        open: function(url, provider) {
            $window.addEventListener('message', this.cancelWatcher);
            var optionsString = stringifyOptions(prepareOptions(auth_popupConstants[provider], $window));
            this.window = $window.open(url, '_blank', optionsString);
            if (!this.window) {
                $interval.cancel(this.interval);
                return false;
            } 
            this.obeservePopup();
            return true;
        },
        obeservePopup: function() {
            this.interval = $interval(function() {
                 if (service.window.closed) {
                     $interval.cancel(service.interval);
                     $rootScope.$emit('window-closed');
                 }
                 service.window.postMessage('requestLogin', '*');
            }, 500);
        }
    };

    return service;
};

'use strict';

module.exports = /*@ngInject*/
function ($rootScope) {
  return {
      subscribe: function(callback) {
          $rootScope.$on('notify-service-event', callback);
      }, 
      notify: function() {
          $rootScope.$emit('notify-service-event');
      }
  };  
};
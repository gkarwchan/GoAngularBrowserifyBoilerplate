'use strict';

module.exports = /*@ngInject*/
  function ($scope) {
    var mq = window.matchMedia( "(min-width: 500px)" );
    $scope.navbarCollapsed = !mq.matches;
  };

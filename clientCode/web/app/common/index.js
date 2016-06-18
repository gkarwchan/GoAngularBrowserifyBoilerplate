'use strict';

require('angular-animate');
require('angular-sanitize');
require('ngToast');

module.exports =  /*@ngInject*/
    angular
    .module('app.common', [
        'ngAnimate',
        'ngToast',
        'ngSanitize'
    ])
    .config(require('./config'))
    .name;

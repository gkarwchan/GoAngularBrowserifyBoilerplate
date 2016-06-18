'use strict';

module.exports = /*@ngInject*/
    function (ngToastProvider) {
        ngToastProvider.configure({
            animation: 'fade',
            verticalPosition: 'bottom',
            horizontalPosition: 'left',
            timeout: 5000,
            maxNumber: 3
        });
    };

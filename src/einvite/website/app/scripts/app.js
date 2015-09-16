'use strict';

angular.module('websiteApp', ['$strap.directives'])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/newevent', {
        templateUrl: 'views/newevent.html',
        controller: 'NewEventCtrl'
      })
      .when('/event/:eventid', {
        templateUrl: 'views/eventdetails.html',
        controller: 'EventCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
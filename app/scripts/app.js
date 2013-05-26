'use strict';

angular.module('govotingApp', ['ngResource'])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/vote/:id', {
        templateUrl: 'views/vote.html',
        controller: 'VoteCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });

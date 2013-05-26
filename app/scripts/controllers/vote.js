'use strict';

angular.module('govotingApp')
  .controller('VoteCtrl', function ($scope, $resource, $routeParams) {
   	var Vote = $resource('/vote/:voteId', {voteId: '@id'});
   	var vote = Vote.get({voteId: $routeParams.id}, function() {
   		$scope.vote = vote;
   		console.log(vote);
   	});
  });

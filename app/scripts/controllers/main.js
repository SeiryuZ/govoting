'use strict';

angular.module('govotingApp')
  .controller('MainCtrl', function ($scope, $resource) {
   	var Votes = $resource('/vote');
   	var Vote = $resource('vote');

   	var votes = Votes.query(function() {
   		$scope.votes = votes;
   	});


   	$scope.showForm = function () {
   		$('.create-vote-form').removeClass('hide');
   	}
   	$scope.saveVote = function (vote) {
   		$('.create-vote-form').addClass('hide');
   		var newVote = new Vote(vote);
   		newVote.$save(
   			function (vote) {
   				console.log ($scope.votes)
   				$scope.votes.unshift(vote)
   				console.log ($scope.votes)

   			},
   			function (error) {
   				console.log(error)
   			}
   		);
   	}

  });

'use strict';

angular.module('websiteApp')
  .controller('ActivitiesCtrl', function ($scope) {
  	$scope.showComplete = true;

    $scope.activities = [
    	{ name: 'activity 1', rating: 10 },
    	{ name: 'activity 2', rating: 9 },
    	{ name: 'activity 3', rating: 8 },
    	{ name: 'activity 4', rating: 7 },
    	{ name: 'activity 5', rating: 6 },
    	{ name: 'activity 6', rating: 5 },
    	{ name: 'activity 7', rating: 4 },
    	{ name: 'activity 8', rating: 3 },
    	{ name: 'activity 9', rating: 2 },
    	{ name: 'activity 10', rating: 1 },
    ];
  });
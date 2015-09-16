'use strict';

angular.module('websiteApp')
  .controller('LocationsCtrl', function ($scope) {
  	$scope.showComplete = true;

    $scope.locations = [
    	{ name: 'location 1', rating: 10 },
    	{ name: 'location 2', rating: 9 },
    	{ name: 'location 3', rating: 8 },
    	{ name: 'location 4', rating: 7 },
    	{ name: 'location 5', rating: 6 },
    	{ name: 'location 6', rating: 5 },
    	{ name: 'location 7', rating: 4 },
    	{ name: 'location 8', rating: 3 },
    	{ name: 'location 9', rating: 2 },
    	{ name: 'location 10', rating: 1 },
    ];
  });

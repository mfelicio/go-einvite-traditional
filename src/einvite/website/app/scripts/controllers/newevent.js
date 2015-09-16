'use strict';

angular.module('websiteApp')
  .controller('NewEventCtrl', function ($scope) {
    
    $scope.event = {
        type: "",

        location: "",

        date: "",

        time: "",

        people: [],

        activities: []
    }
    
    $scope.newPerson = "";

    $scope.newActivity = "";

    $scope.addPerson = function(){
    	if($scope.newPerson){
    		$scope.event.people.push($scope.newPerson);
    	}
    };

    $scope.addActivity = function(){
    	if($scope.newActivity){
    		$scope.event.activities.push($scope.newActivity);
    	}
    };

    $scope.createEvent = function(newEvent){
        console.log('submiting form');
    };

  });
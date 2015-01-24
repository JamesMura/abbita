var application = angular
.module('abbita', ['angularFileUpload', 'ui.router'])
.controller('NewSessionController', ['$scope', 'FileUploader', function($scope, FileUploader) {
	$scope.uploader = new FileUploader();
	$scope.actions = ['Convert to wav', 'Convert to mp4'];
}]);

application.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider){
    $urlRouterProvider.otherwise("/");
    $stateProvider
    .state('home', {
        url: "/",
        templateUrl: 'templates/home.html',
        controller: 'NewSessionController'
    })
}])	

application.filter('bytes', function() {
	return function(bytes, precision) {
		if (isNaN(parseFloat(bytes)) || !isFinite(bytes)) return '-';
		if (typeof precision === 'undefined') precision = 1;
		var units = ['bytes', 'kB', 'MB', 'GB', 'TB', 'PB'],
			number = Math.floor(Math.log(bytes) / Math.log(1024));
		return (bytes / Math.pow(1024, Math.floor(number))).toFixed(precision) +  ' ' + units[number];
	}
});
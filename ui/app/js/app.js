
var cansApp = angular.module('cans', [
  'ui.bootstrap',
  'ngRoute',
  'cansControllers'
]);

cansApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/settings', {
        templateUrl: 'partials/settings.html',
        controller: 'SettingsCtrl'
      }).
      when('/artists', {
        templateUrl: 'partials/artists.html',
        controller: 'ArtistListCtrl'
      }).
      when('/releases', {
        templateUrl: 'partials/releases.html',
        controller: 'ReleaseListCtrl'
      }).
      otherwise({
        redirectTo: '/'
      })
  }]);

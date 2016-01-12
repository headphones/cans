var cansControllers = angular.module('cansControllers', []);

cansControllers.controller('IndexCtrl', function ($scope) {
  $scope.dashboard = {};
});

cansControllers.controller('ArtistListCtrl', ['$scope', '$http',
  function ($scope, $http) {
    $http.get('/mocks/artists.json').success(function(data) {
      $scope.artists = data;
    });

    $scope.orderProp = 'name';
  }]);

cansControllers.controller('ReleaseListCtrl', ['$scope', '$http',
  function ($scope, $http) {
    $http.get('/mocks/releases.json').success(function(data) {
      $scope.artists = data;
    });

    $scope.orderProp = 'name';
  }]);

cansControllers.controller('SettingsCtrl', function ($scope) {
  $scope.dashboard = {};
});

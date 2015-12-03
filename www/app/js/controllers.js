angular.module('demoapp', ['ngMdIcons']);

var phonecatApp = angular.module('cans', ['ui.bootstrap']);

phonecatApp.controller('indexController', function ($scope) {
  $scope.artists = [
    {'name': 'Black Keys',
     'snippet': 'The Black Keys are an American rock duo'},
    {'name': 'Deafheaven',
     'snippet': 'Originally formed as a duo in 2010, Deafheaven are an experimental black metal band from San Francisco, California.'},
    {'name': 'Aphex Twin',
     'snippet': 'Richard David James, best known by his alias Aphex Twin, is an Irish-born British electronic musician and composer. He is also the co-founder of Rephlex Records with Grant Wilson-Claridge. '}
  ];
});
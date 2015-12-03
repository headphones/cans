angular.module('demoapp', ['ngMdIcons']);

var phonecatApp = angular.module('cans', ['ui.bootstrap']);

phonecatApp.controller('indexController', function ($scope) {
  $scope.artists = [
    {'name': 'Black Keys',
     'snippet': 'The Black Keys are an American rock duo'},
    {'name': 'Deafheaven',
     'snippet': 'Originally formed as a duo in 2010, Deafheaven are an experimental black metal band from San Francisco, California.'},
    {'name': 'Aphex Twin',
     'snippet': 'Richard David James, best known by his alias Aphex Twin, is an Irish-born British electronic musician and composer. He is also the co-founder of Rephlex Records with Grant Wilson-Claridge. '},
    {'name': 'My Bloody Valentine',
     'snippet': 'Like the Velvet Underground, Sonic Youth, and the Jesus and Mary Chain before them, My Bloody Valentine redefined what noise meant within the context of pop songwriting. Led by guitarist Kevin Shields, the group released several EPs in the mid-\'80s before recording the era-defining Isn\'t Anything in 1988, a record that merged lilting, ethereal melodies of the Cocteau Twins with crushingly loud, shimmering distortion. Though My Bloody Valentine rejected rock & roll conventions, they didn\'t subscribe to the precious tendencies of anti-rock art-pop bands. Instead, they rode crashing waves of white noise to unpredictable conclusions, particularly since their noise wasn\'t paralyzing like the typical avant-garde noise rock band: it was translucent, glimmering, and beautiful. Shields was a perfectionist, especially when it came to recording, as much of My Bloody Valentine\'s sound was conceived within the studio itself. Nevertheless, the band was known as a formidable live act, even though they rarely moved, or even looked at the audience, while they were on-stage. Their notorious lack of movement was branded \"shoegazing\" by the British music press, and soon there were legions of other shoegazers -- Ride, Lush, the Boo Radleys, Chapterhouse, Slowdive -- that, along with the rolling dance-influenced Madchester scene, dominated British indie rock of the late \'80s and early \'90s. As shoegazing reached its peak in 1991, My Bloody Valentine released Loveless, which broke new sonic ground and was hailed as a masterpiece. Though the band was poised for a popular breakthrough, it disappeared into the studio and didn\'t emerge over the next five years, leaving behind a legacy that proved profoundly influential in the direction of \'90s alternative rock.'}

  ];
});
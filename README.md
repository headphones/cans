# Cans [![Circle CI](https://circleci.com/gh/headphones/cans.svg?style=svg)](https://circleci.com/gh/headphones/cans)

This is a complete rewrite of [Headphones](//github.com/rembo10/headphones/)
in [golang](https://golang.org/). If you are interested in contributing, check out the
[issue tracker](https://github.com/headphones/cans/issues) and file a bug, submit a pull 
request, or just star the project to increase it's visibility. Thanks for taking a look!

# Objectives

The initial goals of this rewrite are (beyond a skeleton set of Headphones' features):

1) First and foremost: a faster, more responsive app
2) Lower resource usage and fewer external deps so we can support resource-constrained devices
3) Angular / Bootstrap UI so we can do more stuff client-side to make the app more 
  responsive and work on a wider variety of devices and viewports.
4) Better search
5) Better data model that supports EPs, singles and other release types.

*NOTE: NBZ support is not currently a priority* 

The long-term goals are to implement the full original feature set, and expand it with:
 
1) Online DRM-free music stores (Bandcamp, etc.)
2) Support multiple database backends (Musicbrainz, Discogs, ???)
3) ???

# Building

To build cans yourself, you need to have [go](https://golang.org/) set up on your 
machine. `go get github.com/headphones/cans` and `go run github.com/headphones/cans` 
should get you up and running from there. If you'd like to run within a Docker container, 
configurations are provided for you;

    docker run -Pit `./docker-container.sh | awk 'END {print $3}'`


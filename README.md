# Cans [![Circle CI](https://circleci.com/gh/headphones/cans.svg?style=svg)](https://circleci.com/gh/headphones/cans)

This is a complete rewrite of [Headphones](//github.com/rembo10/headphones/)
in [golang](https://golang.org/).

# Building

To build cans yourself, you need to have [go](https://golang.org/) set up on your 
machine. `go get github.com/headphones/cans` and `go run github.com/headphones/cans` 
should get you up and running from there. If you'd like to run within a Docker container, 
configurations are provided for you;

    docker run -Pit `./docker-container.sh | awk 'END {print $3}'`

#!/bin/bash

echo "compiling go code..."
export CGO_ENABLED=0
export GOOS=linux
go build -o cans .

echo "building container..."
docker build .

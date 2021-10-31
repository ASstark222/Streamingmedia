#! /bin/bash

# Build web and other server

cd ~/workspace/src/Streamingmedia/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ~/workspace/src/Streamingmedia/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/workspace/src/Streamingmedia/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin.streamserver

cd ~/workspace/src/Streamingmedia/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web
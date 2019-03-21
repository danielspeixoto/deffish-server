#!/usr/bin/env bash
export PORT=5000
export MONGO_CONNECTION=mongodb://localhost:27017
export MONGO_DB_NAME=deffishtest

docker rm --force -v deffishdb
docker run -d -p 27017:27017 --name deffishdb mongo && \
go build main.go && \
 ./main
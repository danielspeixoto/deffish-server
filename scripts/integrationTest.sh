#!/usr/bin/env bash

docker rm --force -v deffishdb
docker run -d -p 27017:27017 --name deffishdb mongo
sleep 20
export PORT=5000
export MONGO_CONNECTION=mongodb://localhost:27017
export MONGO_DB_NAME=deffishtest

go build main.go
./main & RUNNING_PID=$!
sleep 30
go test -count=1 ./test/...
docker rm --force -v deffishdb
kill ${RUNNING_PID}




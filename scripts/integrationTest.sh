#!/usr/bin/env bash

docker rm --force -v deffishdb
docker run -d -p 27017:27017 --name deffishdb mongo
export PORT=5000
export MONGO_CONNECTION=mongodb://localhost:27017
export MONGO_DB_NAME=deffishtest
go build main.go &&
./main & RUNNING_PID=$!
sleep 5./
jest --runInBand test
sleep 5
kill ${RUNNING_PID}




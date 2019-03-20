#!/usr/bin/env bash

docker rm --force -v deffishdb
docker run -d -p 27017:27017 --name deffishdb mongo && \
sleep 10 && \
./scripts/local.sh & \
sleep 20 && \
go test -count=1 ./test/... && \
docker rm --force -v deffishdb && \
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT


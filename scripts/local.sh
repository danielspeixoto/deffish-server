#!/usr/bin/env bash
export PORT=5000
export MONGO_CONNECTION="mongodb+srv://server:Skw1ZwnJXAUx5YEb@cluster0-lf760.mongodb.net/test&retryWrites=true"
export MONGO_DB_NAME="test"

#docker rm --force -v deffishdb
#docker run -d -p 27017:27017 --name deffishdb mongo && \
go build main.go && \
 ./main
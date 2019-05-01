#!/usr/bin/env bash
export PORT=5000
export MONGO_CONNECTION="mongodb+srv://danielspeixoto:Bornnov16@cluster0-lf760.mongodb.net/heroku_wn1s1nxv&retryWrites=true"
export MONGO_DB_NAME="heroku_wn1s1nxv"

#docker rm --force -v deffishdb
#docker run -d -p 27017:27017 --name deffishdb mongo && \
go build main.go && \
 ./main
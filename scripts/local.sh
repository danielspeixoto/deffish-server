export PORT=5000
export MONGO_CONNECTION=mongodb://localhost:27017
export MONGO_DB_NAME=deffishtest

go build main.go && \
 ./main
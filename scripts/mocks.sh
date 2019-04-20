#!/usr/bin/env bash

DIRS=$(find src/boundary/ -maxdepth 1 -mindepth 1 -type d)
for DIR in $DIRS
do
    rm $DIR/Mocks.go
    IFS="/" read -ra foo <<< "$DIR"
    echo ${foo[3]}
    mockgen -source=$DIR/Boundary.go -destination=${DIR}/Mocks.go -package=${foo[3]}
done


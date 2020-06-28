#!/bin/zsh
TIMESTAMP=`date +%T`
FILE_PATH="./build"
FILE_NAME="interface-${TIMESTAMP}"
OUTPUT="${FILE_PATH}/${FILE_NAME}"
GOARM=7 GOARCH=arm GOOS=linux go build -o $OUTPUT
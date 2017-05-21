#!/bin/bash

set -e

if [[ -z ${GOOS} ]]; then GOOS=linux; fi
docker run --rm -i \
       -v ${HOME}/.ssh:/root/.ssh \
       -v ${HOME}/.gitconfig:/root/.gitconfig \
       -v ${PWD}:/go/src/github.com/tskinn/sqs-daemon \
       -w /go/src/github.com/tskinn/sqs-daemon \
       golang:1.7 \
       sh -c " go get -u -d -v && GOOS=${GOOS} go build -o bin/${GOOS}/sqs-deamon ."

if [ $? -eq 0 ]; then
    echo "${GOOS} binary successfully built at bin/${GOOS}/sqs-daemon"
fi

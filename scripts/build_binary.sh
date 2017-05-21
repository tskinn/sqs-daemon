#!/bin/bash

set -e

if [[ -z ${GOOS} ]]; then GOOS=linux; fi
docker run --rm -i \
       -v ${HOME}/.ssh:/root/.ssh \
       -v ${HOME}/.gitconfig:/root/.gitconfig \
       -v ${PWD}:/go/src/github.com/tskinn/sqs-daemon \
       -w /go/src/github.com/tskinn/sqs-daemon \
       golang:1.7 \
       sh -c "echo yes; cd src && go get -u -d -v && GOOS=${GOOS} go build -o ../bin/${GOOS}/sqsd ."

if [ $? -eq 0 ]; then
    echo "${GOOS} binary successfully built at bin/${GOOS}/sqsd"
fi

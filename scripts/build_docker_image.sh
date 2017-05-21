#!/bin/bash

. ./scripts/utils.sh

IMAGE_NUMBER=$1
if [[ IMAGE_NUMBER == "" ]]; then
	 IMAGE_NUMBER=local-test
fi

docker build -t tskinn/sqsd:${IMAGE_NUMBER} .

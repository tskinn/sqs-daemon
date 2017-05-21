#!/bin/bash

IMAGE_NUMBER=$1
if [[ IMAGE_NUMBER == "" ]]; then
	 IMAGE_NUMBER=local-test
fi

docker push tskinn/sqsd:${IMAGE_NUMBER}

#!/bin/bash

echo "build admission-webhook-example bin"
# dep ensure -v
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o admission-webhook-example

echo "build admission-webhook-example image"
: ${DOCKER_USER:? required}
docker build --no-cache -t ${DOCKER_USER}/admission-webhook-example:v1 .
# docker push ${DOCKER_USER}/admission-webhook-example:v1

rm -rf admission-webhook-example
echo "build done"

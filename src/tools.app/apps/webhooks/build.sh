#!/bin/bash
set -u

DOCKER_USER="zhengjin"

function build_bin() {
  echo "build admission-webhook-example bin"
  #dep ensure -v
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o admission-webhook-example
}

function build_image() {
  bin_file="admission-webhook-example"
  if [[ ! -f ${bin_file} ]]; then
    echo "go bin file ${bin_file} not exist."
    exit -1
  fi
  
  echo "build admission-webhook-example image"
  : ${DOCKER_USER:? required}
  docker build --no-cache -t ${DOCKER_USER}/admission-webhook-example:v1 .
  #docker push ${DOCKER_USER}/admission-webhook-example:v1

  #rm -rf ${bin_file} 
}

if [[ $1 == "bin" ]]; then
  build_bin
fi

if [[ $1 = "image" ]]; then
  build_image
fi

echo "build done"

#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

# build {path of folder} = build a docker image
# -t or --tag = Sets a name and optionally a tag in the "name:tag" formant
IMAGE=jacky-node-app:1.0
CONTAINER=jacky-node-app-1-0

docker stop $CONTAINER
docker container rm $CONTAINER

docker image rm $IMAGE
docker build --no-cache -t $IMAGE .
docker images


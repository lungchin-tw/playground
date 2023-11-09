#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

CONTAINER=jacky-node-app-1-0


docker run --name $CONTAINER -d -p 8000:80 jacky-node-app:1.0
docker logs $CONTAINER
docker ps



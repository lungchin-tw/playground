#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

CONTAINER=allinone-latest


docker run --name $CONTAINER -d -p 8080:8080 allinone:latest
docker logs $CONTAINER
docker ps



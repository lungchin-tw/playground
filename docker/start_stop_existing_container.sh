#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

docker ps
docker start c6013bfff3de

docker ps
sleep 3

docker stop c6013bfff3de
docker ps



#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

NETWORK=jacky-network

docker network rm $NETWORK
docker network create $NETWORK
docker network connect --ip 172.17.0.2 $NETWORK postgres-gamedb
docker network connect --ip 172.17.0.3 $NETWORK postgres-foreign
docker network connect $NETWORK ubuntu-psql
docker network inspect $NETWORK > network.log

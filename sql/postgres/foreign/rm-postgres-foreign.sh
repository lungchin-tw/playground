#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

DB=foreign

# --name = Assign a name to the container
# -p or --publish = Publish a container's port to the host
# -p {HOST_PORT}:{CONTAINER_PORT}
docker stop postgres-$DB
docker network disconnect bridge postgres-$DB
docker container rm postgres-$DB
docker ps -a




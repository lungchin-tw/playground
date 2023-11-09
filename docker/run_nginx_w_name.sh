#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

# --name = Assign a name to the container
# -p or --publish = Publish a container's port to the host
# -p {HOST_PORT}:{CONTAINER_PORT}
docker run --name jacky-nginx-demo -d -p 9000:80 nginx:1.25.2
docker logs jacky-nginx-demo
docker ps
docker inspect jacky-nginx-demo



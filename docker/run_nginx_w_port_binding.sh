#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

# -p or --publish = Publish a container's port to the host
# -p {HOST_PORT}:{CONTAINER_PORT}
docker run -d -p 8080:80 nginx:1.25.2




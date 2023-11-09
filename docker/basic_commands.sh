#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

# List all Docker images
docker images

# List running containers
docker ps

# -a or --all = List all containers
docker ps -a

# Login to a container
# docker exec -it 041668a2a002 bash
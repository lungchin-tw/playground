#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

# Pull Nginx image with specific version
docker pull nginx:1.25.2

# Pull Nginx image without specific version, then it will pull the latest image
docker pull nginx


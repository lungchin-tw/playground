#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x


# If you want to run the image in the background
# add the -d option
# docker run -d nginx:1.25.2

# Run a new Nginx image with specific version but without specifying its name
docker run nginx:1.25.2



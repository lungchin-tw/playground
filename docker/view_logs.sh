#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

# view a container's logs
docker logs 3b4d63efe31798df34f433805802d59616e00454f7512c591e7992f639f8f556



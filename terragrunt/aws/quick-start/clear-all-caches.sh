#!/bin/bash

set +x

echo '$(dirname $0):' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '$(pwd):' $(pwd)

echo 'whoami' : $(whoami)

echo '[pushd $(dirname $0)]:'
pushd $(dirname $0)
echo '[pwd]:' $(pwd)

find . -type d -name ".terragrunt-cache" -prune -exec rm -rf {} \;

echo '[popd]:'
popd
echo '[pwd]:' $(pwd)

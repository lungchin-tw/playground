#! /usr/bin/bash

echo '$(dirname $0):' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '$(pwd):' $(pwd)
pushd $(dirname $0)

set -x

openssl verify -CAfile ca-cert.pem server-cert.pem
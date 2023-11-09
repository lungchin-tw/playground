#! /usr/bin/bash

echo '$(dirname $0):' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '$(pwd):' $(pwd)
pushd $(dirname $0)

set -x

# -noout = Don't output the original encoded value, we only want to see the readable text format
openssl x509 -in ca-cert.pem -noout -text
openssl x509 -in server-cert.pem -noout -text
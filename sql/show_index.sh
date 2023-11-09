#! /usr/bin/bash

# set -x

echo '[dirname $0]:' $(dirname $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

mysql -u jacky -p playground -e  \
'SHOW INDEX FROM user_details;'
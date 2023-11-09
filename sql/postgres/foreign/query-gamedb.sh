#! /usr/bin/bash

export 

USER=jacky
PASSWORD=123456
DB=gamedb
HOST_PORT=15432


echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

psql --version

psql postgresql://$USER:$PASSWORD@localhost:$HOST_PORT/$DB <<EOF
\o query.sql
\dt
SELECT * FROM user_data;
SELECT * FROM class;
EOF
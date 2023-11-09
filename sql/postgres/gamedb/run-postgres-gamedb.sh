#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

USER=jacky
PASSWORD=123456
DB=gamedb
HOST_PORT=15432 

# --name = Assign a name to the container
# -p or --publish = Publish a container's port to the host
# -p {HOST_PORT}:{CONTAINER_PORT}
docker run --name postgres-$DB \
-e POSTGRES_PASSWORD=$PASSWORD \
-e POSTGRES_USER=$USER \
-e POSTGRES_DB=$DB \
-p $HOST_PORT:5432 \
-d postgres-demo:latest

docker logs postgres-$DB
docker ps

sleep 3

docker network connect --ip 172.17.0.2 bridge postgres-$DB 


psql postgresql://$USER:$PASSWORD@localhost:$HOST_PORT/$DB -f user.sql -f class.sql

psql postgresql://$USER:$PASSWORD@localhost:$HOST_PORT/$DB <<EOF
\o run.log
CREATE USER fdwuser WITH PASSWORD 'secret';
GRANT USAGE ON SCHEMA PUBLIC TO fdwuser;
GRANT SELECT ON class TO fdwuser;
GRANT SELECT ON user_data TO fdwuser;
\du+
SELECT * FROM pg_catalog.pg_user;
EOF




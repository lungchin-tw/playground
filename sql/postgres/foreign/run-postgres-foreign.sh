#! /usr/bin/bash

echo '[dirname $0]:' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

set -x

USER=jacky
PASSWORD=123456
DB=foreign
HOST_PORT=25432

# --name = Assign a name to the container
# -p or --publish = Publish a container's port to the host
# -p {HOST_PORT}:{CONTAINER_PORT}
docker run --name postgres-$DB \
-e POSTGRES_PASSWORD=$PASSWORD \
-e POSTGRES_USER=$USER \
-e POSTGRES_DB=$DB \
-p $HOST_PORT:5432 \
-d postgres-demo:latest

docker network connect --ip 172.17.0.3 bridge postgres-$DB
docker logs postgres-$DB
docker ps

sleep 3

psql --version

psql postgresql://$USER:$PASSWORD@localhost:$HOST_PORT/$DB <<EOF
\o run.log
SHOW hba_file;
SHOW data_directory;
CREATE EXTENSION IF NOT EXISTS postgres_fdw;
SELECT * FROM pg_extension;
CREATE SERVER fdw_gamedb FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host '172.17.0.2', port '5432', dbname 'gamedb');
\des
SELECT * FROM pg_foreign_server;
CREATE USER MAPPING FOR jacky SERVER fdw_gamedb OPTIONS (user 'fdwuser', password 'secret');
SELECT * FROM pg_user_mapping;
SELECT * FROM pg_user_mappings;
GRANT USAGE ON FOREIGN SERVER fdw_gamedb TO jacky;
IMPORT FOREIGN SCHEMA public FROM SERVER fdw_gamedb INTO public;
EOF





#! /usr/bin/bash

# set -x

echo '[dirname $0]:' $(dirname $0)
echo '[pwd]:' $(pwd)
pushd $(dirname $0)

mysql -u jacky -p playground -e  \
'
EXPLAIN SELECT * FROM user_details
WHERE user_id = (SELECT min(user_id) FROM user_details);
'
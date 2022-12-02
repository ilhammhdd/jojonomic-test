#!/bin/sh

mkdir /home/ilhammhdd/jojonomic-test/pgdata
chmod 777 /home/ilhammhdd/jojonomic-test/pgdata/

docker network create jojonomic-test-net
docker volume create --driver local --opt device=/home/ilhammhdd/jojonomic-test/pgdata --opt type=none --opt o=bind jojonomic-test-pgdata-vol
docker volume create --driver local --opt device=/mnt/c/Users/milha/Home/jojonomic-test/misc/postgres_migrations --opt type=none --opt o=bind jojonomic-test-pgmigrations-vol
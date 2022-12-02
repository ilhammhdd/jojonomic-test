#!/bin/sh

docker network rm jojonomic-test-net
docker volume rm jojonomic-test-pgdata-vol
docker volume rm jojonomic-test-pgmigrations-vol
docker network prune
docker volume prune

sudo rm -rf ~/jojonomic-test/pgdata
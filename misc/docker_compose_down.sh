#!/bin/sh

docker compose -f kafka_zookeeper.yaml -f kafdrop.yaml down -v
docker network rm jojonomic-test-net
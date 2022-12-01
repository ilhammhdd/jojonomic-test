#!/bin/sh

docker network create jojonomic-test-net
docker compose -f kafka_zookeeper.yaml -f kafdrop.yaml up --detach
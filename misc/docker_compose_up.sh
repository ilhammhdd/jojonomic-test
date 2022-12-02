#!/bin/sh

docker compose -f kafka_zookeeper.yaml -f kafdrop.yaml -f postgres.yaml up --detach
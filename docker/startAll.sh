#!/bin/bash

components=("kafka" "etcd" "redis" "grafana" "m3db")

for comp in ${components[@]}; do
    docker-compose -f $comp/docker-compose.yml down
    docker-compose -f $comp/docker-compose.yml up -d
done
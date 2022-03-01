#!/bin/bash

 docker network create \
    --driver=bridge \
    --subnet=172.18.0.0/23 \
    sr_cluster_network

# components=("kafka" "etcd" "redis" "grafana" "m3db" "flink" "druid")
#components=("kafka" "etcd" "redis" "m3db" "grafana")
components=("etcd" "m3db" "grafana")

for comp in ${components[@]}; do
    docker-compose -f $comp/docker-compose.yml down
done

root_path=/data
#data_folders=("etcd1" "etcd2" "etcd3" "zk1" "zk2" "zk3" "kfk1" "kfk2" "kfk3" "m3db1" "m3db2" "m3db3" "grafana")
data_folders=("etcd1" "etcd2" "etcd3" "m3db1" "m3db2" "m3db3" "m3c1")

for dir in ${data_folders[@]}; do
    wd=$root_path/$dir

    if [[ -d $wd ]]; then
        echo "Removing " $wd
        sudo rm -rf $wd
    fi
    
    sudo mkdir -p $wd
    sudo chown nobody:nogroup $wd
    sudo chmod 777 $wd -R
done


for comp in ${components[@]}; do
    docker-compose -f $comp/docker-compose.yml build
    docker-compose -f $comp/docker-compose.yml up -d
done


echo "sleeping for creating topic 5s"
sleep 5
bash kafka/createTopic.sh

echo "m3db detup namespace"
bash m3db/m3dbTest.sh
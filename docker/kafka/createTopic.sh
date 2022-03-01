#!/bin/bash
replicationFactor=3 #no. of brokers
partition=5
topics=("websvc" "ts")


partition_bigdata=5
topic_bigdata=("tsbig")

function createTopics() {
    for topic in ${topics[@]};
    do
    echo "Creating Topic : " $topic

    docker exec -t kafka1 /bin/sh -c \
    "bin/kafka-topics.sh --create --zookeeper 172.18.0.51:2181,172.18.0.52:2181,172.18.0.53:2181 \
    --replication-factor $replicationFactor  --partitions $partition \
    --topic $topic"

    echo "Topic Created : " $topic " - Partition Count : " $partition
    done
}

function createBigTopics() {
    for topic in ${topic_bigdata[@]};
    do
    echo "Creating Topic : " $topic

    docker exec -t kafka1 /bin/sh -c \
    "bin/kafka-topics.sh --create --zookeeper 172.18.0.51:2181,172.18.0.52:2181,172.18.0.53:2181 \
    --replication-factor $replicationFactor  --partitions $partition_bigdata \
    --topic $topic"

    echo "Topic Created : " $topic " - Partition Count : " $partition_bigdata
    done
}

createTopics
createBigTopics


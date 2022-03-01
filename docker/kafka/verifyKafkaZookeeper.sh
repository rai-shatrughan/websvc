#!/bin/bash

docker exec kafka1 bash -c \
"bin/kafka-topics.sh --describe --topic ts --zookeeper 172.18.0.51:2181,172.18.0.52:2181,172.18.0.53:2181"
	
for i in 172.18.0.51 172.18.0.52 172.18.0.53; do
   docker exec zookeeper1 bash -c "echo stat | nc $i 2181 | grep Mode"
done

#docker-compose logs | grep started
#docker-compose logs | grep controller 
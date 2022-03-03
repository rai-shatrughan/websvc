# websvc
This is Microservice setup.

It has a Gateway(based on lura and a custom from scratch gateway), Pick any of two.
> - go run gateway/gateway.go &
> - go run gateway/lugat/main.go &

It has three light weight microservices
> - auth-server - Implementation (In progress)
> - device-server - API to manage sensors
> - timeseries-server - API to push and retrieve sensors data

How to start:
> - Install docker and docker-compose if not installed
> - create a "/data" directory and give following permission 
> `this directory is used for data persistence from docker containers`
> - "sudo chown nobody:nogroup /data/"
> - run following command "cd docker; bash buildNStartAll.sh" 
`Note - dropping into docker directory is important. docker-compose will look for a .env file which is present into docker dir.`
> - once script finishes you will see few docker conatiners like kafka,zookeeper,etcd up n running.
> - drop into project root directory
> - run "bash startWebsvc.sh"
> - services are available at 8080 port
> - swagger specs are available at *swaggers* dir

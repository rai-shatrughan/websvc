#!/bin/bash

trap "exit" INT TERM ERR
trap "kill 0" EXIT

#webservices
go run auth/cmd/auth-server/main.go --host 0.0.0.0 --port 9001 --scheme http & 
go run device/cmd/device-server/main.go --host 0.0.0.0 --port 8001 --scheme http & 
go run timeseries/cmd/timeseries-server/main.go --host 0.0.0.0 --port 8002 --scheme http &

#gateway/lugat - use any one
#gateway is high performant than lugat as of Feb20, 2022
#go run gateway/gateway.go &
go run gateway/lugat/main.go &

go run middleware/writer/ts_writer.go &

wait
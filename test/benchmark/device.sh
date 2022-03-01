#!/bin/bash

go-wrk -c 100 -d 10 -T 50000 -M GET \
-H "X-API-Key: srkey12345" \
http://localhost:8080/api/v1/device/version

go-wrk -c 100 -d 10 -T 50000 -M GET \
-H "X-API-Key: srkey12345" \
http://localhost:8080/api/v1/device/6fdae6af-226d-48bd-8b61-699758137eb3

#!/bin/bash
go-wrk -c 10 -d 10 -T 50000 -M GET \
-H "X-API-Key: srkey12345" \
http://localhost:8080/api/v1/timeseries/6fdae6af-226d-48bd-8b61-699758137eb3?duration=1m
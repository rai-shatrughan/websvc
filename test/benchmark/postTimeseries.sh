#!/bin/bash
start=`date`
echo "Start Time - " $start

go-wrk -c 8000 -d 10 -T 30000 -M PUT \
-H "X-API-Key: srkey12345" \
-H "Content-Type: application/json" \
-body @json/timeseries.json \
http://localhost:8002/api/v1/timeseries/6fdae6af-226d-48bd-8b61-699758137eb3

end=`date`
echo "End Time - " $end

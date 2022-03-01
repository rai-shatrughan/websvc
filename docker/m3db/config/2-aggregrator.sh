curl -vvvsSf -H "Cluster-Environment-Name: namespace/m3db-cluster-name" \
-X POST http://localhost:7201/api/v1/services/m3aggregator/placement/init -d '{
    "num_shards": 64,
    "replication_factor": 2,
    "instances": [
        {
            "id": "m3aggregator01:6000",
            "isolation_group": "availability-zone-a",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "m3aggregator01:6000",
            "hostname": "m3aggregator01",
            "port": 6000
        },
        {
            "id": "m3aggregator02:6000",
            "isolation_group": "availability-zone-b",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "m3aggregator02:6000",
            "hostname": "m3aggregator02",
            "port": 6000
        }
    ]
}'



curl -vvvsSf -H "Cluster-Environment-Name: namespace/m3db-cluster-name" \
-H "Topic-Name: aggregator_ingest" -X POST http://localhost:7201/api/v1/topic/init -d '{
    "numberOfShards": 64
}'

curl -vvvsSf -H "Cluster-Environment-Name: namespace/m3db-cluster-name" \
-H "Topic-Name: aggregator_ingest" -X POST http://localhost:7201/api/v1/topic -d '{
  "consumerService": {
    "serviceId": {
      "name": "m3aggregator",
      "environment": "namespace/m3db-cluster-name",
      "zone": "embedded"
    },
    "consumptionType": "REPLICATED",
    "messageTtlNanos": "300000000000"
  }
}'

curl -vvvsSf -H "Cluster-Environment-Name: namespace/m3db-cluster-name" \
-H "Topic-Name: aggregated_metrics" -X POST http://localhost:7201/api/v1/topic/init -d '{
    "numberOfShards": 64
}'
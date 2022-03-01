curl -X POST http://localhost:7201/api/v1/database/create -d '{
  "namespaceName": "ts",
  "retentionTime": "8760h"
}' | jq .

curl -X POST http://localhost:7201/api/v1/database/create -d '{
  "namespaceName": "unaggregated",
  "retentionTime": "8760h"
}' | jq .

curl -X POST http://localhost:7201/api/v1/database/create -d '{
  "namespaceName": "default",
  "retentionTime": "8760h"
}' | jq .

curl http://localhost:7201/api/v1/services/m3db/namespace | jq .

echo "******************************"

curl -X POST http://localhost:7201/api/v1/services/m3db/namespace/ready -d '{
  "name": "default"
}' | jq .

echo "******************************"

curl -X POST http://localhost:7201/api/v1/json/write -d '{
  "tags": 
    {
      "__name__": "third_avenue",
      "city": "new_york",
      "checkout": "1"
    },
    "timestamp": '\"$(date "+%s")\"',
    "value": 3347.26
}'

curl -X POST http://localhost:7201/api/v1/json/write -d '{
  "tags": 
    {
      "__name__": "third_avenue",
      "city": "new_york",
      "checkout": "1"
    },
    "timestamp": '\"$(date "+%s")\"',
    "value": 5347.26
}'

curl -X POST http://localhost:7201/api/v1/json/write -d '{
  "tags": 
    {
      "__name__": "third_avenue",
      "city": "new_york",
      "checkout": "1"
    },
    "timestamp": '\"$(date "+%s")\"',
    "value": 9347.26
}'
  
curl -X "POST" -G "http://localhost:7201/api/v1/query_range" \
  -d "query=third_avenue" \
  -d "start=$(date "+%s" -d "45 seconds ago")" \
  -d "end=$( date +%s )" \
  -d "step=5s" | jq .  



# curl -X POST http://localhost:7201/api/v1/database/create -d '{
#   "type": "local",
#   "namespaceName": "ts",
#   "retentionTime": "8760h"
# }' | jq .

# curl -X POST http://localhost:7201/api/v1/services/m3db/namespace/ready -d '{
#   "name": "ts"
# }' | jq .
 
#   ====================WORKS WITH NON_CLUSTER SETUP ONLY============
#   curl -X POST http://localhost:7201/api/v1/database/create -d '{
#   "type": "local",
#   "namespaceName": "default",
#   "retentionTime": "8760h"
# }'
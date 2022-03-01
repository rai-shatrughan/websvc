kubectl port-forward --address 0.0.0.0 service/grafana 3100:3000 &
kubectl port-forward --address 0.0.0.0 service/m3coordinator-simple-cluster 7201:7201 &

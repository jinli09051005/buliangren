#!/bin/bash


kubectl delete -f config/
#hack/update-codegen.sh 
apiserver-boot build container --image demo.harbor.com/aggregator/demo-apiserver:latest
docker push demo.harbor.com/aggregator/demo-apiserver:latest
kubectl apply -f config/
kubectl -n demo-apiserver logs -f --tail=100  $(kubectl -n demo-apiserver get pods | grep apiserver-apiserver | awk '{print $1}')

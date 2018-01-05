#!/bin/bash -e

endpoint=$(kubectl -n example-system get ep -l app=redis-operator --no-headers | awk '{print $2}')

case $1 in
    create-crd)
	    curl -X POST http://${endpoint}/redisopapi/v1/crd -d \
'{
  "recipe":
    {
      "name": "redises",
      "group": "example.com",
      "version": "v1alpha1",
      "scope": 0,
      "plural": "redises",
      "singular": "redis",
      "kind": "Redis"
    }
}'
        ;;

    *)
        echo "Unknown test"
        ;;
	
esac

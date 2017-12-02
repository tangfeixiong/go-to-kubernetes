#!/bin/bash -e

case $1 in
    vnc)
	    curl -X POST http://172.17.4.50:10072/v1/vnc -d \
'{
  "vnc_addr": "172.17.0.4:5900",
  "vnc_secret": "not used",
  "auth_token": "not_used"
}'
        ;;



    *)
        echo "Valid test: vnc"
        ;;
	
esac

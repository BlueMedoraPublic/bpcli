#!/bin/bash

UNIX_TIME=$(date +%s)
export COLLECTOR_NAME="intigration-test-${UNIX_TIME}"
export COLLECTOR_UUID=$(uuidgen)
export COLLECTOR_SECRET_KEY="e223ccf9-3e93-4f62-ba6e-2a9803b41bc2"
export API_ADDRESS="https://production.api.bindplane.bluemedora.com"
export BINDPLANE_HOME="/opt/bluemedora/bindplane-collector"

clean_collector() {
    ./bpcli collector list | grep $COLLECTOR_NAME | awk '{print $2}' | xargs -I{} ./bpcli collector delete --id {}
    docker ps | grep $COLLECTOR_NAME | awk '{print $1}' | xargs -I{} docker rm -f {}
}

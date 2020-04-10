#!/bin/bash

export UNIX_TIME=$(date +%s)
export COLLECTOR_NAME="intigration-test-${UNIX_TIME}"
export COLLECTOR_UUID=$(uuidgen)
export COLLECTOR_SECRET_KEY="e223ccf9-3e93-4f62-ba6e-2a9803b41bc2"
export API_ADDRESS="https://production.api.bindplane.bluemedora.com"
export BINDPLANE_HOME="/opt/bluemedora/bindplane-collector"

export MYSQL_ROOT_PASSWORD="password"

clean() {
    echo "cleaning up deployment. . ."
    ./bpcli collector list | grep $UNIX_TIME | awk '{print $2}' | xargs --no-run-if-empty -n1 ./bpcli collector delete --id

    ./bpcli credential list | grep $UNIX_TIME | awk '{print $2}' | xargs --no-run-if-empty -n1 ./bpcli credential delete --id
    rm -f *.cred

    ./bpcli source list | grep $UNIX_TIME | awk '{print $2}' | xargs --no-run-if-empty -n1 ./bpcli source delete --id
    rm -f *.source

    docker ps | grep $UNIX_TIME | awk '{print $1}' | xargs --no-run-if-empty -I{} docker rm -f {} >> /dev/null
}

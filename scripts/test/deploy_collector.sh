#!/bin/bash

set -e

ENTRYPOINT="/opt/bluemedora/bindplane-collector/scripts/run_collector_in_docker.sh"

docker run \
    -d \
    --name=$COLLECTOR_NAME \
    --entrypoint=$ENTRYPOINT \
    -e "COLLECTOR_NAME=${COLLECTOR_NAME}" \
    -e "COLLECTOR_UUID=${COLLECTOR_UUID}" \
    -e "COLLECTOR_SECRET_KEY=${COLLECTOR_SECRET_KEY}" \
    -e "API_ADDRESS=${API_ADDRESS}" \
    -e "BINDPLANE_HOME=${BINDPLANE_HOME}" \
    docker.io/bluemedora/bindplane-metrics-collector:latest

t=0
max=12
while :
do
    echo "looking for collector with uuid ${COLLECTOR_UUID}"
    ./bpcli collector list | grep ${COLLECTOR_UUID} && break

    echo "waiting for collector. . ."
    if [ "$t" != "$max" ]
    then
        t=$((t+1))
        echo "waiting for api to return the collector: ${COLLECTOR_UUID}"
        sleep 5
    else
        echo "could not find collector ${COLLECTOR_NAME} with UUID ${COLLECTOR_UUID} after 60 seconds"
        exit 1
    fi
done

#!/bin/bash

set -eE  # same as: `set -o errexit -o errtrace`

cd "$(dirname "$0")"

# globals
UNIX_TIME=$(date +%s)

# collector
COLLECTOR_NAME="intigration-test-${UNIX_TIME}"
COLLECTOR_UUID=$(uuidgen)
COLLECTOR_SECRET_KEY="e223ccf9-3e93-4f62-ba6e-2a9803b41bc2"
API_ADDRESS="https://production.api.bindplane.bluemedora.com"
BINDPLANE_HOME="/opt/bluemedora/bindplane-collector"

# postgres cred / source
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="password"
SOURCE_NAME="${UNIX_TIME}-psql-integration"
SOURCE_FILE="${SOURCE_NAME}.source"
CRED_NAME="${UNIX_TIME}-psql-integration"
CRED_FILE="${CRED_NAME}.cred"

# cleanup when done
clean () {
    echo "cleaning up deployment. . ."
    
    rm -f *.source
    rm -f *.cred

    ./bpcli source list --json | jq ".[] | select(.name | contains(\"${UNIX_TIME}\"))" | jq .id | xargs --no-run-if-empty -n1 ./bpcli source delete --id
    ./bpcli credential list | grep $UNIX_TIME | awk '{print $2}' | xargs --no-run-if-empty -n1 ./bpcli credential delete --id
    sleep 10
    ./bpcli collector delete --id $COLLECTOR_UUID || sleep 30 && ./bpcli collector delete --id $COLLECTOR_UUID
    docker ps | grep $UNIX_TIME | awk '{print $1}' | xargs --no-run-if-empty -I{} docker rm -f {} >> /dev/null

    exit
}
trap clean ERR

docker_bindplane_collector () {
    docker run -d --name=$COLLECTOR_NAME \
        --entrypoint="/opt/bluemedora/bindplane-collector/scripts/run_collector_in_docker.sh" \
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
            clean
        fi
    done

}

docker_psql () {
    docker run \
        -d \
        -p 5432:5432 \
        --name=$SOURCE_NAME \
        -e "POSTGRES_PASSWORD=${POSTGRES_PASSWORD}" \
        postgres:9.6
    sleep 15
}


pgsql_credential () {
    CRED_TYPE_ID=$(./bpcli source type get --id postgresql --json | jq '.credential_types[0].id' | tr -d '"')
    echo "detected postgresql credential type ${CRED_TYPE_ID}"

    # test this api call, dont actuall use it
    TEMPLATE=$(./bpcli credential type template --id $CRED_TYPE_ID)
    echo "detected postgresql credential type template $TEMPLATE"

    cat <<- EOF > $CRED_FILE
{
  "name": "${CRED_NAME}",
  "credential_type_id": "${CRED_TYPE_ID}",
  "parameters": {
    "password": "${POSTGRES_PASSWORD}",
    "username": "${POSTGRES_USER}"
  }
}
EOF

    CRED_ID=$(./bpcli credential create --file $CRED_FILE --json | jq .id | tr -d '"')
    echo "created postgresql credential, ID: ${CRED_ID}"

    t=0
    max=12
    while :
    do
        ./bpcli credential list | grep ${CRED_ID} && break
        if [ "$t" != "$max" ]
        then
            t=$((t+1))
            echo "waiting for api to return postgresql credential: ${CRED_ID}"
            sleep 5
        else
            echo "could not find postgresql credential ${CRED_ID} with UUID ${CRED_ID} after 60 seconds"
            clean
        fi
    done
}

pgsql_source () {
    POSTGRES_ADDR=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' $SOURCE_NAME)

    # test this api call, dont actuall use it
    TEMPLATE=$(./bpcli source type template --id postgresql)
    echo "detected postgresql source type template $TEMPLATE"

    cat <<- EOF > $SOURCE_FILE
{
  "name": "${SOURCE_NAME}",
  "source_type": "postgresql",
  "collector_id": "${COLLECTOR_UUID}",
  "collection_interval": 60,
  "credentials": {
    "credentials": "${CRED_ID}"
  },
  "configuration": {
      "collection_mode": "normal",
      "function_count": 20,
      "host": "${POSTGRES_ADDR}",
      "monitor_indexes": true,
      "monitor_sequences": false,
      "monitor_sessions": false,
      "monitor_tables": false,
      "monitor_triggers": false,
      "order_functions_by": "calls",
      "order_queries_by": "calls",
      "port": 5432,
      "query_count": 10,
      "show_query_text": true,
      "ssl_config": "No SSL"
  }
}
EOF

    JOB_ID=$(./bpcli source create --file $SOURCE_FILE --json | jq .job_id | tr -d '"')
    echo "created postgresql source, job_id: ${JOB_ID}"

    t=0
    max=6
    while :
    do
        ./bpcli job get --id ${JOB_ID} --json | jq .status | grep Complete && echo "postgres source create" && break
        if [ "$t" != "$max" ]
        then
            t=$((t+1))
            echo "waiting for api to return source: ${SOURCE_NAME}"
            sleep 5
        else
            echo "could not find source: ${SOURCE_NAME}"
            echo "job output: $(./bpcli job get --id ${JOB_ID} --json)"
            clean
        fi
    done

}

docker_bindplane_collector
docker_psql
pgsql_credential
pgsql_source
clean

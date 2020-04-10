#!/bin/bash

set -e

SOURCE_NAME="${UNIX_TIME}-mysql-integration"
SOURCE_FILE="${SOURCE_NAME}.source"

docker run \
    -d \
    --name=$SOURCE_NAME \
    -e "MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}" \
    -e "MYSQL_DATABASE=test" \
    mysql/mysql-server:8.0

sleep 15

# test this api call, dont actuall use it
TEMPLATE=$(./bpcli source type template --id mysql)
echo "detected mysql source type template $TEMPLATE"

cat <<- EOF > $SOURCE_FILE
{
  "name": "${SOURCE_NAME}",
  "source_type": "mysql",
  "collector_id": "${COLLECTOR_UUID}",
  "collection_interval": 60,
  "credentials": {
    "credentials": "${CRED_ID}"
  },
  "configuration": {
    "collection_mode": "normal",
    "connection_timeout": 15,
    "host": "localhost",
    "monitor_databases": "user",
    "monitor_queries": true,
    "monitor_tables": true,
    "order_queries_by": "avg_latency",
    "order_tablespaces_by": "file_size",
    "port": 3306,
    "query_count": 10,
    "query_history_interval": 24,
    "query_timeout": 5,
    "ssl_config": "No Verify",
    "table_space_count": 200
  }
}
EOF
JOB_ID=$(./bpcli source create --file $SOURCE_FILE --json | jq .job_id | tr -d '"')
echo "created mysql source, job_id: ${JOB_ID}"

t=0
max=6
while :
do
    ./bpcli source list | grep $SOURCE_NAME && break
    if [ "$t" != "$max" ]
    then
        t=$((t+1))
        echo "waiting for api to return source: ${SOURCE_NAME}"
        sleep 5
    else
        echo "could not find source: ${SOURCE_NAME}"
        echo "job output: $(./bpcli job get --id ${JOB_ID} --json)"
        exit 1
    fi
done

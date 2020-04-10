#!/bin/bash

set -e

CRED_NAME="${UNIX_TIME}-mysql-integration"
CRED_FILE="${CRED_NAME}.cred"

CRED_TYPE_ID=$(./bpcli source type get --id mysql --json | jq '.credential_types[0].id' | tr -d '"')
echo "detected mysql credential type ${CRED_TYPE_ID}"

# test this api call, dont actuall use it
TEMPLATE=$(./bpcli credential type template --id $CRED_TYPE_ID)
echo "detected mysql credential type template $TEMPLATE"

cat <<- EOF > $CRED_FILE
{
  "name": "${CRED_NAME}",
  "credential_type_id": "${CRED_TYPE_ID}",
  "parameters": {
    "password": "${MYSQL_ROOT_PASSWORD}",
    "username": "root"
  }
}
EOF
CRED_ID=$(./bpcli credential create --file $CRED_FILE --json | jq .id | tr -d '"')
echo "created mysql credential, ID: ${CRED_ID}"
export CRED_ID

t=0
max=12
while :
do
    ./bpcli credential list | grep ${CRED_ID} && break
    if [ "$t" != "$max" ]
    then
        t=$((t+1))
        echo "waiting for api to return mysql credential: ${CRED_ID}"
        sleep 5
    else
        echo "could not find mysql credential ${CRED_ID} with UUID ${CRED_ID} after 60 seconds"
        exit 1
    fi
done

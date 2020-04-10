#!/bin/bash

set -eE  # same as: `set -o errexit -o errtrace`
cd "$(dirname "$0")"

. ./env.sh

if [ -z ${COLLECTOR_NAME+x} ]; then echo "COLLECTOR_NAME is unset"; exit 1; fi
if [ -z ${COLLECTOR_UUID+x} ]; then echo "COLLECTOR_UUID is unset"; exit 1; fi
if [ -z ${COLLECTOR_SECRET_KEY+x} ]; then echo "COLLECTOR_SECRET_KEY is unset"; exit 1; fi
if [ -z ${API_ADDRESS+x} ]; then echo "API_ADDRESS is unset"; exit 1; fi
if [ -z ${BINDPLANE_HOME+x} ]; then echo "BINDPLANE_HOME is unset"; exit 1; fi

trap clean ERR

./deploy_collector.sh
./deploy_credential.sh
./deploy_source.sh

clean

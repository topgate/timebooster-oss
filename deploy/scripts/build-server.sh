#! /bin/bash -eu

export TIMEBOOSTER_ARTIFACTS=`pwd`/build
export TIMEBOOSTER_SERVICE_ACCOUNT=`cat "server/gae/assets/firebase-admin.json" | jq -r '.["client_email"]'`

./server/scripts/build.sh

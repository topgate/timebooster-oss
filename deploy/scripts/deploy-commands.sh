#! /bin/bash -eu

.timebooster/scripts/gcloud-auth.sh .timebooster/service-account.json

# Deploy Commands
export GCP_PROJECT_ID=`gcloud config get-value project | tr -d "\n"`
export TIMEBOOSTER_ARTIFACTS=`pwd`/build

gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/runner" "gs://${GCP_PROJECT_ID}.appspot.com/tools/stable/"
gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/cli" "gs://${GCP_PROJECT_ID}.appspot.com/tools/stable/"

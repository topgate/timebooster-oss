#! /bin/bash -eu

.timebooster/scripts/gcloud-auth.sh .timebooster/service-account.json

export GCP_PROJECT_ID=`cat "server/gae/assets/firebase-admin.json" | jq -r '.["project_id"]'`
export TIMEBOOSTER_ARTIFACTS=`pwd`/build
export TIMEBOOSTER_SERVICE_ACCOUNT=`cat "server/gae/assets/firebase-admin.json" | jq -r '.["client_email"]'`

# Endpoint Deploy
cat deploy/openapi.yaml \
  | sed "s/GCP_PROJECT_ID/$GCP_PROJECT_ID/g" \
  > deploy/deploy.yaml
gcloud service-management deploy deploy/deploy.yaml

cd server

export WORKSPACE=`pwd`/gae
export GOPATH=`pwd`/.gopath/linux:`pwd`:${GOPATH}

# 本番Deploy
echo "Deploy stable"
appcfg.py --secure --no_cookies \
    --application=${GCP_PROJECT_ID} \
    --env_variable=GCP_PROJECT_ID:"${GCP_PROJECT_ID}" \
    --env_variable=TIMEBOOSTER_SERVICE_ACCOUNT:"${TIMEBOOSTER_SERVICE_ACCOUNT}" \
    --version="stable" \
    --oauth2_access_token=$(gcloud auth print-access-token 2> /dev/null) \
    update ./gae

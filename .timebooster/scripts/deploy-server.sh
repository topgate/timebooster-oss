#! /bin/bash -eu

.timebooster/scripts/gcloud-auth.sh .timebooster/service-account.json
cd server

export WORKSPACE=`pwd`/gae
export GOPATH=`pwd`/.gopath/linux:`pwd`:${GOPATH}
export GCP_PROJECT_ID=`cat "./gae/assets/firebase-admin.json" | jq -r '.["project_id"]'`

# バージョンごとのデプロイ
appcfg.py --secure --no_cookies \
    --application=${GCP_PROJECT_ID} \
    --env_variable=GCP_PROJECT_ID:"${GCP_PROJECT_ID}" \
    --env_variable=TIMEBOOSTER_SERVICE_ACCOUNT:"${TIMEBOOSTER_SERVICE_ACCOUNT}" \
    --version="build${CIRCLE_BUILD_NUM}" \
    --oauth2_access_token=$(gcloud auth print-access-token 2> /dev/null) \
    update ./gae

if [[ "${CIRCLE_BRANCH:-nil}" =~ develop ]]; then
    echo "Deploy beta"
    appcfg.py --secure --no_cookies \
        --application=${GCP_PROJECT_ID} \
        --env_variable=GCP_PROJECT_ID:"${GCP_PROJECT_ID}" \
        --env_variable=TIMEBOOSTER_SERVICE_ACCOUNT:"${TIMEBOOSTER_SERVICE_ACCOUNT}" \
        --version="beta" \
        --oauth2_access_token=$(gcloud auth print-access-token 2> /dev/null) \
        update ./gae
fi

if [[ "${CIRCLE_BRANCH:-nil}" =~ master ]]; then
    echo "Deploy stable"
    appcfg.py --secure --no_cookies \
        --application=${GCP_PROJECT_ID} \
        --env_variable=GCP_PROJECT_ID:"${GCP_PROJECT_ID}" \
        --env_variable=TIMEBOOSTER_SERVICE_ACCOUNT:"${TIMEBOOSTER_SERVICE_ACCOUNT}" \
        --version="stable" \
        --oauth2_access_token=$(gcloud auth print-access-token 2> /dev/null) \
        update ./gae
fi

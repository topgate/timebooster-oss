#! /bin/bash -eu

.timebooster/scripts/gcloud-auth.sh .timebooster/service-account.json
export GCP_PROJECT_ID=`gcloud config get-value project | tr -d "\n"`

gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/runner" "gs://${GCP_PROJECT_ID}.appspot.com/tools/$CIRCLE_BUILD_NUM/"
gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/cli" "gs://${GCP_PROJECT_ID}.appspot.com/tools/$CIRCLE_BUILD_NUM/"

if [[ "${CIRCLE_BRANCH:-nil}" =~ develop ]]; then
    echo "Deploy beta"
    gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/runner" "gs://${GCP_PROJECT_ID}.appspot.com/tools/beta/"
    gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/cli" "gs://${GCP_PROJECT_ID}.appspot.com/tools/beta/"
fi

if [[ "${CIRCLE_BRANCH:-nil}" =~ master ]]; then
    echo "Deploy stable"
    gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/runner" "gs://${GCP_PROJECT_ID}.appspot.com/tools/stable/"
    gsutil rsync -r -U "$TIMEBOOSTER_ARTIFACTS/cli" "gs://${GCP_PROJECT_ID}.appspot.com/tools/stable/"
fi

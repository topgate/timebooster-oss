#! /bin/bash -eu

SERVICE_ACCOUNT_JSON=$1

SERVICE_ACCOUNT_EMAIL=`cat ${SERVICE_ACCOUNT_JSON} | jq -r '.["client_email"]'`
GCP_PROJECT_ID=`cat ${SERVICE_ACCOUNT_JSON} | jq -r '.["project_id"]'`

# echo "GCP Service Account File=$SERVICE_ACCOUNT_JSON"
# echo "GCP Service Account=$SERVICE_ACCOUNT_EMAIL"
# echo "GCP Project ID=$GCP_PROJECT_ID"

# login gcloud
gcloud auth activate-service-account ${SERVICE_ACCOUNT_EMAIL} --key-file ${SERVICE_ACCOUNT_JSON} --project ${GCP_PROJECT_ID} &> /dev/null
gcloud config set project ${GCP_PROJECT_ID} &> /dev/null
gcloud config set account ${SERVICE_ACCOUNT_EMAIL} &> /dev/null

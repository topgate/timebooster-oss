#! /bin/bash -eu

cd server

export WORKSPACE=`pwd`/gae
export GOPATH=`pwd`/.gopath/linux:`pwd`:${GOPATH}
export GCP_PROJECT_ID=`cat "./gae/assets/firebase-admin.json" | jq -r '.["project_id"]'`

# 依存関係取得
prjdep restore

# config設定
go run ./scripts/config.go

# test & build check
goapp test -tags gaetest -parallel 128 ./src/...
goapp build ./src/...

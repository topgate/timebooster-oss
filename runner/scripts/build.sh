#! /bin/bash -eu

cd runner

export GOPATH=`pwd`/.gopath/linux:${GOPATH}
export GCP_PROJECT_ID=`cat "../server/gae/assets/firebase-admin.json" | jq -r '.["project_id"]'`
export TIMEBOOSTER_PROJECT_ID=$GCP_PROJECT_ID

# restore dependencies
go version
prjdep restore

# test
go test ./...

# build all artifacts
INSTALL_TARGET=$TIMEBOOSTER_ARTIFACTS/runner
mkdir $INSTALL_TARGET

export EXE_NAME="tbrun"

echo "########################"
echo "## Build for Linux"
echo "########################"
export GOOS=linux
export GOARCH=amd64

go build -o $EXE_NAME
mkdir $INSTALL_TARGET/linux
mv ./$EXE_NAME $INSTALL_TARGET/linux

echo "########################"
echo "## Build for Mac"
echo "########################"
export GOOS=darwin
export GOARCH=amd64

go build -o $EXE_NAME
mkdir $INSTALL_TARGET/mac
mv ./$EXE_NAME $INSTALL_TARGET/mac

echo "########################"
echo "## Build for Windows"
echo "########################"
export GOOS=windows
export GOARCH=amd64

go build -o $EXE_NAME.exe
mkdir $INSTALL_TARGET/windows
mv ./$EXE_NAME.exe $INSTALL_TARGET/windows

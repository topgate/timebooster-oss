#! /bin/bash -eu

cd cli

export GOPATH=`pwd`/.gopath/linux:${GOPATH}

# ビルドバージョンを更新する
sed -i "s/__BUILD_VERSION__/${CIRCLE_BUILD_NUM:-__BUILD_VERSION__}/g" "./timebooster.go"
cat "./timebooster.go" | grep "app.Version"

# restore dependencies
go version
prjdep restore

# test
go test ./...

# build all artifacts
INSTALL_TARGET=$TIMEBOOSTER_ARTIFACTS/cli
mkdir $INSTALL_TARGET

export EXE_NAME="timebooster"

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

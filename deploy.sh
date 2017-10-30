#! /bin/sh -eu

if [[ $(uname) = CYGWIN* ]]; then
    export WORK_PATH="/`cygpath -w ${PWD}`"
    export WORK_PATH=`echo $WORK_PATH | sed -e 's/:\\\\/\//g' | sed -e 's/\\\\/\//g'`
else
    export WORK_PATH=$PWD
fi

if [ ! -e ".timebooster/service-account.json" ]; then
  echo ".timebooster/service-account.json (Deploy service account) not found."
  exit 1
fi

if [ ! -e "server/gae/assets/firebase-admin.json" ]; then
  echo "server/gae/assets/firebase-admin.json (GAE/Go Service account) not found."
  exit 1
fi

chmod -R 755 .timebooster/scripts
chmod -R 755 ./cli/scripts
chmod -R 755 ./runner/scripts
chmod -R 755 ./deploy/scripts

rm -rf build/
mkdir build

docker build -t timebooster/go:latest -f .timebooster/dockerfiles/go.dockerfile .
docker build -t timebooster/appengine:latest -f .timebooster/dockerfiles/appengine.dockerfile .

# コマンドをビルド
docker run --rm \
       -v ${WORK_PATH}:/work \
       -w /work \
       timebooster/go:latest \
       ./deploy/scripts/build-commands.sh

# サーバーをビルド
docker run --rm \
      -v ${WORK_PATH}:/work \
      -w /work \
      timebooster/appengine:latest \
      ./deploy/scripts/build-server.sh

# コマンドをデプロイ
docker run --rm \
       -v ${WORK_PATH}:/work \
       -w /work \
       timebooster/appengine:latest \
       ./deploy/scripts/deploy-commands.sh

# サーバーをデプロイ
docker run --rm \
      -v ${WORK_PATH}:/work \
      -w /work \
      timebooster/appengine:latest \
      ./deploy/scripts/deploy-server.sh

#! /bin/bash -eu

if [ "`which docker`" == "" ]; then
  echo "Error [docker] not found"
  exit 1
fi

if [ "`which swagger-codegen`" == "" ]; then
  echo "Install   [swagger-codegen]"
  go get -u -f github.com/eaglesakura/swagger-codegen
else
  echo "Installed [swagger-codegen]"
fi

swagger-codegen init

if [ "`which prjdep`" == "" ]; then
  echo "Install   [prjdep]"
  go get -u -f github.com/eaglesakura/prjdep
else
  echo "Installed [prjdep]"
fi

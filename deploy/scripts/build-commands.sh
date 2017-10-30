#! /bin/bash -eu

export TIMEBOOSTER_ARTIFACTS=`pwd`/build

./cli/scripts/build.sh
./runner/scripts/build.sh

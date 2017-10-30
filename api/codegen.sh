#! /bin/bash -eu

if [[ $(uname) = CYGWIN* ]]; then
    echo "Run on Cygwin"
    export WORK_PATH="/`cygpath -w ${PWD}`"
    export WORK_PATH=`echo $WORK_PATH | sed -e 's/:\\\\/\//g' | sed -e 's/\\\\/\//g'`
else
    echo "Run on Unix"
    export WORK_PATH=$PWD
fi

# Generate API
echo "$WORK_PATH"
docker build -t timebooster/swagger:latest ./api
docker run --rm \
       -v ${WORK_PATH}:/work \
       -w /work \
       timebooster/swagger:latest \
       groovy ./api/generator.groovy

# # Copy API
rm -rf server/src/api_v1
rm -rf cli/api_v1
rm -rf runner/api_v1
cp -f api/build/swagger.json server/gae/assets/swagger.json
cp -r api/build/server server/src/api_v1
cp -r api/build/client cli/api_v1
cp -r api/build/client runner/api_v1

#! /bin/bash -eu
rm -rf server/src/api_v1
rm -rf cli/api_v1
rm -rf runner/api_v1

# execute
swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output server/src/api_v1 \
  --target go-server
swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output server/gae/assets/ \
  --target swagger-json

swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output cli/api_v1 \
  --target go-client
swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output runner/api_v1 \
  --target go-client

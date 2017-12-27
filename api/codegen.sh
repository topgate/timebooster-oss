#! /bin/bash -eu

# execute
swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output server/src/api_v1 \
  --target go-server \
  --with-clean true

swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output server/gae/assets/ \
  --target swagger-json

swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output cli/api_v1 \
  --target go-client \
  --with-clean true

swagger-codegen \
  generate --swagger api/swagger.yaml --config api/config.json \
  --output runner/api_v1 \
  --target go-client \
  --with-clean true

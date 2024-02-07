ROOT_PATH=$(cd $(dirname $0)/../ && pwd)
FRONTEND_GEN_PATH="${ROOT_PATH}/messenger-frontend"
BACKEND_GEN_PATH="${ROOT_PATH}/messenger-backend"

protoc --proto_path ${ROOT_PATH}/proto/ \
       --go_out ${BACKEND_GEN_PATH}/hello/api/gen/v1 \
       --go_opt paths=source_relative \
       --go-grpc_out ${BACKEND_GEN_PATH}/hello/api/gen/v1 \
       --go-grpc_opt paths=source_relative \
       --proto_path proto/hello \
       hello.proto

protoc --proto_path ${ROOT_PATH}/proto/ \
       --grpc-gateway_out ${BACKEND_GEN_PATH}/hello/api/gen/v1 \
       --grpc-gateway_opt logtostderr=true \
       --grpc-gateway_opt paths=source_relative \
       --proto_path proto/hello \
       hello.proto

# Path to this cli protobufjs-cli
PROTOC_GEN_JS_PATH="${FRONTEND_GEN_PATH}/node_modules/.bin/pbjs"
PROTOC_GEN_TS_PATH="${FRONTEND_GEN_PATH}/node_modules/.bin/pbts"

${PROTOC_GEN_JS_PATH} -t static-module -w es6 ${ROOT_PATH}/proto/hello/hello.proto \
       --no-create --no-encode --no-verify --no-delimited \
       -o ${FRONTEND_GEN_PATH}/api/gen/v1/hello/hello_pb.js

${PROTOC_GEN_TS_PATH} ${FRONTEND_GEN_PATH}/api/gen/v1/hello/hello_pb.js \
       -o ${FRONTEND_GEN_PATH}/api/gen/v1/hello/hello_pb.d.ts

# Path to this plugin ts-protoc-gen
# PROTOC_GEN_TS_PATH="${FRONTEND_GEN_PATH}/node_modules/.bin/protoc-gen-ts"

# protoc --proto_path ${ROOT_PATH}/proto/ \
#        --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
#        --js_out="import_style=commonjs,binary:${FRONTEND_GEN_PATH}/api/gen/v1/hello" \
#        --ts_out="${FRONTEND_GEN_PATH}/api/gen/v1/hello" \
#        --proto_path proto/hello \
#        proto/hello/hello.proto

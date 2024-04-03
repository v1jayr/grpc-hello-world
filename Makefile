
PROTOC_GEN_TS_PATH="./nodejs/node_modules/.bin/protoc-gen-ts"

NODEJS_OUT_DIR="./nodejs/lib/protogen"

GO_OUT_DIR="lib/protogen"

gen-go:
	protoc --go_out="${GO_OUT_DIR}" --go_opt=paths=source_relative --go-grpc_out="${GO_OUT_DIR}" --go-grpc_opt=paths=source_relative service.proto

gen-ts:
	protoc --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
			--js_out="import_style=commonjs,binary:${NODEJS_OUT_DIR}" \
			--ts_out="${NODEJS_OUT_DIR}" \
			service.proto
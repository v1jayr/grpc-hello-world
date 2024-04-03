gen-go:
	protoc --go_out=lib/protogen --go_opt=paths=source_relative --go-grpc_out=lib/protogen --go-grpc_opt=paths=source_relative service.proto
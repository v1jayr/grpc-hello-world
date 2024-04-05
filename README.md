# grpc-hello-world
Example grpc service with http bindings. Server and Http gateway implementation is golang and client is typescript.

## Generate stubs
`make proto-gen`

Uses
* protoc-gen-go & protoc-gen-go-grpc to generate go server and client stubs.
* protoc-gen-grpc-gateway to generate go http proxy.
* protoc-gen-grpc-gateway-ts to generate typescript client stubs.

## Build & Publish javascript
`make npm-build-patch` builds the typescript code, auto updates patch version and does npm publish to github npm registry.

## Clone and Reuse
1. Update all occurance of repository name in *package.json*
2. Update `go_package` in *service.proto*
3. Update `name` in *buf.yaml*
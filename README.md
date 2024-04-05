#grpc-hello-world

## Generate stubs
`make proto-gen`

Uses
* protoc-gen-go & protoc-gen-go-grpc to generate go server and client stubs.
* protoc-gen-grpc-gateway to generate go http proxy.
* protoc-gen-grpc-gateway-ts to generate typescript client stubs.

## Build & Publish javascript
`npm run build`
`npm publish`
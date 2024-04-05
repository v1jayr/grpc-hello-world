start-server:
	go run main.go
	go run main.go --proxy

proto-gen:
	buf generate

npm-build-patch:
	cd nodejs; npm run build; npm version patch; npm publish

npm-build-minor:
	cd nodejs; npm run build; npm version minor; npm publish
MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")
TOOLKIT_PKG = ${MODULE_NAME}/gen/cmd/toolkit

install_toolkit:
	@go install "${TOOLKIT_PKG}/..."

install_goimports:
	@go install golang.org/x/tools/cmd/goimports@latest

## TODO add source format as a githook
format: install_goimports
	go mod tidy
	goimports -w -l -local "${MODULE_NAME}" ./

generate: install_toolkit
	go generate ./...

migrate:
	go run cmd/demo/main.go migrate

run:
	go run cmd/demo/main.go


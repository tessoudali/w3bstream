MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")

install_toolkit:
	@go install github.com/iotexproject/Bumblebee/gen/cmd/...

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


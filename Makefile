
install_toolkit:
	go install github.com/iotexproject/Bumblebee/gen/cmd/...

generate: install_toolkit
	go generate ./...

migrate:
	go run cmd/demo/main.go migrate

run:
	go run cmd/demo/main.go

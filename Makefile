MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")

install_toolkit:
	@go install github.com/iotexproject/Bumblebee/gen/cmd/...

install_goimports:
	@go install golang.org/x/tools/cmd/goimports@latest

## TODO add source format as a githook
format: install_goimports
	go mod tidy
	goimports -w -l -local "${MODULE_NAME}" ./

## gen code
generate: install_toolkit
	go generate ./...

## to migrate database models, if model defines changed, make this entry
migrate:
	go run cmd/srv-applet-mgr/main.go migrate

## build srv-applet-mgr
build_server:
	@cd cmd/srv-applet-mgr && go build
	@mkdir -p build
	@mv cmd/srv-applet-mgr/srv-applet-mgr build
	@cp -r cmd/srv-applet-mgr/config build/config
	@echo 'succeed! srv-applet-mgr =>build/srv-applet-mgr*'
	@echo 'succeed! config =>build/config/'
	@echo 'modify config/local.yaml to use your server config'

## migrate first
run_server: build_server
	@cd build && ./srv-applet-mgr

## make pub_client
build_pub_client:
	@cd cmd/pub_client && go build
	@mkdir -p build
	@mv cmd/pub_client/pub_client build
	@echo 'succeed! pub_client => build/pub_client*'

clean:
	@rm -rf build/{config,pub_client,srv-applet-mgr}
	@echo 'remove build/{config,pub_client,srv-applet-mgr}'

run_depends:
	docker-compose -f testutil/docker-compose-pg.yaml up -d
	docker-compose -f testutil/docker-compose-mqtt.yaml up -d

build: build_server build_pub_client

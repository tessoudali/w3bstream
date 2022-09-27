MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")

update_go_module:
	go mod tidy

install_toolkit: update_go_module
	@go install github.com/iotexproject/Bumblebee/gen/cmd/...

install_goimports: update_go_module
	@go install golang.org/x/tools/cmd/goimports@latest

## TODO add source format as a githook
format: install_goimports
	go mod tidy
	goimports -w -l -local "${MODULE_NAME}" ./

## gen code
generate: install_toolkit
	go generate ./...

## to migrate database models, if model defines changed, make this entry
migrate: update_go_module
	go run cmd/srv-applet-mgr/main.go migrate

## build srv-applet-mgr
build_server: update_go_module
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

## create admin account
create_admin: build_server
	@cd build && ./srv-applet-mgr init_admin

## make pub_client
build_pub_client: update_go_module
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

wasm_demo: update_go_module
	@cd pkg/modules/vm/testdata && make all

build: build_server build_pub_client


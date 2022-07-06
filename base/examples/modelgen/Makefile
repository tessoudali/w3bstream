
gen_models:
	@echo "generate models"
	go generate github.com/iotexproject/w3bstream/base/examples/modelgen/pkg/models
	@echo ""

install_tools:
	@echo "install gen tools"
	cd ../../cmd && make install
	@echo ""

start_postgres:
	docker-compose -f docker-compose-pg.yaml up

init_db:
	@echo "init database"
	psql 'postgres://pguser:pgpassword@127.0.0.1:5432' -c 'create database sqlxdemo' && psql 'postgres://pguser:pgpassword@127.0.0.1:5432' -c 'create schema sqlxdemo'
	@echo ""

download:
	@echo "download go modules"
	go mod download
	@echo ""

dep:
	go get -u ./...

build: download install_tools gen_models
	@echo "build demo"
	cd cmd && go build -v && mv cmd ../sqlxdemo
	@echo ""

run: build
	./sqlxdemo

clean:
	docker-compose -f docker-compose-pg.yaml down; rm -rf data sqlxdemo

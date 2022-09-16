build_wasm:
	@cd testdata && tinygo build -o greet.wasm -scheduler=none --no-debug -target=wasi greet.go
	@echo built to testdata/greet.wasm

greet_iotex: build_wasm
	go run greet.go IoTeX

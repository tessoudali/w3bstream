build_wasm:
	@cd testdata && tinygo build -o word_count.wasm -scheduler=none --no-debug -target=wasi word_count.go
	@echo built to testdata/word_count.wasm

word_count: build_wasm
	go run word_count.go "a ab abc def def abc ab a xxxx dddd ffff"

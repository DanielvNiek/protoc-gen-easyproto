test:
	mkdir -p bin
	mkdir -p out
	mkdir -p goout
	go build -o bin/protoc-gen-easyproto .
	export PATH=$$PATH:$(PWD)/bin && \
	protoc --easyproto_out=./out --go_out=./goout ./pb/test.proto


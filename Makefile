test:
	mkdir -p bin
	go build -o bin/protoc-gen-easyproto .
	export PATH=$$PATH:$(PWD)/bin && \
	protoc --easyproto_out=./out ./pb/test.proto


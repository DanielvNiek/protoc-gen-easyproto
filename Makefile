test:
	mkdir -p bin
	mkdir -p out
	mkdir -p goout
	go build -o ~/go/bin/protoc-gen-easyproto .
	protoc --easyproto_out=./out --go_out=./goout ./pb/test.proto


build:
	go build -i -v -o go-program.exe .

lint:
	go fmt .
	go tool vet .

proto:
    cd go-rpc/proto/ && protoc --go_out=plugins=grpc:. handler.proto
    cd go-rpc/proto-fast/ && protoc --gofast_out=. handler.proto

.PHONY: proto lint build

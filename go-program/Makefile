build:
	go build -i -v -o go-program.exe .

lint:
	go fmt .
	go tool vet .

proto:
    cd go-rpc/proto/ && protoc --gofast_out=plugins=grpc:. handler.proto

.PHONY: proto lint build

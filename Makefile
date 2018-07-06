all: $(GOPATH)/bin/protoc-gen-go $(GOPATH)/src/google.golang.org/grpc service/service.pb.go grpc-server grpc-client

$(GOPATH)/bin/protoc-gen-go:
	go get -u -v github.com/golang/protobuf/protoc-gen-go

$(GOPATH)/src/google.golang.org/grpc:
	go get -u -v google.golang.org/grpc

service/service.pb.go: service.proto
	mkdir -p $(dir $@)
	protoc --go_out=plugins=grpc:$(dir $@) $<

grpc-server:
	go build -o $@ ./server

grpc-client:
	go build -o $@ ./client

clean:
	rm -rf service/service.pb.go grpc-server grpc-client

run-server:
	./grpc-server

run-client:
	./grpc-client

.PHONY: all clean run-server run-client

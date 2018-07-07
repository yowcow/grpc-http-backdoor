TARGETS := \
	$(GOPATH)/bin/dep \
	$(GOPATH)/bin/protoc-gen-go \
	$(GOPATH)/src/google.golang.org/grpc \
	service/service.pb.go \
	grpc-server \
	grpc-client

all: $(TARGETS)

$(GOPATH)/bin/dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

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

test:
	go test ./...

.PHONY: all clean run-server run-client test

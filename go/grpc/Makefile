
PRO_DIR=./proto
GO_DST_DIR=./go
CPP_DST_DIR=./cpp

# Get a brand new protoc binary, need sudo permission
protoc-install:
	wget https://github.com/google/protobuf/releases/download/v3.5.1/protobuf-cpp-3.5.1.tar.gz
	tar xzvf protobuf-cpp-3.5.1.tar.gz && rm protobuf-cpp-3.5.1.tar.gz
	make -C protobuf-3.5.1 distclean
	cd protobuf-3.5.1 && ./configure && make -j4
	make -C protobuf-3.5.1 install && ldconfig
	rm -rf protobuf-3.5.1

deps:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

gen:
	protoc -I=$(PRO_DIR) --go_out=plugins=grpc:$(GO_DST_DIR) $(PRO_DIR)/rpc.proto
	protoc -I=$(PRO_DIR) --cpp_out=$(CPP_DST_DIR) $(PRO_DIR)/rpc.proto

.PHONY: gen deps
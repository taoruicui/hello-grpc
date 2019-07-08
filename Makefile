DOCKER_RUN=docker run -w /root/protobuf -v `pwd`:/root/protobuf 951896542015.dkr.ecr.us-west-1.amazonaws.com/tools/tooling-image:protobuf

# DOCKER_RUN=docker run -w /root/protobuf -v `pwd`:/defs  namely/protoc-all

PROTO_FILES := $(shell find api -type f -iname "*.proto")

.PHONY: proto
proto: pkg/api/ api/proto
	$(DOCKER_RUN) protoc -Iapi/proto -I/root/include -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:pkg/api $(PROTO_FILES)
	$(DOCKER_RUN) protoc -Iapi/proto -I/root/include -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:pkg/api $(PROTO_FILES)
	$(DOCKER_RUN) protoc -Iapi/proto -I/root/include -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:pkg/api $(PROTO_FILES)

.PHONY: clean
clean:
	@rm -rf pkg/api/*

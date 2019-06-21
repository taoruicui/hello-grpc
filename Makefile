DOCKER_RUN=docker run -w /root/protobuf -v `pwd`:/root/protobuf proto:test

PROTO_FILES := $(shell find api -type f -iname "*.proto")

.PHONY: proto
proto: pkg/api/ api/proto
	$(DOCKER_RUN) protoc -I api/proto --go_out=plugins=grpc:pkg/api $(PROTO_FILES)

.PHONY: clean
clean:
	@rm -rf pkg/api/*

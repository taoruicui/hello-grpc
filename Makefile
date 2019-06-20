PROTO_FILES := $(shell find api -type f -iname "*.proto")

.PHONY: proto
proto: pkg/api/ api/proto
	@protoc -I api/proto --go_out=plugins=grpc:pkg/api $(PROTO_FILES)

.PHONY: clean
clean:
	@rm -rf pkg/api/*
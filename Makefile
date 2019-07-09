DOCKER_RUN = docker run -w /root/protobuf -v `pwd`:/root/protobuf 951896542015.dkr.ecr.us-west-1.amazonaws.com/tools/tooling-image:protobuf

PROTO_FILES := $(shell find api -type f -iname "*.proto")
TARGET_DIR = pkg/api
SOURCE_DIR = api/proto

all: clean directories proto

.PHONY: proto
proto: $(TARGET_DIR) $(SOURCE_DIR)
	$(DOCKER_RUN) protoc -I$(SOURCE_DIR) -I/root/include -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:$(TARGET_DIR) $(PROTO_FILES)
	$(DOCKER_RUN) protoc -I$(SOURCE_DIR) -I/root/include -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:$(TARGET_DIR) $(PROTO_FILES)
	$(DOCKER_RUN) protoc -I$(SOURCE_DIR) -I/root/include -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:$(TARGET_DIR) $(PROTO_FILES)

.PHONY: clean
clean:
	rm -rf $(TARGET_DIR)

.PHONY: directories
directories:
	mkdir -p $(TARGET_DIR)
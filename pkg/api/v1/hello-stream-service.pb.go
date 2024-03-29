// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/hello-stream-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StreamDataRequest struct {
	Api                  string      `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	StreamData           *StreamData `protobuf:"bytes,2,opt,name=streamData,proto3" json:"streamData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StreamDataRequest) Reset()         { *m = StreamDataRequest{} }
func (m *StreamDataRequest) String() string { return proto.CompactTextString(m) }
func (*StreamDataRequest) ProtoMessage()    {}
func (*StreamDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5f019a934c23b4, []int{0}
}

func (m *StreamDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamDataRequest.Unmarshal(m, b)
}
func (m *StreamDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamDataRequest.Marshal(b, m, deterministic)
}
func (m *StreamDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamDataRequest.Merge(m, src)
}
func (m *StreamDataRequest) XXX_Size() int {
	return xxx_messageInfo_StreamDataRequest.Size(m)
}
func (m *StreamDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamDataRequest proto.InternalMessageInfo

func (m *StreamDataRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *StreamDataRequest) GetStreamData() *StreamData {
	if m != nil {
		return m.StreamData
	}
	return nil
}

type StreamData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamData) Reset()         { *m = StreamData{} }
func (m *StreamData) String() string { return proto.CompactTextString(m) }
func (*StreamData) ProtoMessage()    {}
func (*StreamData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5f019a934c23b4, []int{1}
}

func (m *StreamData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamData.Unmarshal(m, b)
}
func (m *StreamData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamData.Marshal(b, m, deterministic)
}
func (m *StreamData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamData.Merge(m, src)
}
func (m *StreamData) XXX_Size() int {
	return xxx_messageInfo_StreamData.Size(m)
}
func (m *StreamData) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamData.DiscardUnknown(m)
}

var xxx_messageInfo_StreamData proto.InternalMessageInfo

func (m *StreamData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StreamData) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StreamDataResponse struct {
	Api                  string      `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	StreamData           *StreamData `protobuf:"bytes,2,opt,name=streamData,proto3" json:"streamData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StreamDataResponse) Reset()         { *m = StreamDataResponse{} }
func (m *StreamDataResponse) String() string { return proto.CompactTextString(m) }
func (*StreamDataResponse) ProtoMessage()    {}
func (*StreamDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5f019a934c23b4, []int{2}
}

func (m *StreamDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamDataResponse.Unmarshal(m, b)
}
func (m *StreamDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamDataResponse.Marshal(b, m, deterministic)
}
func (m *StreamDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamDataResponse.Merge(m, src)
}
func (m *StreamDataResponse) XXX_Size() int {
	return xxx_messageInfo_StreamDataResponse.Size(m)
}
func (m *StreamDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamDataResponse proto.InternalMessageInfo

func (m *StreamDataResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *StreamDataResponse) GetStreamData() *StreamData {
	if m != nil {
		return m.StreamData
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamDataRequest)(nil), "v1.StreamDataRequest")
	proto.RegisterType((*StreamData)(nil), "v1.StreamData")
	proto.RegisterType((*StreamDataResponse)(nil), "v1.StreamDataResponse")
}

func init() { proto.RegisterFile("v1/hello-stream-service.proto", fileDescriptor_9e5f019a934c23b4) }

var fileDescriptor_9e5f019a934c23b4 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x65, 0x57, 0x20, 0xf1, 0x40, 0x15, 0x35, 0x50, 0x55, 0x55, 0x41, 0x95, 0xa7, 0xa8,
	0x52, 0xe3, 0x26, 0x6c, 0x1d, 0xa1, 0x03, 0x73, 0x2a, 0xd8, 0x18, 0xdc, 0xe4, 0xa9, 0x58, 0x0a,
	0x71, 0x88, 0x5d, 0x1f, 0x80, 0x2b, 0x70, 0x18, 0x0e, 0xc2, 0x15, 0x38, 0x08, 0x8a, 0x2d, 0x41,
	0xd4, 0x11, 0xd8, 0x12, 0xbf, 0xff, 0x7d, 0xff, 0xef, 0x5f, 0x86, 0x4b, 0x97, 0x88, 0x27, 0x2c,
	0x4b, 0x3d, 0x37, 0xb6, 0x41, 0xf9, 0x3c, 0x37, 0xd8, 0x38, 0x95, 0x63, 0x5c, 0x37, 0xda, 0x6a,
	0x46, 0x5d, 0x32, 0x9e, 0x6c, 0xb5, 0xde, 0x96, 0x28, 0x64, 0xad, 0x84, 0xac, 0x2a, 0x6d, 0xa5,
	0x55, 0xba, 0x32, 0x41, 0xc1, 0xef, 0x61, 0xb0, 0xf6, 0x9b, 0x2b, 0x69, 0x65, 0x86, 0x2f, 0x3b,
	0x34, 0x96, 0x9d, 0x42, 0x4f, 0xd6, 0x6a, 0x44, 0xa6, 0x24, 0x3a, 0xca, 0xda, 0x4f, 0x16, 0x03,
	0x98, 0x6f, 0xd9, 0x88, 0x4e, 0x49, 0x74, 0x9c, 0xf6, 0x63, 0x97, 0xc4, 0x9d, 0xe5, 0x8e, 0x82,
	0xa7, 0x00, 0x3f, 0x13, 0xd6, 0x07, 0xaa, 0x0a, 0x8f, 0xeb, 0x65, 0x54, 0x15, 0xec, 0x1c, 0x0e,
	0x9c, 0x2c, 0x77, 0xe8, 0x41, 0x24, 0x0b, 0x3f, 0xfc, 0x01, 0x58, 0x37, 0x8a, 0xa9, 0x75, 0x65,
	0xf0, 0xef, 0x59, 0xd2, 0x77, 0x0a, 0xec, 0xae, 0xed, 0x28, 0xcc, 0xd7, 0xa1, 0x21, 0x96, 0x43,
	0xff, 0x46, 0xad, 0x54, 0x83, 0xb9, 0x0d, 0x03, 0x76, 0xb1, 0x07, 0x09, 0x6d, 0x8c, 0x87, 0xfb,
	0xc7, 0x21, 0x19, 0xbf, 0x7a, 0xfd, 0xf8, 0x7c, 0xa3, 0x23, 0x7e, 0x26, 0x5c, 0x22, 0x82, 0x9f,
	0xd8, 0xa8, 0xc2, 0x13, 0x97, 0x64, 0x16, 0x91, 0x05, 0x61, 0x8f, 0x70, 0x72, 0x5b, 0x2a, 0xac,
	0x7e, 0x69, 0x31, 0xf1, 0x16, 0x43, 0x3e, 0xe8, 0x58, 0xe4, 0x9e, 0xe7, 0x0d, 0x5a, 0x7c, 0x7b,
	0x1d, 0x6c, 0xfe, 0x0f, 0x6f, 0x3c, 0x6f, 0x49, 0x66, 0x0b, 0xb2, 0x39, 0xf4, 0x6f, 0xe4, 0xfa,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x59, 0x91, 0x2a, 0xe9, 0x66, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloStreamServiceClient is the client API for HelloStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloStreamServiceClient interface {
	BiDirectStream(ctx context.Context, opts ...grpc.CallOption) (HelloStreamService_BiDirectStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloStreamService_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *StreamDataRequest, opts ...grpc.CallOption) (HelloStreamService_ServerStreamClient, error)
}

type helloStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloStreamServiceClient(cc *grpc.ClientConn) HelloStreamServiceClient {
	return &helloStreamServiceClient{cc}
}

func (c *helloStreamServiceClient) BiDirectStream(ctx context.Context, opts ...grpc.CallOption) (HelloStreamService_BiDirectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloStreamService_serviceDesc.Streams[0], "/v1.HelloStreamService/BiDirectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloStreamServiceBiDirectStreamClient{stream}
	return x, nil
}

type HelloStreamService_BiDirectStreamClient interface {
	Send(*StreamDataRequest) error
	Recv() (*StreamDataResponse, error)
	grpc.ClientStream
}

type helloStreamServiceBiDirectStreamClient struct {
	grpc.ClientStream
}

func (x *helloStreamServiceBiDirectStreamClient) Send(m *StreamDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloStreamServiceBiDirectStreamClient) Recv() (*StreamDataResponse, error) {
	m := new(StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloStreamServiceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloStreamService_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloStreamService_serviceDesc.Streams[1], "/v1.HelloStreamService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloStreamServiceClientStreamClient{stream}
	return x, nil
}

type HelloStreamService_ClientStreamClient interface {
	Send(*StreamDataRequest) error
	CloseAndRecv() (*StreamDataResponse, error)
	grpc.ClientStream
}

type helloStreamServiceClientStreamClient struct {
	grpc.ClientStream
}

func (x *helloStreamServiceClientStreamClient) Send(m *StreamDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloStreamServiceClientStreamClient) CloseAndRecv() (*StreamDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloStreamServiceClient) ServerStream(ctx context.Context, in *StreamDataRequest, opts ...grpc.CallOption) (HelloStreamService_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloStreamService_serviceDesc.Streams[2], "/v1.HelloStreamService/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloStreamServiceServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloStreamService_ServerStreamClient interface {
	Recv() (*StreamDataResponse, error)
	grpc.ClientStream
}

type helloStreamServiceServerStreamClient struct {
	grpc.ClientStream
}

func (x *helloStreamServiceServerStreamClient) Recv() (*StreamDataResponse, error) {
	m := new(StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloStreamServiceServer is the server API for HelloStreamService service.
type HelloStreamServiceServer interface {
	BiDirectStream(HelloStreamService_BiDirectStreamServer) error
	ClientStream(HelloStreamService_ClientStreamServer) error
	ServerStream(*StreamDataRequest, HelloStreamService_ServerStreamServer) error
}

// UnimplementedHelloStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloStreamServiceServer struct {
}

func (*UnimplementedHelloStreamServiceServer) BiDirectStream(srv HelloStreamService_BiDirectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectStream not implemented")
}
func (*UnimplementedHelloStreamServiceServer) ClientStream(srv HelloStreamService_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (*UnimplementedHelloStreamServiceServer) ServerStream(req *StreamDataRequest, srv HelloStreamService_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}

func RegisterHelloStreamServiceServer(s *grpc.Server, srv HelloStreamServiceServer) {
	s.RegisterService(&_HelloStreamService_serviceDesc, srv)
}

func _HelloStreamService_BiDirectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloStreamServiceServer).BiDirectStream(&helloStreamServiceBiDirectStreamServer{stream})
}

type HelloStreamService_BiDirectStreamServer interface {
	Send(*StreamDataResponse) error
	Recv() (*StreamDataRequest, error)
	grpc.ServerStream
}

type helloStreamServiceBiDirectStreamServer struct {
	grpc.ServerStream
}

func (x *helloStreamServiceBiDirectStreamServer) Send(m *StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloStreamServiceBiDirectStreamServer) Recv() (*StreamDataRequest, error) {
	m := new(StreamDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloStreamService_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloStreamServiceServer).ClientStream(&helloStreamServiceClientStreamServer{stream})
}

type HelloStreamService_ClientStreamServer interface {
	SendAndClose(*StreamDataResponse) error
	Recv() (*StreamDataRequest, error)
	grpc.ServerStream
}

type helloStreamServiceClientStreamServer struct {
	grpc.ServerStream
}

func (x *helloStreamServiceClientStreamServer) SendAndClose(m *StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloStreamServiceClientStreamServer) Recv() (*StreamDataRequest, error) {
	m := new(StreamDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloStreamService_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloStreamServiceServer).ServerStream(m, &helloStreamServiceServerStreamServer{stream})
}

type HelloStreamService_ServerStreamServer interface {
	Send(*StreamDataResponse) error
	grpc.ServerStream
}

type helloStreamServiceServerStreamServer struct {
	grpc.ServerStream
}

func (x *helloStreamServiceServerStreamServer) Send(m *StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HelloStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.HelloStreamService",
	HandlerType: (*HelloStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BiDirectStream",
			Handler:       _HelloStreamService_BiDirectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _HelloStreamService_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _HelloStreamService_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/hello-stream-service.proto",
}

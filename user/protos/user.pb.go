// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

// protoc --go_out=plugins=grpc:. knowing.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type FindUserByMobileRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByMobileRequest) Reset()         { *m = FindUserByMobileRequest{} }
func (m *FindUserByMobileRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserByMobileRequest) ProtoMessage()    {}
func (*FindUserByMobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *FindUserByMobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByMobileRequest.Unmarshal(m, b)
}
func (m *FindUserByMobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByMobileRequest.Marshal(b, m, deterministic)
}
func (m *FindUserByMobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByMobileRequest.Merge(m, src)
}
func (m *FindUserByMobileRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserByMobileRequest.Size(m)
}
func (m *FindUserByMobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByMobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByMobileRequest proto.InternalMessageInfo

func (m *FindUserByMobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FindIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindIdRequest) Reset()         { *m = FindIdRequest{} }
func (m *FindIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindIdRequest) ProtoMessage()    {}
func (*FindIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *FindIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindIdRequest.Unmarshal(m, b)
}
func (m *FindIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindIdRequest.Marshal(b, m, deterministic)
}
func (m *FindIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindIdRequest.Merge(m, src)
}
func (m *FindIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindIdRequest.Size(m)
}
func (m *FindIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindIdRequest proto.InternalMessageInfo

func (m *FindIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*FindUserByMobileRequest)(nil), "protos.FindUserByMobileRequest")
	proto.RegisterType((*UserResponse)(nil), "protos.UserResponse")
	proto.RegisterType((*FindIdRequest)(nil), "protos.FindIdRequest")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x86, 0x5c, 0xe2, 0x6e,
	0x99, 0x79, 0x29, 0xa1, 0xc5, 0xa9, 0x45, 0x4e, 0x95, 0xbe, 0xf9, 0x49, 0x99, 0x39, 0xa9, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xb9, 0x60, 0x01, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0xc9, 0x88, 0x8b, 0x07, 0xa4, 0x3c, 0x28, 0xb5, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x05, 0xac, 0x86, 0x39, 0x88, 0x29, 0x33,
	0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0xac, 0x0b, 0xcc, 0x56, 0x92,
	0xe7, 0xe2, 0x05, 0x59, 0xe3, 0x99, 0x02, 0x33, 0x1c, 0x4d, 0x93, 0x51, 0x37, 0x23, 0x17, 0x3b,
	0xd8, 0xd4, 0x82, 0x64, 0x21, 0x4f, 0x2e, 0x01, 0x74, 0x37, 0x09, 0xc9, 0x43, 0xdc, 0x5d, 0xac,
	0x87, 0xc3, 0xb5, 0x52, 0x22, 0x30, 0x05, 0x28, 0x6e, 0x33, 0xe5, 0x62, 0x83, 0xd8, 0x2b, 0x24,
	0x8a, 0x6c, 0x00, 0xdc, 0x1d, 0xd8, 0xb5, 0x25, 0x41, 0x42, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0x6d, 0xc1, 0x46, 0x32, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserRpcClient is the client API for UserRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserRpcClient interface {
	FindUserByMobile(ctx context.Context, in *FindUserByMobileRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FindId(ctx context.Context, in *FindIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userRpcClient struct {
	cc *grpc.ClientConn
}

func NewUserRpcClient(cc *grpc.ClientConn) UserRpcClient {
	return &userRpcClient{cc}
}

func (c *userRpcClient) FindUserByMobile(ctx context.Context, in *FindUserByMobileRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/protos.UserRpc/FindUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) FindId(ctx context.Context, in *FindIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/protos.UserRpc/FindId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRpcServer is the server API for UserRpc service.
type UserRpcServer interface {
	FindUserByMobile(context.Context, *FindUserByMobileRequest) (*UserResponse, error)
	FindId(context.Context, *FindIdRequest) (*UserResponse, error)
}

func RegisterUserRpcServer(s *grpc.Server, srv UserRpcServer) {
	s.RegisterService(&_UserRpc_serviceDesc, srv)
}

func _UserRpc_FindUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).FindUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserRpc/FindUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).FindUserByMobile(ctx, req.(*FindUserByMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_FindId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).FindId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserRpc/FindId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).FindId(ctx, req.(*FindIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserRpc",
	HandlerType: (*UserRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserByMobile",
			Handler:    _UserRpc_FindUserByMobile_Handler,
		},
		{
			MethodName: "FindId",
			Handler:    _UserRpc_FindId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

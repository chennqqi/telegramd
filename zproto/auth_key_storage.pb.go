// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_key_storage.proto

/*
Package zproto is a generated protocol buffer package.

It is generated from these files:
	auth_key_storage.proto
	auth_session.proto
	chat_test_data.proto
	message.proto
	rpc_metadata.proto
	sync.proto

It has these top-level messages:
	GetAuthKeyReq
	BoolRsp
	AuthKeyData
	AuthUserReq
	AuthUserRsp
	OnlineStatus
	ChatMessage
	ChatSession
	VoidRsp2
	MessageDataEmpty
	MessageData
	VoidRsp
	AuthKeyMetadata
	RpcMetadata
	ServerAuthReq
	DeliveryUpdatesToUsers
	PushUpdatesData
*/
package zproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// /////////////////////////////////////////////////////////////////////
// ServerAuthReq ==> VoidRsp
// SERVER_AUTH_REQ
type GetAuthKeyReq struct {
	AuthKeyId int64 `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
}

func (m *GetAuthKeyReq) Reset()                    { *m = GetAuthKeyReq{} }
func (m *GetAuthKeyReq) String() string            { return proto.CompactTextString(m) }
func (*GetAuthKeyReq) ProtoMessage()               {}
func (*GetAuthKeyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetAuthKeyReq) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

type BoolRsp struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *BoolRsp) Reset()                    { *m = BoolRsp{} }
func (m *BoolRsp) String() string            { return proto.CompactTextString(m) }
func (*BoolRsp) ProtoMessage()               {}
func (*BoolRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BoolRsp) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AuthKeyData struct {
	AuthKeyId int64  `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	AuthKey   []byte `protobuf:"bytes,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
}

func (m *AuthKeyData) Reset()                    { *m = AuthKeyData{} }
func (m *AuthKeyData) String() string            { return proto.CompactTextString(m) }
func (*AuthKeyData) ProtoMessage()               {}
func (*AuthKeyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthKeyData) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *AuthKeyData) GetAuthKey() []byte {
	if m != nil {
		return m.AuthKey
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAuthKeyReq)(nil), "zproto.GetAuthKeyReq")
	proto.RegisterType((*BoolRsp)(nil), "zproto.BoolRsp")
	proto.RegisterType((*AuthKeyData)(nil), "zproto.AuthKeyData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCAuthKeyStorage service

type RPCAuthKeyStorageClient interface {
	PutAuthKey(ctx context.Context, in *AuthKeyData, opts ...grpc.CallOption) (*VoidRsp, error)
	GetAuthKey(ctx context.Context, in *GetAuthKeyReq, opts ...grpc.CallOption) (*AuthKeyData, error)
}

type rPCAuthKeyStorageClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthKeyStorageClient(cc *grpc.ClientConn) RPCAuthKeyStorageClient {
	return &rPCAuthKeyStorageClient{cc}
}

func (c *rPCAuthKeyStorageClient) PutAuthKey(ctx context.Context, in *AuthKeyData, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/zproto.RPCAuthKeyStorage/PutAuthKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyStorageClient) GetAuthKey(ctx context.Context, in *GetAuthKeyReq, opts ...grpc.CallOption) (*AuthKeyData, error) {
	out := new(AuthKeyData)
	err := grpc.Invoke(ctx, "/zproto.RPCAuthKeyStorage/GetAuthKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCAuthKeyStorage service

type RPCAuthKeyStorageServer interface {
	PutAuthKey(context.Context, *AuthKeyData) (*VoidRsp, error)
	GetAuthKey(context.Context, *GetAuthKeyReq) (*AuthKeyData, error)
}

func RegisterRPCAuthKeyStorageServer(s *grpc.Server, srv RPCAuthKeyStorageServer) {
	s.RegisterService(&_RPCAuthKeyStorage_serviceDesc, srv)
}

func _RPCAuthKeyStorage_PutAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthKeyData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyStorageServer).PutAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zproto.RPCAuthKeyStorage/PutAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyStorageServer).PutAuthKey(ctx, req.(*AuthKeyData))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKeyStorage_GetAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyStorageServer).GetAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zproto.RPCAuthKeyStorage/GetAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyStorageServer).GetAuthKey(ctx, req.(*GetAuthKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCAuthKeyStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zproto.RPCAuthKeyStorage",
	HandlerType: (*RPCAuthKeyStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutAuthKey",
			Handler:    _RPCAuthKeyStorage_PutAuthKey_Handler,
		},
		{
			MethodName: "GetAuthKey",
			Handler:    _RPCAuthKeyStorage_GetAuthKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_key_storage.proto",
}

func init() { proto.RegisterFile("auth_key_storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0x2d, 0xc9,
	0x88, 0xcf, 0x4e, 0xad, 0x8c, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0xab, 0x02, 0xd3, 0x52, 0xbc, 0xb9, 0xa9, 0xc5, 0xc5, 0x70, 0x61, 0x25,
	0x7d, 0x2e, 0x5e, 0xf7, 0xd4, 0x12, 0xc7, 0xd2, 0x92, 0x0c, 0xef, 0xd4, 0xca, 0xa0, 0xd4, 0x42,
	0x21, 0x39, 0x2e, 0x6e, 0xb8, 0x09, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x9c,
	0x89, 0x10, 0x05, 0x9e, 0x29, 0x4a, 0x8a, 0x5c, 0xec, 0x4e, 0xf9, 0xf9, 0x39, 0x41, 0xc5, 0x05,
	0x42, 0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x60, 0x55, 0xac, 0x41, 0x50, 0x9e,
	0x92, 0x07, 0x17, 0x37, 0xd4, 0x40, 0x97, 0xc4, 0x92, 0x44, 0x42, 0x26, 0x0a, 0x49, 0x72, 0x71,
	0xc0, 0xe4, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0xd8, 0xa1, 0x92, 0x46, 0x8d, 0x8c, 0x5c,
	0x82, 0x41, 0x01, 0xce, 0x50, 0xd3, 0x82, 0x21, 0x1e, 0x12, 0x32, 0xe2, 0xe2, 0x0a, 0x28, 0x85,
	0xb9, 0x59, 0x48, 0x58, 0x0f, 0xe2, 0x33, 0x3d, 0x24, 0x3b, 0xa5, 0xf8, 0x61, 0x82, 0x61, 0xf9,
	0x99, 0x29, 0x20, 0xb7, 0x5a, 0x70, 0x71, 0x21, 0xfc, 0x29, 0x24, 0x0a, 0x93, 0x46, 0xf1, 0xbb,
	0x14, 0x36, 0xa3, 0x9c, 0xf4, 0xb9, 0x84, 0x93, 0xf3, 0x73, 0xf5, 0xf2, 0x52, 0x93, 0x4a, 0x73,
	0x12, 0x33, 0x73, 0xa1, 0x4a, 0x9c, 0x44, 0xa2, 0x02, 0x40, 0x34, 0xaa, 0xd3, 0x3c, 0x98, 0x02,
	0x18, 0x93, 0xd8, 0xc0, 0x0a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x9b, 0x93, 0x5f,
	0x8a, 0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ucenter.proto

package pb

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

type ReqCreateBranch struct {
}

func (m *ReqCreateBranch) Reset()                    { *m = ReqCreateBranch{} }
func (m *ReqCreateBranch) String() string            { return proto.CompactTextString(m) }
func (*ReqCreateBranch) ProtoMessage()               {}
func (*ReqCreateBranch) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type ReqDeleteBranch struct {
}

func (m *ReqDeleteBranch) Reset()                    { *m = ReqDeleteBranch{} }
func (m *ReqDeleteBranch) String() string            { return proto.CompactTextString(m) }
func (*ReqDeleteBranch) ProtoMessage()               {}
func (*ReqDeleteBranch) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

type ReqCreateUser struct {
}

func (m *ReqCreateUser) Reset()                    { *m = ReqCreateUser{} }
func (m *ReqCreateUser) String() string            { return proto.CompactTextString(m) }
func (*ReqCreateUser) ProtoMessage()               {}
func (*ReqCreateUser) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

type ReqDeleteUser struct {
}

func (m *ReqDeleteUser) Reset()                    { *m = ReqDeleteUser{} }
func (m *ReqDeleteUser) String() string            { return proto.CompactTextString(m) }
func (*ReqDeleteUser) ProtoMessage()               {}
func (*ReqDeleteUser) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

type ReqCheckAPIPermission struct {
	Type UserType `protobuf:"varint,1,opt,name=type,enum=pb.UserType" json:"type,omitempty"`
	Api  string   `protobuf:"bytes,2,opt,name=api" json:"api,omitempty"`
}

func (m *ReqCheckAPIPermission) Reset()                    { *m = ReqCheckAPIPermission{} }
func (m *ReqCheckAPIPermission) String() string            { return proto.CompactTextString(m) }
func (*ReqCheckAPIPermission) ProtoMessage()               {}
func (*ReqCheckAPIPermission) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *ReqCheckAPIPermission) GetType() UserType {
	if m != nil {
		return m.Type
	}
	return UserType_UT_NORMAL
}

func (m *ReqCheckAPIPermission) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type RspCheckAPIPermission struct {
	Passed bool `protobuf:"varint,1,opt,name=passed" json:"passed,omitempty"`
}

func (m *RspCheckAPIPermission) Reset()                    { *m = RspCheckAPIPermission{} }
func (m *RspCheckAPIPermission) String() string            { return proto.CompactTextString(m) }
func (*RspCheckAPIPermission) ProtoMessage()               {}
func (*RspCheckAPIPermission) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *RspCheckAPIPermission) GetPassed() bool {
	if m != nil {
		return m.Passed
	}
	return false
}

type ReqCheckResourcePermission struct {
	Type     UserType `protobuf:"varint,1,opt,name=type,enum=pb.UserType" json:"type,omitempty"`
	Resource int64    `protobuf:"varint,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *ReqCheckResourcePermission) Reset()                    { *m = ReqCheckResourcePermission{} }
func (m *ReqCheckResourcePermission) String() string            { return proto.CompactTextString(m) }
func (*ReqCheckResourcePermission) ProtoMessage()               {}
func (*ReqCheckResourcePermission) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *ReqCheckResourcePermission) GetType() UserType {
	if m != nil {
		return m.Type
	}
	return UserType_UT_NORMAL
}

func (m *ReqCheckResourcePermission) GetResource() int64 {
	if m != nil {
		return m.Resource
	}
	return 0
}

type RspCheckResourcePermission struct {
	Passed bool `protobuf:"varint,1,opt,name=passed" json:"passed,omitempty"`
}

func (m *RspCheckResourcePermission) Reset()                    { *m = RspCheckResourcePermission{} }
func (m *RspCheckResourcePermission) String() string            { return proto.CompactTextString(m) }
func (*RspCheckResourcePermission) ProtoMessage()               {}
func (*RspCheckResourcePermission) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{7} }

func (m *RspCheckResourcePermission) GetPassed() bool {
	if m != nil {
		return m.Passed
	}
	return false
}

type ReqUserLogin struct {
	User     string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *ReqUserLogin) Reset()                    { *m = ReqUserLogin{} }
func (m *ReqUserLogin) String() string            { return proto.CompactTextString(m) }
func (*ReqUserLogin) ProtoMessage()               {}
func (*ReqUserLogin) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{8} }

func (m *ReqUserLogin) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqUserLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ReqUserLogin) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type RspUserLogin struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *RspUserLogin) Reset()                    { *m = RspUserLogin{} }
func (m *RspUserLogin) String() string            { return proto.CompactTextString(m) }
func (*RspUserLogin) ProtoMessage()               {}
func (*RspUserLogin) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{9} }

func (m *RspUserLogin) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RspUserLogin) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ReqCreateBranch)(nil), "pb.ReqCreateBranch")
	proto.RegisterType((*ReqDeleteBranch)(nil), "pb.ReqDeleteBranch")
	proto.RegisterType((*ReqCreateUser)(nil), "pb.ReqCreateUser")
	proto.RegisterType((*ReqDeleteUser)(nil), "pb.ReqDeleteUser")
	proto.RegisterType((*ReqCheckAPIPermission)(nil), "pb.ReqCheckAPIPermission")
	proto.RegisterType((*RspCheckAPIPermission)(nil), "pb.RspCheckAPIPermission")
	proto.RegisterType((*ReqCheckResourcePermission)(nil), "pb.ReqCheckResourcePermission")
	proto.RegisterType((*RspCheckResourcePermission)(nil), "pb.RspCheckResourcePermission")
	proto.RegisterType((*ReqUserLogin)(nil), "pb.ReqUserLogin")
	proto.RegisterType((*RspUserLogin)(nil), "pb.RspUserLogin")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UCenter service

type UCenterClient interface {
	// UserLogin 登陆
	UserLogin(ctx context.Context, in *ReqUserLogin, opts ...grpc.CallOption) (*RspUserLogin, error)
	// UserLogout 登出
	UserLogout(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// 创建用户
	CreateUser(ctx context.Context, in *ReqCreateUser, opts ...grpc.CallOption) (*CommonResponse, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *ReqDeleteUser, opts ...grpc.CallOption) (*CommonResponse, error)
	// 创建部门
	CreateBranch(ctx context.Context, in *ReqCreateBranch, opts ...grpc.CallOption) (*CommonResponse, error)
	// 删除部门
	DeleteBranch(ctx context.Context, in *ReqDeleteBranch, opts ...grpc.CallOption) (*CommonResponse, error)
	// 检查权限
	CheckAPIPermission(ctx context.Context, in *ReqCheckAPIPermission, opts ...grpc.CallOption) (*RspCheckAPIPermission, error)
	// 检查资源权限
	CheckResourcePermission(ctx context.Context, in *ReqCheckResourcePermission, opts ...grpc.CallOption) (*RspCheckResourcePermission, error)
}

type uCenterClient struct {
	cc *grpc.ClientConn
}

func NewUCenterClient(cc *grpc.ClientConn) UCenterClient {
	return &uCenterClient{cc}
}

func (c *uCenterClient) UserLogin(ctx context.Context, in *ReqUserLogin, opts ...grpc.CallOption) (*RspUserLogin, error) {
	out := new(RspUserLogin)
	err := grpc.Invoke(ctx, "/pb.UCenter/UserLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) UserLogout(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/pb.UCenter/UserLogout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) CreateUser(ctx context.Context, in *ReqCreateUser, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/pb.UCenter/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) DeleteUser(ctx context.Context, in *ReqDeleteUser, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/pb.UCenter/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) CreateBranch(ctx context.Context, in *ReqCreateBranch, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/pb.UCenter/CreateBranch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) DeleteBranch(ctx context.Context, in *ReqDeleteBranch, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/pb.UCenter/DeleteBranch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) CheckAPIPermission(ctx context.Context, in *ReqCheckAPIPermission, opts ...grpc.CallOption) (*RspCheckAPIPermission, error) {
	out := new(RspCheckAPIPermission)
	err := grpc.Invoke(ctx, "/pb.UCenter/CheckAPIPermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCenterClient) CheckResourcePermission(ctx context.Context, in *ReqCheckResourcePermission, opts ...grpc.CallOption) (*RspCheckResourcePermission, error) {
	out := new(RspCheckResourcePermission)
	err := grpc.Invoke(ctx, "/pb.UCenter/CheckResourcePermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UCenter service

type UCenterServer interface {
	// UserLogin 登陆
	UserLogin(context.Context, *ReqUserLogin) (*RspUserLogin, error)
	// UserLogout 登出
	UserLogout(context.Context, *EmptyRequest) (*EmptyResponse, error)
	// 创建用户
	CreateUser(context.Context, *ReqCreateUser) (*CommonResponse, error)
	// 删除用户
	DeleteUser(context.Context, *ReqDeleteUser) (*CommonResponse, error)
	// 创建部门
	CreateBranch(context.Context, *ReqCreateBranch) (*CommonResponse, error)
	// 删除部门
	DeleteBranch(context.Context, *ReqDeleteBranch) (*CommonResponse, error)
	// 检查权限
	CheckAPIPermission(context.Context, *ReqCheckAPIPermission) (*RspCheckAPIPermission, error)
	// 检查资源权限
	CheckResourcePermission(context.Context, *ReqCheckResourcePermission) (*RspCheckResourcePermission, error)
}

func RegisterUCenterServer(s *grpc.Server, srv UCenterServer) {
	s.RegisterService(&_UCenter_serviceDesc, srv)
}

func _UCenter_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).UserLogin(ctx, req.(*ReqUserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_UserLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).UserLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/UserLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).UserLogout(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).CreateUser(ctx, req.(*ReqCreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).DeleteUser(ctx, req.(*ReqDeleteUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateBranch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/CreateBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).CreateBranch(ctx, req.(*ReqCreateBranch))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_DeleteBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteBranch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).DeleteBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/DeleteBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).DeleteBranch(ctx, req.(*ReqDeleteBranch))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_CheckAPIPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCheckAPIPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).CheckAPIPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/CheckAPIPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).CheckAPIPermission(ctx, req.(*ReqCheckAPIPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCenter_CheckResourcePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCheckResourcePermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCenterServer).CheckResourcePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UCenter/CheckResourcePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCenterServer).CheckResourcePermission(ctx, req.(*ReqCheckResourcePermission))
	}
	return interceptor(ctx, in, info, handler)
}

var _UCenter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UCenter",
	HandlerType: (*UCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UCenter_UserLogin_Handler,
		},
		{
			MethodName: "UserLogout",
			Handler:    _UCenter_UserLogout_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UCenter_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UCenter_DeleteUser_Handler,
		},
		{
			MethodName: "CreateBranch",
			Handler:    _UCenter_CreateBranch_Handler,
		},
		{
			MethodName: "DeleteBranch",
			Handler:    _UCenter_DeleteBranch_Handler,
		},
		{
			MethodName: "CheckAPIPermission",
			Handler:    _UCenter_CheckAPIPermission_Handler,
		},
		{
			MethodName: "CheckResourcePermission",
			Handler:    _UCenter_CheckResourcePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ucenter.proto",
}

func init() { proto.RegisterFile("ucenter.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x8f, 0xda, 0x30,
	0x10, 0x15, 0x9b, 0xd5, 0x2e, 0x8c, 0xb2, 0xdd, 0xc5, 0xd5, 0xb6, 0xac, 0x0f, 0x2b, 0x94, 0x13,
	0x27, 0x50, 0xa1, 0x52, 0xcf, 0x85, 0xf6, 0xd0, 0x8f, 0x03, 0xb2, 0xca, 0xa1, 0xbd, 0x85, 0x30,
	0x82, 0xa8, 0x4d, 0x6c, 0x3c, 0x89, 0x2a, 0xfe, 0x75, 0x7f, 0x42, 0x65, 0xc7, 0x49, 0x5c, 0x11,
	0x0e, 0xbd, 0x79, 0xde, 0xcc, 0x7b, 0x6f, 0x32, 0x7a, 0x81, 0xbb, 0x32, 0xc1, 0xbc, 0x40, 0x3d,
	0x55, 0x5a, 0x16, 0x92, 0x5d, 0xa9, 0x2d, 0x0f, 0x13, 0x99, 0x65, 0x32, 0xaf, 0x10, 0x0e, 0x25,
	0xd5, 0xdd, 0x68, 0x08, 0xf7, 0x02, 0x8f, 0x2b, 0x8d, 0x71, 0x81, 0x4b, 0x1d, 0xe7, 0xc9, 0xc1,
	0x41, 0x1f, 0xf0, 0x17, 0x36, 0xd0, 0x3d, 0xdc, 0x35, 0x53, 0x1b, 0x42, 0xed, 0x80, 0x6a, 0xc6,
	0x02, 0x5f, 0xe0, 0xd1, 0x4c, 0x1c, 0x30, 0xf9, 0xf9, 0x7e, 0xfd, 0x69, 0x8d, 0x3a, 0x4b, 0x89,
	0x52, 0x99, 0xb3, 0x31, 0x5c, 0x17, 0x27, 0x85, 0xa3, 0xde, 0xb8, 0x37, 0x79, 0x31, 0x0f, 0xa7,
	0x6a, 0x3b, 0x35, 0x84, 0x6f, 0x27, 0x85, 0xc2, 0x76, 0xd8, 0x03, 0x04, 0xb1, 0x4a, 0x47, 0x57,
	0xe3, 0xde, 0x64, 0x20, 0xcc, 0x33, 0x9a, 0xc1, 0xa3, 0x20, 0xd5, 0x21, 0xf6, 0x0a, 0x6e, 0x54,
	0x4c, 0x84, 0x3b, 0x2b, 0xd7, 0x17, 0xae, 0x8a, 0x7e, 0x00, 0xaf, 0xdd, 0x05, 0x92, 0x2c, 0x75,
	0x82, 0xff, 0xb5, 0x02, 0x87, 0xbe, 0x76, 0x3c, 0xbb, 0x47, 0x20, 0x9a, 0x3a, 0x7a, 0x0b, 0xbc,
	0x5e, 0xa6, 0x43, 0xfb, 0xd2, 0x46, 0x02, 0x42, 0x81, 0x47, 0x63, 0xf3, 0x55, 0xee, 0xd3, 0x9c,
	0x31, 0xb8, 0x36, 0x57, 0xb7, 0x53, 0x03, 0x61, 0xdf, 0xc6, 0xd5, 0x4c, 0xff, 0x96, 0x7a, 0xe7,
	0xbe, 0xbe, 0xa9, 0xcd, 0xbc, 0xdd, 0x39, 0xa8, 0xe6, 0xcd, 0x3b, 0x5a, 0x42, 0x28, 0x48, 0xb5,
	0x9a, 0x23, 0xb8, 0xcd, 0x90, 0x28, 0xde, 0xa3, 0x93, 0xad, 0x4b, 0xd3, 0xa1, 0x32, 0x49, 0x90,
	0xc8, 0x0a, 0xf7, 0x45, 0x5d, 0xce, 0xff, 0x04, 0x70, 0xbb, 0x59, 0xd9, 0x7c, 0xb0, 0x19, 0x0c,
	0x5a, 0xb1, 0x07, 0x73, 0x16, 0x7f, 0x65, 0x5e, 0x21, 0xbe, 0xe1, 0x1b, 0x00, 0x57, 0xc8, 0xb2,
	0xa8, 0x18, 0x1f, 0x33, 0x55, 0x9c, 0x04, 0x1e, 0x4b, 0xa4, 0x82, 0x0f, 0x3d, 0x84, 0x94, 0xcc,
	0x09, 0xd9, 0x02, 0xa0, 0x8d, 0x0d, 0x1b, 0x3a, 0x93, 0x16, 0xe2, 0xcc, 0x40, 0x2b, 0x1b, 0x4f,
	0x9f, 0xd4, 0x46, 0xab, 0x21, 0xb5, 0x50, 0x27, 0xe9, 0x1d, 0x84, 0x7e, 0x8c, 0xd9, 0xcb, 0x7f,
	0xbc, 0x2a, 0xf0, 0x12, 0xd1, 0x0f, 0x7b, 0x43, 0xf4, 0xc1, 0x4e, 0xe2, 0x67, 0x60, 0x1d, 0x19,
	0x7d, 0xaa, 0x7d, 0xcf, 0x5a, 0xfc, 0xc9, 0x5d, 0xb4, 0x83, 0xf5, 0x1d, 0x5e, 0x5f, 0x8a, 0xd8,
	0xb3, 0x2f, 0x78, 0xde, 0xe7, 0xcf, 0xbe, 0xea, 0x79, 0x7f, 0x7b, 0x63, 0xff, 0xf4, 0xc5, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xb0, 0x49, 0xa8, 0x18, 0x04, 0x00, 0x00,
}
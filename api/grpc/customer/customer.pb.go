// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package customer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type RegisterRequest struct {
	EmailAddress         string   `protobuf:"bytes,1,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	GivenName            string   `protobuf:"bytes,2,opt,name=givenName,proto3" json:"givenName,omitempty"`
	FamilyName           string   `protobuf:"bytes,3,opt,name=familyName,proto3" json:"familyName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *RegisterRequest) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *RegisterRequest) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

type RegisterResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ConfirmEmailAddressRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	ConfirmationHash     string   `protobuf:"bytes,3,opt,name=confirmationHash,proto3" json:"confirmationHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmEmailAddressRequest) Reset()         { *m = ConfirmEmailAddressRequest{} }
func (m *ConfirmEmailAddressRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmEmailAddressRequest) ProtoMessage()    {}
func (*ConfirmEmailAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *ConfirmEmailAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmEmailAddressRequest.Unmarshal(m, b)
}
func (m *ConfirmEmailAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmEmailAddressRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmEmailAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmEmailAddressRequest.Merge(m, src)
}
func (m *ConfirmEmailAddressRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmEmailAddressRequest.Size(m)
}
func (m *ConfirmEmailAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmEmailAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmEmailAddressRequest proto.InternalMessageInfo

func (m *ConfirmEmailAddressRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConfirmEmailAddressRequest) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *ConfirmEmailAddressRequest) GetConfirmationHash() string {
	if m != nil {
		return m.ConfirmationHash
	}
	return ""
}

type ChangeEmailAddressRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress         string   `protobuf:"bytes,2,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeEmailAddressRequest) Reset()         { *m = ChangeEmailAddressRequest{} }
func (m *ChangeEmailAddressRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeEmailAddressRequest) ProtoMessage()    {}
func (*ChangeEmailAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *ChangeEmailAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEmailAddressRequest.Unmarshal(m, b)
}
func (m *ChangeEmailAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEmailAddressRequest.Marshal(b, m, deterministic)
}
func (m *ChangeEmailAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEmailAddressRequest.Merge(m, src)
}
func (m *ChangeEmailAddressRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeEmailAddressRequest.Size(m)
}
func (m *ChangeEmailAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEmailAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEmailAddressRequest proto.InternalMessageInfo

func (m *ChangeEmailAddressRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChangeEmailAddressRequest) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "customer.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "customer.RegisterResponse")
	proto.RegisterType((*ConfirmEmailAddressRequest)(nil), "customer.ConfirmEmailAddressRequest")
	proto.RegisterType((*ChangeEmailAddressRequest)(nil), "customer.ChangeEmailAddressRequest")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xd3, 0x7e, 0xc9, 0x17, 0xb8, 0x21, 0x88, 0x63, 0xa2, 0x30, 0x12, 0x62, 0x06, 0x63,
	0x14, 0x93, 0x4e, 0xd0, 0x1d, 0x3b, 0x43, 0x48, 0x5c, 0x69, 0xc2, 0xd2, 0xdd, 0x40, 0x87, 0x32,
	0x09, 0x9d, 0xa9, 0x9d, 0x81, 0x84, 0x88, 0x1b, 0x17, 0xbe, 0x80, 0x8f, 0xe6, 0x2b, 0xb8, 0xf5,
	0x1d, 0x0c, 0xd3, 0x56, 0x2a, 0x7f, 0x5c, 0xb9, 0xec, 0x3d, 0xb7, 0xf7, 0xfc, 0xe6, 0xdc, 0x0b,
	0xe5, 0xe1, 0x54, 0x1b, 0x15, 0xf2, 0xd8, 0x8b, 0x62, 0x65, 0x14, 0x2a, 0x64, 0xdf, 0xf8, 0x38,
	0x50, 0x2a, 0x98, 0x70, 0x6a, 0xeb, 0x83, 0xe9, 0x88, 0xf2, 0x30, 0x32, 0xf3, 0xa4, 0x0d, 0xd7,
	0x53, 0x91, 0x45, 0x82, 0x32, 0x29, 0x95, 0x61, 0x46, 0x28, 0xa9, 0x13, 0x95, 0x68, 0xd8, 0xeb,
	0xf3, 0x40, 0x68, 0xc3, 0xe3, 0x3e, 0x7f, 0x9c, 0x72, 0x6d, 0x10, 0x81, 0x12, 0x0f, 0x99, 0x98,
	0xdc, 0xf8, 0x7e, 0xcc, 0xb5, 0xae, 0x3a, 0x27, 0xce, 0x79, 0xb1, 0xff, 0xa3, 0x86, 0xea, 0x50,
	0x0c, 0xc4, 0x8c, 0xcb, 0x3b, 0x16, 0xf2, 0xaa, 0x6b, 0x1b, 0x56, 0x05, 0xd4, 0x00, 0x18, 0xb1,
	0x50, 0x4c, 0xe6, 0x56, 0xfe, 0x67, 0xe5, 0x5c, 0x85, 0x10, 0xa8, 0xac, 0x4c, 0x75, 0xa4, 0xa4,
	0xe6, 0xa8, 0x0c, 0xae, 0xf0, 0x53, 0x2f, 0x57, 0xf8, 0x64, 0x01, 0xb8, 0xab, 0xe4, 0x48, 0xc4,
	0x61, 0x2f, 0x67, 0x9c, 0x31, 0xae, 0x75, 0x6f, 0x30, 0xbb, 0x5b, 0x98, 0x5b, 0x50, 0x19, 0x26,
	0x13, 0x6d, 0x02, 0xb7, 0x4c, 0x8f, 0x53, 0xb6, 0x8d, 0x3a, 0xb9, 0x87, 0x5a, 0x77, 0xcc, 0x64,
	0xc0, 0xff, 0xc8, 0xfc, 0xea, 0xd3, 0x85, 0x42, 0x37, 0xdd, 0x17, 0x7a, 0x80, 0x42, 0xf6, 0x7e,
	0x54, 0xf3, 0xbe, 0xd7, 0xba, 0xb6, 0x08, 0x8c, 0xb7, 0x49, 0x49, 0x5c, 0xe4, 0xe8, 0xe5, 0xfd,
	0xe3, 0xcd, 0xdd, 0x27, 0x25, 0x3a, 0x6b, 0xd3, 0xac, 0xad, 0xe3, 0xb4, 0xd0, 0xab, 0x03, 0x07,
	0x5b, 0x82, 0x43, 0xa7, 0xab, 0x61, 0xbb, 0x73, 0xc5, 0x87, 0x5e, 0x72, 0x2d, 0x5e, 0x76, 0x4a,
	0x5e, 0x6f, 0x79, 0x4a, 0xa4, 0x6d, 0xed, 0x2e, 0xf1, 0x59, 0xde, 0x8e, 0x3e, 0x09, 0xff, 0x99,
	0xda, 0x67, 0xb2, 0x64, 0x0c, 0x4d, 0x83, 0x5c, 0x82, 0x2c, 0x00, 0x6d, 0x46, 0x88, 0x9a, 0x39,
	0x8c, 0x5d, 0x01, 0xef, 0xa4, 0xb8, 0xb0, 0x14, 0x4d, 0xdc, 0xf8, 0x9d, 0xa2, 0xe3, 0xb4, 0x06,
	0xff, 0xed, 0xaf, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x66, 0x56, 0xad, 0x35, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	ConfirmEmailAddress(ctx context.Context, in *ConfirmEmailAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ChangeEmailAddress(ctx context.Context, in *ChangeEmailAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/customer.Customer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) ConfirmEmailAddress(ctx context.Context, in *ConfirmEmailAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/customer.Customer/ConfirmEmailAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) ChangeEmailAddress(ctx context.Context, in *ChangeEmailAddressRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/customer.Customer/ChangeEmailAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	ConfirmEmailAddress(context.Context, *ConfirmEmailAddressRequest) (*empty.Empty, error)
	ChangeEmailAddress(context.Context, *ChangeEmailAddressRequest) (*empty.Empty, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedCustomerServer) ConfirmEmailAddress(ctx context.Context, req *ConfirmEmailAddressRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmEmailAddress not implemented")
}
func (*UnimplementedCustomerServer) ChangeEmailAddress(ctx context.Context, req *ChangeEmailAddressRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEmailAddress not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_ConfirmEmailAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmEmailAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).ConfirmEmailAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/ConfirmEmailAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).ConfirmEmailAddress(ctx, req.(*ConfirmEmailAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_ChangeEmailAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEmailAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).ChangeEmailAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/ChangeEmailAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).ChangeEmailAddress(ctx, req.(*ChangeEmailAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Customer_Register_Handler,
		},
		{
			MethodName: "ConfirmEmailAddress",
			Handler:    _Customer_ConfirmEmailAddress_Handler,
		},
		{
			MethodName: "ChangeEmailAddress",
			Handler:    _Customer_ChangeEmailAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}

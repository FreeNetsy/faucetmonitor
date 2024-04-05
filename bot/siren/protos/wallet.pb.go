// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/wallet.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetAddressRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressRequest) Reset()         { *m = GetAddressRequest{} }
func (m *GetAddressRequest) String() string { return proto.CompactTextString(m) }
func (*GetAddressRequest) ProtoMessage()    {}
func (*GetAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b906fab2aed9696f, []int{0}
}

func (m *GetAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressRequest.Unmarshal(m, b)
}
func (m *GetAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressRequest.Marshal(b, m, deterministic)
}
func (m *GetAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressRequest.Merge(m, src)
}
func (m *GetAddressRequest) XXX_Size() int {
	return xxx_messageInfo_GetAddressRequest.Size(m)
}
func (m *GetAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressRequest proto.InternalMessageInfo

func (m *GetAddressRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetAddressRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type GetAddressResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressResponse) Reset()         { *m = GetAddressResponse{} }
func (m *GetAddressResponse) String() string { return proto.CompactTextString(m) }
func (*GetAddressResponse) ProtoMessage()    {}
func (*GetAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b906fab2aed9696f, []int{1}
}

func (m *GetAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressResponse.Unmarshal(m, b)
}
func (m *GetAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressResponse.Marshal(b, m, deterministic)
}
func (m *GetAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressResponse.Merge(m, src)
}
func (m *GetAddressResponse) XXX_Size() int {
	return xxx_messageInfo_GetAddressResponse.Size(m)
}
func (m *GetAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressResponse proto.InternalMessageInfo

func (m *GetAddressResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type WithdrawRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawRequest) Reset()         { *m = WithdrawRequest{} }
func (m *WithdrawRequest) String() string { return proto.CompactTextString(m) }
func (*WithdrawRequest) ProtoMessage()    {}
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b906fab2aed9696f, []int{2}
}

func (m *WithdrawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawRequest.Unmarshal(m, b)
}
func (m *WithdrawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawRequest.Marshal(b, m, deterministic)
}
func (m *WithdrawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawRequest.Merge(m, src)
}
func (m *WithdrawRequest) XXX_Size() int {
	return xxx_messageInfo_WithdrawRequest.Size(m)
}
func (m *WithdrawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawRequest proto.InternalMessageInfo

func (m *WithdrawRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *WithdrawRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *WithdrawRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *WithdrawRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type WithdrawResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawResponse) Reset()         { *m = WithdrawResponse{} }
func (m *WithdrawResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawResponse) ProtoMessage()    {}
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b906fab2aed9696f, []int{3}
}

func (m *WithdrawResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawResponse.Unmarshal(m, b)
}
func (m *WithdrawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawResponse.Merge(m, src)
}
func (m *WithdrawResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawResponse.Size(m)
}
func (m *WithdrawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetAddressRequest)(nil), "protos.GetAddressRequest")
	proto.RegisterType((*GetAddressResponse)(nil), "protos.GetAddressResponse")
	proto.RegisterType((*WithdrawRequest)(nil), "protos.WithdrawRequest")
	proto.RegisterType((*WithdrawResponse)(nil), "protos.WithdrawResponse")
}

func init() { proto.RegisterFile("protos/wallet.proto", fileDescriptor_b906fab2aed9696f) }

var fileDescriptor_b906fab2aed9696f = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0x7e, 0xf3, 0x82, 0x92, 0x70, 0x0b, 0x70, 0x48, 0x34, 0x64, 0x42, 0x11, 0x03, 0x0b, 0x89,
	0x54, 0x66, 0x06, 0x8a, 0x2a, 0x60, 0x83, 0x2e, 0x95, 0x58, 0x50, 0x1a, 0x9f, 0x88, 0xa5, 0x26,
	0x0e, 0xf6, 0x59, 0x05, 0x7e, 0x03, 0x3f, 0x1a, 0x11, 0x27, 0xb4, 0x50, 0x36, 0x3a, 0x3e, 0x77,
	0xbe, 0xe7, 0x4b, 0x86, 0x83, 0x46, 0x2b, 0x56, 0x26, 0x5b, 0xe4, 0xf3, 0x39, 0x71, 0xda, 0x22,
	0xf4, 0xdd, 0x30, 0xb9, 0x81, 0xfd, 0x6b, 0xe2, 0x4b, 0x21, 0x34, 0x19, 0x33, 0xa1, 0x67, 0x4b,
	0x86, 0x71, 0x00, 0x81, 0x35, 0xa4, 0x1f, 0xa5, 0x88, 0xbc, 0x63, 0xef, 0x74, 0x67, 0xe2, 0x7f,
	0xc2, 0x5b, 0x81, 0x31, 0x84, 0x85, 0xd5, 0x9a, 0xea, 0xe2, 0x35, 0xfa, 0xdf, 0x6e, 0xbe, 0x70,
	0x92, 0x02, 0xae, 0x32, 0x99, 0x46, 0xd5, 0x86, 0x30, 0x82, 0x20, 0x77, 0xa3, 0x8e, 0xaa, 0x87,
	0xc9, 0x0b, 0xec, 0x4e, 0x25, 0x97, 0x42, 0xe7, 0x8b, 0xbf, 0xe8, 0xe2, 0x21, 0xf8, 0x79, 0xa5,
	0x6c, 0xcd, 0xd1, 0x96, 0xbb, 0x71, 0x68, 0x55, 0x79, 0xfb, 0xbb, 0x32, 0xc2, 0xde, 0x52, 0xd9,
	0xf9, 0x1c, 0xbe, 0x7b, 0xe0, 0x4f, 0xdb, 0x82, 0xf0, 0x0a, 0x60, 0x19, 0x04, 0x8f, 0x5c, 0x61,
	0x26, 0x5d, 0xab, 0x29, 0x8e, 0x7f, 0x5b, 0x75, 0xb9, 0x2f, 0x20, 0xec, 0x35, 0x70, 0xd0, 0xbf,
	0xfb, 0x91, 0x37, 0x8e, 0xd6, 0x17, 0x9d, 0x9d, 0x7b, 0x08, 0xc7, 0x5c, 0x92, 0x26, 0x5b, 0xe1,
	0x78, 0x03, 0x7e, 0x92, 0x7f, 0xc3, 0x3b, 0x08, 0x46, 0x92, 0x0b, 0x25, 0xeb, 0x0d, 0x31, 0x8e,
	0x4e, 0x1e, 0x92, 0x27, 0xc9, 0xa5, 0x9d, 0xa5, 0x85, 0xaa, 0x32, 0xf1, 0xc6, 0x54, 0x94, 0x99,
	0x91, 0x9a, 0xea, 0xb3, 0x46, 0x99, 0xcc, 0x9d, 0xce, 0xdc, 0x4f, 0x3b, 0xff, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x54, 0xa1, 0x00, 0x9d, 0x87, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletClient interface {
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
}

type walletClient struct {
	cc *grpc.ClientConn
}

func NewWalletClient(cc *grpc.ClientConn) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/protos.Wallet/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, "/protos.Wallet/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
type WalletServer interface {
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
}

func RegisterWalletServer(s *grpc.Server, srv WalletServer) {
	s.RegisterService(&_Wallet_serviceDesc, srv)
}

func _Wallet_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Wallet/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Wallet/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddress",
			Handler:    _Wallet_GetAddress_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Wallet_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/wallet.proto",
}

// EthereumClient is the client API for Ethereum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EthereumClient interface {
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
}

type ethereumClient struct {
	cc *grpc.ClientConn
}

func NewEthereumClient(cc *grpc.ClientConn) EthereumClient {
	return &ethereumClient{cc}
}

func (c *ethereumClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/protos.Ethereum/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EthereumServer is the server API for Ethereum service.
type EthereumServer interface {
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
}

func RegisterEthereumServer(s *grpc.Server, srv EthereumServer) {
	s.RegisterService(&_Ethereum_serviceDesc, srv)
}

func _Ethereum_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Ethereum/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ethereum_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Ethereum",
	HandlerType: (*EthereumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddress",
			Handler:    _Ethereum_GetAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/wallet.proto",
}

// BitcoinClient is the client API for Bitcoin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BitcoinClient interface {
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
}

type bitcoinClient struct {
	cc *grpc.ClientConn
}

func NewBitcoinClient(cc *grpc.ClientConn) BitcoinClient {
	return &bitcoinClient{cc}
}

func (c *bitcoinClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/protos.Bitcoin/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BitcoinServer is the server API for Bitcoin service.
type BitcoinServer interface {
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
}

func RegisterBitcoinServer(s *grpc.Server, srv BitcoinServer) {
	s.RegisterService(&_Bitcoin_serviceDesc, srv)
}

func _Bitcoin_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Bitcoin/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bitcoin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Bitcoin",
	HandlerType: (*BitcoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddress",
			Handler:    _Bitcoin_GetAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/wallet.proto",
}

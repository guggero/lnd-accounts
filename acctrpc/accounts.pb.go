// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package acctrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateAccountRequest struct {
	//*
	//The initial account balance in satoshis representing the maximum amount that
	//can be spent.
	AccountBalance uint64 `protobuf:"varint,1,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	//*
	//The expiration date of the account as a timestamp. Set to zero to never
	//expire.
	ExpirationDate       int64    `protobuf:"varint,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetAccountBalance() uint64 {
	if m != nil {
		return m.AccountBalance
	}
	return 0
}

func (m *CreateAccountRequest) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

type Account struct {
	//*
	//The ID of the account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//*
	//The initial balance in satoshis that was set when the account was created.
	InitialBalance uint64 `protobuf:"varint,2,opt,name=initial_balance,json=initialBalance,proto3" json:"initial_balance,omitempty"`
	//*
	//The current balance in satoshis.
	CurrentBalance uint64 `protobuf:"varint,3,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
	//*
	//Timestamp of the last time the account was updated.
	LastUpdate int64 `protobuf:"varint,4,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	//*
	//Timestamp of the account's expiration date. Zero means it does not expire.
	ExpirationDate       int64    `protobuf:"varint,5,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetInitialBalance() uint64 {
	if m != nil {
		return m.InitialBalance
	}
	return 0
}

func (m *Account) GetCurrentBalance() uint64 {
	if m != nil {
		return m.CurrentBalance
	}
	return 0
}

func (m *Account) GetLastUpdate() int64 {
	if m != nil {
		return m.LastUpdate
	}
	return 0
}

func (m *Account) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

type CreateAccountResponse struct {
	//*
	//The created account.
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}

func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(m, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ListAccountsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccountsRequest) Reset()         { *m = ListAccountsRequest{} }
func (m *ListAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccountsRequest) ProtoMessage()    {}
func (*ListAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}

func (m *ListAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccountsRequest.Unmarshal(m, b)
}
func (m *ListAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ListAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsRequest.Merge(m, src)
}
func (m *ListAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccountsRequest.Size(m)
}
func (m *ListAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsRequest proto.InternalMessageInfo

type ListAccountsResponse struct {
	//*
	//All accounts in the account database.
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListAccountsResponse) Reset()         { *m = ListAccountsResponse{} }
func (m *ListAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccountsResponse) ProtoMessage()    {}
func (*ListAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{4}
}

func (m *ListAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccountsResponse.Unmarshal(m, b)
}
func (m *ListAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ListAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsResponse.Merge(m, src)
}
func (m *ListAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccountsResponse.Size(m)
}
func (m *ListAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsResponse proto.InternalMessageInfo

func (m *ListAccountsResponse) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type RemoveAccountRequest struct {
	//*
	//The hexadecimal ID of the account to remove.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAccountRequest) Reset()         { *m = RemoveAccountRequest{} }
func (m *RemoveAccountRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveAccountRequest) ProtoMessage()    {}
func (*RemoveAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{5}
}

func (m *RemoveAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAccountRequest.Unmarshal(m, b)
}
func (m *RemoveAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAccountRequest.Marshal(b, m, deterministic)
}
func (m *RemoveAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAccountRequest.Merge(m, src)
}
func (m *RemoveAccountRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveAccountRequest.Size(m)
}
func (m *RemoveAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAccountRequest proto.InternalMessageInfo

func (m *RemoveAccountRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAccountResponse) Reset()         { *m = RemoveAccountResponse{} }
func (m *RemoveAccountResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveAccountResponse) ProtoMessage()    {}
func (*RemoveAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{6}
}

func (m *RemoveAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAccountResponse.Unmarshal(m, b)
}
func (m *RemoveAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAccountResponse.Marshal(b, m, deterministic)
}
func (m *RemoveAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAccountResponse.Merge(m, src)
}
func (m *RemoveAccountResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveAccountResponse.Size(m)
}
func (m *RemoveAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "acctrpc.CreateAccountRequest")
	proto.RegisterType((*Account)(nil), "acctrpc.Account")
	proto.RegisterType((*CreateAccountResponse)(nil), "acctrpc.CreateAccountResponse")
	proto.RegisterType((*ListAccountsRequest)(nil), "acctrpc.ListAccountsRequest")
	proto.RegisterType((*ListAccountsResponse)(nil), "acctrpc.ListAccountsResponse")
	proto.RegisterType((*RemoveAccountRequest)(nil), "acctrpc.RemoveAccountRequest")
	proto.RegisterType((*RemoveAccountResponse)(nil), "acctrpc.RemoveAccountResponse")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x8e, 0xd3, 0x30,
	0x18, 0x55, 0xd2, 0x81, 0x0e, 0xdf, 0xcc, 0x04, 0xc6, 0x93, 0x6a, 0x4a, 0xd5, 0x42, 0xe5, 0x05,
	0x2d, 0x15, 0x24, 0xa2, 0xec, 0xd8, 0xd1, 0x76, 0xc9, 0x2a, 0x12, 0x1b, 0x36, 0x95, 0xeb, 0x58,
	0xa9, 0xa5, 0xd4, 0x0e, 0x89, 0x53, 0x21, 0x21, 0x36, 0x5c, 0x81, 0xb3, 0x70, 0x00, 0xce, 0xc0,
	0x15, 0x38, 0x08, 0xaa, 0xe3, 0x64, 0x92, 0x28, 0xdd, 0x3e, 0xbf, 0xbc, 0x3f, 0xc7, 0xe0, 0x10,
	0x4a, 0x65, 0x2e, 0x54, 0xe6, 0x25, 0xa9, 0x54, 0x12, 0xf5, 0x09, 0xa5, 0x2a, 0x4d, 0xe8, 0x68,
	0x1c, 0x49, 0x19, 0xc5, 0xcc, 0x27, 0x09, 0xf7, 0x89, 0x10, 0x52, 0x11, 0xc5, 0xa5, 0x30, 0x34,
	0xbc, 0x07, 0x77, 0x9d, 0x32, 0xa2, 0xd8, 0xc7, 0xe2, 0xf3, 0x80, 0x7d, 0xcd, 0x59, 0xa6, 0xd0,
	0x0c, 0x9e, 0x1a, 0xc1, 0xed, 0x8e, 0xc4, 0x44, 0x50, 0x36, 0xb4, 0xa6, 0xd6, 0xfc, 0x22, 0x28,
	0x7d, 0x56, 0x05, 0x7a, 0x22, 0xb2, 0x6f, 0x09, 0x4f, 0xb5, 0xea, 0x36, 0x24, 0x8a, 0x0d, 0xed,
	0xa9, 0x35, 0xef, 0x05, 0xce, 0x03, 0xbc, 0x21, 0x8a, 0xe1, 0xdf, 0x16, 0xf4, 0x8d, 0x09, 0x72,
	0xc0, 0xe6, 0xa1, 0x16, 0x7c, 0x12, 0xd8, 0x3c, 0x3c, 0x89, 0x70, 0xc1, 0x15, 0x27, 0x71, 0xe5,
	0x66, 0x17, 0x6e, 0x06, 0xae, 0xb9, 0xd1, 0x3c, 0x4d, 0x59, 0x2d, 0x56, 0xaf, 0x20, 0x1a, 0xb8,
	0x24, 0xbe, 0x84, 0xab, 0x98, 0x64, 0x6a, 0x9b, 0x27, 0x3a, 0xd2, 0x85, 0x8e, 0x04, 0x27, 0xe8,
	0xb3, 0x46, 0xba, 0x72, 0x3f, 0xea, 0xcc, 0xbd, 0x86, 0x41, 0x6b, 0xa1, 0x2c, 0x91, 0x22, 0x63,
	0x68, 0x01, 0x7d, 0xb3, 0x85, 0x6e, 0x72, 0xb5, 0x7c, 0xe6, 0x99, 0xcd, 0xbd, 0x92, 0x5a, 0x12,
	0xf0, 0x00, 0xee, 0x3e, 0xf1, 0x4c, 0x19, 0x3c, 0x33, 0x2b, 0xe3, 0x0d, 0xb8, 0x4d, 0xd8, 0x48,
	0xbf, 0x81, 0xcb, 0xf2, 0x3a, 0x87, 0xd6, 0xb4, 0xd7, 0xa9, 0x5d, 0x31, 0xf0, 0x2b, 0x70, 0x03,
	0x76, 0x90, 0xc7, 0xf6, 0x1d, 0xb6, 0x56, 0xc6, 0xf7, 0x30, 0x68, 0xf1, 0x0a, 0xbb, 0xe5, 0x1f,
	0x1b, 0x2e, 0xcb, 0x0c, 0x28, 0x82, 0x9b, 0x46, 0x5f, 0x34, 0xa9, 0xac, 0xbb, 0xfe, 0x94, 0xd1,
	0x8b, 0x73, 0xc7, 0x85, 0x38, 0xbe, 0xff, 0xf9, 0xf7, 0xdf, 0x2f, 0xfb, 0x16, 0x5f, 0xfb, 0xc7,
	0x77, 0x7e, 0x99, 0xf9, 0x83, 0xb5, 0x40, 0x04, 0xae, 0xeb, 0xe5, 0xd1, 0xb8, 0x12, 0xea, 0x98,
	0x6a, 0x34, 0x39, 0x73, 0x6a, 0x5c, 0x5c, 0xed, 0xe2, 0xa0, 0x86, 0x0b, 0xe2, 0x70, 0xd3, 0x68,
	0x5c, 0xeb, 0xd2, 0xb5, 0x58, 0xad, 0x4b, 0xe7, 0x50, 0xf8, 0xb9, 0x76, 0xb9, 0x5b, 0xdc, 0xd6,
	0x5d, 0xfc, 0xef, 0x3c, 0xfc, 0xb1, 0x7a, 0xfd, 0x65, 0x16, 0x71, 0xb5, 0xcf, 0x77, 0x1e, 0x95,
	0x07, 0x3f, 0xca, 0xa3, 0x88, 0xa5, 0xd2, 0x8f, 0x45, 0xf8, 0xb6, 0xe2, 0x19, 0xed, 0xdd, 0x63,
	0xfd, 0xf4, 0xde, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xcf, 0x36, 0xe1, 0xb3, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	//* lncli: `createaccount`
	//CreateAccount adds an entry to the account database. This entry represents
	//an amount of satoshis (account balance) that can be spent using off-chain
	//transactions (e.g. paying invoices).
	//
	//Macaroons can be created to be locked to an account. This makes sure that
	//the bearer of the macaroon can only spend at most that amount of satoshis
	//through the daemon that has issued the macaroon.
	//
	//Accounts only assert a maximum amount spendable. Having a certain account
	//balance does not guarantee that the node has the channel liquidity to
	//actually spend that amount.
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	//* lncli: `listaccounts`
	//ListAccounts returns all accounts that are currently stored in the account
	//database.
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	//* lncli: `removeaccount`
	//RemoveAccount removes the given account from the account database.
	RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...grpc.CallOption) (*RemoveAccountResponse, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/acctrpc.Accounts/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/acctrpc.Accounts/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...grpc.CallOption) (*RemoveAccountResponse, error) {
	out := new(RemoveAccountResponse)
	err := c.cc.Invoke(ctx, "/acctrpc.Accounts/RemoveAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	//* lncli: `createaccount`
	//CreateAccount adds an entry to the account database. This entry represents
	//an amount of satoshis (account balance) that can be spent using off-chain
	//transactions (e.g. paying invoices).
	//
	//Macaroons can be created to be locked to an account. This makes sure that
	//the bearer of the macaroon can only spend at most that amount of satoshis
	//through the daemon that has issued the macaroon.
	//
	//Accounts only assert a maximum amount spendable. Having a certain account
	//balance does not guarantee that the node has the channel liquidity to
	//actually spend that amount.
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	//* lncli: `listaccounts`
	//ListAccounts returns all accounts that are currently stored in the account
	//database.
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	//* lncli: `removeaccount`
	//RemoveAccount removes the given account from the account database.
	RemoveAccount(context.Context, *RemoveAccountRequest) (*RemoveAccountResponse, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acctrpc.Accounts/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acctrpc.Accounts/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_RemoveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).RemoveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acctrpc.Accounts/RemoveAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).RemoveAccount(ctx, req.(*RemoveAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "acctrpc.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Accounts_CreateAccount_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _Accounts_ListAccounts_Handler,
		},
		{
			MethodName: "RemoveAccount",
			Handler:    _Accounts_RemoveAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}

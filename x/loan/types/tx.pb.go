// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loan/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Loan    string `protobuf:"bytes,2,opt,name=loan,proto3" json:"loan,omitempty"`
}

func (m *MsgRequest) Reset()         { *m = MsgRequest{} }
func (m *MsgRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()    {}
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f8f93b4237d328, []int{0}
}
func (m *MsgRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequest.Merge(m, src)
}
func (m *MsgRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequest proto.InternalMessageInfo

func (m *MsgRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRequest) GetLoan() string {
	if m != nil {
		return m.Loan
	}
	return ""
}

type MsgRequestResponse struct {
}

func (m *MsgRequestResponse) Reset()         { *m = MsgRequestResponse{} }
func (m *MsgRequestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestResponse) ProtoMessage()    {}
func (*MsgRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f8f93b4237d328, []int{1}
}
func (m *MsgRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestResponse.Merge(m, src)
}
func (m *MsgRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestResponse proto.InternalMessageInfo

type MsgRequestLoan struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee        string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Collateral string `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Deadline   string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *MsgRequestLoan) Reset()         { *m = MsgRequestLoan{} }
func (m *MsgRequestLoan) String() string { return proto.CompactTextString(m) }
func (*MsgRequestLoan) ProtoMessage()    {}
func (*MsgRequestLoan) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f8f93b4237d328, []int{2}
}
func (m *MsgRequestLoan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestLoan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestLoan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestLoan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestLoan.Merge(m, src)
}
func (m *MsgRequestLoan) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestLoan) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestLoan.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestLoan proto.InternalMessageInfo

func (m *MsgRequestLoan) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRequestLoan) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MsgRequestLoan) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *MsgRequestLoan) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *MsgRequestLoan) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

type MsgRequestLoanResponse struct {
}

func (m *MsgRequestLoanResponse) Reset()         { *m = MsgRequestLoanResponse{} }
func (m *MsgRequestLoanResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestLoanResponse) ProtoMessage()    {}
func (*MsgRequestLoanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f8f93b4237d328, []int{3}
}
func (m *MsgRequestLoanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestLoanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestLoanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestLoanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestLoanResponse.Merge(m, src)
}
func (m *MsgRequestLoanResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestLoanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestLoanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestLoanResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequest)(nil), "loan.loan.MsgRequest")
	proto.RegisterType((*MsgRequestResponse)(nil), "loan.loan.MsgRequestResponse")
	proto.RegisterType((*MsgRequestLoan)(nil), "loan.loan.MsgRequestLoan")
	proto.RegisterType((*MsgRequestLoanResponse)(nil), "loan.loan.MsgRequestLoanResponse")
}

func init() { proto.RegisterFile("loan/tx.proto", fileDescriptor_b8f8f93b4237d328) }

var fileDescriptor_b8f8f93b4237d328 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xcd, 0x9a, 0xda, 0xda, 0x11, 0x45, 0x07, 0x2d, 0x6b, 0xc0, 0x45, 0x73, 0x12, 0x84, 0x14,
	0xf4, 0xe6, 0x45, 0xf0, 0xaa, 0xbd, 0xe4, 0xe8, 0x6d, 0x6d, 0xc7, 0x22, 0xc4, 0x6c, 0xcc, 0x6e,
	0xa1, 0xfe, 0x45, 0xc1, 0x9f, 0xf2, 0xd8, 0xa3, 0x47, 0x49, 0x7e, 0x44, 0x76, 0xed, 0x36, 0x15,
	0x6a, 0x2f, 0xcb, 0xbc, 0xf7, 0x66, 0x1e, 0xb3, 0xf3, 0x60, 0x2f, 0x53, 0x32, 0xef, 0x9b, 0x69,
	0x52, 0x94, 0xca, 0x28, 0xec, 0x5a, 0x98, 0xd8, 0x27, 0xbe, 0x01, 0x18, 0xe8, 0x71, 0x4a, 0x6f,
	0x13, 0xd2, 0x06, 0x39, 0x74, 0x86, 0x25, 0x49, 0xa3, 0x4a, 0xce, 0xce, 0xd8, 0x45, 0x37, 0xf5,
	0x10, 0x11, 0x5a, 0xb6, 0x9f, 0x6f, 0x39, 0xda, 0xd5, 0xf1, 0x11, 0x60, 0x33, 0x9b, 0x92, 0x2e,
	0x54, 0xae, 0x29, 0x9e, 0x31, 0xd8, 0x6f, 0xe8, 0x07, 0x25, 0xf3, 0x0d, 0xb6, 0x3d, 0x68, 0xcb,
	0x57, 0x35, 0xc9, 0xcd, 0xc2, 0x78, 0x81, 0xf0, 0x00, 0xc2, 0x67, 0x22, 0x1e, 0x3a, 0xd2, 0x96,
	0x28, 0x00, 0x86, 0x2a, 0xcb, 0xa4, 0xa1, 0x52, 0x66, 0xbc, 0xe5, 0x84, 0x15, 0x06, 0x23, 0xd8,
	0x19, 0x91, 0x1c, 0x65, 0x2f, 0x39, 0xf1, 0x6d, 0xa7, 0x2e, 0x71, 0xcc, 0xa1, 0xf7, 0x77, 0x23,
	0xbf, 0xec, 0xd5, 0x07, 0x83, 0x70, 0xa0, 0xc7, 0x78, 0x0b, 0x1d, 0x7f, 0x83, 0xe3, 0x64, 0x79,
	0x9d, 0xa4, 0x99, 0x8a, 0x4e, 0xd7, 0xd2, 0xde, 0x08, 0xef, 0x61, 0x77, 0xf5, 0xc7, 0x27, 0x6b,
	0xbb, 0xad, 0x14, 0x9d, 0xff, 0x2b, 0x79, 0xb3, 0xbb, 0xcb, 0xcf, 0x4a, 0xb0, 0x79, 0x25, 0xd8,
	0x77, 0x25, 0xd8, 0xac, 0x16, 0xc1, 0xbc, 0x16, 0xc1, 0x57, 0x2d, 0x82, 0xc7, 0x43, 0x17, 0xe4,
	0xb4, 0xff, 0x9b, 0xe7, 0x7b, 0x41, 0xfa, 0xa9, 0xed, 0x32, 0xbd, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0xd0, 0xe7, 0x92, 0xe4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Request(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgRequestResponse, error)
	RequestLoan(ctx context.Context, in *MsgRequestLoan, opts ...grpc.CallOption) (*MsgRequestLoanResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Request(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgRequestResponse, error) {
	out := new(MsgRequestResponse)
	err := c.cc.Invoke(ctx, "/loan.loan.Msg/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestLoan(ctx context.Context, in *MsgRequestLoan, opts ...grpc.CallOption) (*MsgRequestLoanResponse, error) {
	out := new(MsgRequestLoanResponse)
	err := c.cc.Invoke(ctx, "/loan.loan.Msg/RequestLoan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Request(context.Context, *MsgRequest) (*MsgRequestResponse, error)
	RequestLoan(context.Context, *MsgRequestLoan) (*MsgRequestLoanResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Request(ctx context.Context, req *MsgRequest) (*MsgRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedMsgServer) RequestLoan(ctx context.Context, req *MsgRequestLoan) (*MsgRequestLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLoan not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loan.loan.Msg/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Request(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loan.loan.Msg/RequestLoan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestLoan(ctx, req.(*MsgRequestLoan))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loan.loan.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Msg_Request_Handler,
		},
		{
			MethodName: "RequestLoan",
			Handler:    _Msg_RequestLoan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loan/tx.proto",
}

func (m *MsgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Loan) > 0 {
		i -= len(m.Loan)
		copy(dAtA[i:], m.Loan)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Loan)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRequestLoan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestLoan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestLoan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestLoanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestLoanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestLoanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Loan)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRequestLoan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestLoanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Loan", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Loan = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRequestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRequestLoan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestLoan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestLoan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRequestLoanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestLoanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestLoanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

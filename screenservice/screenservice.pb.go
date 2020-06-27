// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screenservice.proto

package screenservice

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

type ScreenRequest struct {
	Line1                string   `protobuf:"bytes,1,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2                string   `protobuf:"bytes,2,opt,name=line2,proto3" json:"line2,omitempty"`
	Line3                string   `protobuf:"bytes,3,opt,name=line3,proto3" json:"line3,omitempty"`
	Line4                string   `protobuf:"bytes,4,opt,name=line4,proto3" json:"line4,omitempty"`
	Line5                string   `protobuf:"bytes,5,opt,name=line5,proto3" json:"line5,omitempty"`
	ReturnScreen         string   `protobuf:"bytes,6,opt,name=returnScreen,proto3" json:"returnScreen,omitempty"`
	Length               int32    `protobuf:"varint,7,opt,name=length,proto3" json:"length,omitempty"`
	ShowCountdown        int32    `protobuf:"varint,8,opt,name=showCountdown,proto3" json:"showCountdown,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScreenRequest) Reset()         { *m = ScreenRequest{} }
func (m *ScreenRequest) String() string { return proto.CompactTextString(m) }
func (*ScreenRequest) ProtoMessage()    {}
func (*ScreenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_screenservice_86f100ab84a8b024, []int{0}
}
func (m *ScreenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreenRequest.Unmarshal(m, b)
}
func (m *ScreenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreenRequest.Marshal(b, m, deterministic)
}
func (dst *ScreenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreenRequest.Merge(dst, src)
}
func (m *ScreenRequest) XXX_Size() int {
	return xxx_messageInfo_ScreenRequest.Size(m)
}
func (m *ScreenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScreenRequest proto.InternalMessageInfo

func (m *ScreenRequest) GetLine1() string {
	if m != nil {
		return m.Line1
	}
	return ""
}

func (m *ScreenRequest) GetLine2() string {
	if m != nil {
		return m.Line2
	}
	return ""
}

func (m *ScreenRequest) GetLine3() string {
	if m != nil {
		return m.Line3
	}
	return ""
}

func (m *ScreenRequest) GetLine4() string {
	if m != nil {
		return m.Line4
	}
	return ""
}

func (m *ScreenRequest) GetLine5() string {
	if m != nil {
		return m.Line5
	}
	return ""
}

func (m *ScreenRequest) GetReturnScreen() string {
	if m != nil {
		return m.ReturnScreen
	}
	return ""
}

func (m *ScreenRequest) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ScreenRequest) GetShowCountdown() int32 {
	if m != nil {
		return m.ShowCountdown
	}
	return 0
}

type ScreenResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScreenResponse) Reset()         { *m = ScreenResponse{} }
func (m *ScreenResponse) String() string { return proto.CompactTextString(m) }
func (*ScreenResponse) ProtoMessage()    {}
func (*ScreenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_screenservice_86f100ab84a8b024, []int{1}
}
func (m *ScreenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreenResponse.Unmarshal(m, b)
}
func (m *ScreenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreenResponse.Marshal(b, m, deterministic)
}
func (dst *ScreenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreenResponse.Merge(dst, src)
}
func (m *ScreenResponse) XXX_Size() int {
	return xxx_messageInfo_ScreenResponse.Size(m)
}
func (m *ScreenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScreenResponse proto.InternalMessageInfo

func (m *ScreenResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ScreenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ScreenRequest)(nil), "screenservice.ScreenRequest")
	proto.RegisterType((*ScreenResponse)(nil), "screenservice.ScreenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScreenServerClient is the client API for ScreenServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScreenServerClient interface {
	SendScreen(ctx context.Context, in *ScreenRequest, opts ...grpc.CallOption) (*ScreenResponse, error)
}

type screenServerClient struct {
	cc *grpc.ClientConn
}

func NewScreenServerClient(cc *grpc.ClientConn) ScreenServerClient {
	return &screenServerClient{cc}
}

func (c *screenServerClient) SendScreen(ctx context.Context, in *ScreenRequest, opts ...grpc.CallOption) (*ScreenResponse, error) {
	out := new(ScreenResponse)
	err := c.cc.Invoke(ctx, "/screenservice.ScreenServer/SendScreen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScreenServerServer is the server API for ScreenServer service.
type ScreenServerServer interface {
	SendScreen(context.Context, *ScreenRequest) (*ScreenResponse, error)
}

func RegisterScreenServerServer(s *grpc.Server, srv ScreenServerServer) {
	s.RegisterService(&_ScreenServer_serviceDesc, srv)
}

func _ScreenServer_SendScreen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScreenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServerServer).SendScreen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/screenservice.ScreenServer/SendScreen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServerServer).SendScreen(ctx, req.(*ScreenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScreenServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "screenservice.ScreenServer",
	HandlerType: (*ScreenServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendScreen",
			Handler:    _ScreenServer_SendScreen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "screenservice.proto",
}

func init() { proto.RegisterFile("screenservice.proto", fileDescriptor_screenservice_86f100ab84a8b024) }

var fileDescriptor_screenservice_86f100ab84a8b024 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0x60, 0xa3, 0x26, 0xd5, 0xa1, 0xf1, 0xb0, 0x8a, 0x0c, 0xa2, 0x50, 0x82, 0x87, 0x9e,
	0x0a, 0x26, 0xed, 0x0b, 0xe8, 0xd1, 0x5b, 0x72, 0xf4, 0x54, 0xdb, 0xa1, 0x2d, 0xd4, 0xd9, 0xba,
	0xb3, 0x69, 0x1f, 0xd7, 0x57, 0x11, 0xb3, 0xbb, 0x2c, 0x7b, 0xf0, 0xf8, 0x7f, 0x03, 0xcb, 0xfc,
	0xb3, 0x70, 0x2b, 0x2b, 0x43, 0xc4, 0x42, 0xe6, 0xb8, 0x5b, 0xd1, 0xec, 0x60, 0xb4, 0xd5, 0xaa,
	0x4c, 0xb0, 0xfa, 0xc9, 0xa0, 0xec, 0x06, 0x69, 0xe9, 0xbb, 0x27, 0xb1, 0xea, 0x0e, 0xf2, 0xfd,
	0x8e, 0xe9, 0x05, 0xb3, 0x49, 0x36, 0xbd, 0x6e, 0x5d, 0x08, 0x5a, 0xe3, 0x79, 0xd4, 0x3a, 0x68,
	0x83, 0x17, 0x51, 0x9b, 0xa0, 0x73, 0xbc, 0x8c, 0x3a, 0x0f, 0xba, 0xc0, 0x3c, 0xea, 0x42, 0x55,
	0x30, 0x36, 0x64, 0x7b, 0xc3, 0x6e, 0x09, 0x2c, 0x86, 0x61, 0x62, 0xea, 0x1e, 0x8a, 0x3d, 0xf1,
	0xc6, 0x6e, 0x71, 0x34, 0xc9, 0xa6, 0x79, 0xeb, 0x93, 0x7a, 0x86, 0x52, 0xb6, 0xfa, 0xf4, 0xa6,
	0x7b, 0xb6, 0x6b, 0x7d, 0x62, 0xbc, 0x1a, 0xc6, 0x29, 0x56, 0xaf, 0x70, 0x13, 0x0a, 0xca, 0x41,
	0xb3, 0xd0, 0xdf, 0x7b, 0x62, 0x97, 0xb6, 0x17, 0x5f, 0xd1, 0x27, 0x85, 0x30, 0xfa, 0x22, 0x91,
	0xe5, 0x86, 0x7c, 0xcb, 0x10, 0xeb, 0x0f, 0x18, 0xbb, 0x37, 0x3a, 0x32, 0x47, 0x32, 0xea, 0x1d,
	0xa0, 0x23, 0x5e, 0xfb, 0xfd, 0x1e, 0x67, 0xe9, 0xa1, 0x93, 0x7b, 0x3e, 0x3c, 0xfd, 0x33, 0x75,
	0xcb, 0x54, 0x67, 0x9f, 0xc5, 0xf0, 0x31, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xc8,
	0xd4, 0xad, 0xaf, 0x01, 0x00, 0x00,
}
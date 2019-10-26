// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package impb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FetchInput struct {
	ApiKey               string   `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	Referer              string   `protobuf:"bytes,3,opt,name=referer,proto3" json:"referer,omitempty"`
	Roof                 string   `protobuf:"bytes,4,opt,name=roof,proto3" json:"roof,omitempty"`
	UserID               int64    `protobuf:"varint,5,opt,name=userID,proto3" json:"userID,omitempty"`
	SizeOp               string   `protobuf:"bytes,6,opt,name=sizeOp,proto3" json:"sizeOp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchInput) Reset()         { *m = FetchInput{} }
func (m *FetchInput) String() string { return proto.CompactTextString(m) }
func (*FetchInput) ProtoMessage()    {}
func (*FetchInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{0}
}

func (m *FetchInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchInput.Unmarshal(m, b)
}
func (m *FetchInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchInput.Marshal(b, m, deterministic)
}
func (m *FetchInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchInput.Merge(m, src)
}
func (m *FetchInput) XXX_Size() int {
	return xxx_messageInfo_FetchInput.Size(m)
}
func (m *FetchInput) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchInput.DiscardUnknown(m)
}

var xxx_messageInfo_FetchInput proto.InternalMessageInfo

func (m *FetchInput) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *FetchInput) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *FetchInput) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *FetchInput) GetRoof() string {
	if m != nil {
		return m.Roof
	}
	return ""
}

func (m *FetchInput) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *FetchInput) GetSizeOp() string {
	if m != nil {
		return m.SizeOp
	}
	return ""
}

type ImageInput struct {
	ApiKey               string   `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Image                []byte   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Roof                 string   `protobuf:"bytes,4,opt,name=roof,proto3" json:"roof,omitempty"`
	UserID               int64    `protobuf:"varint,5,opt,name=userID,proto3" json:"userID,omitempty"`
	SizeOp               string   `protobuf:"bytes,6,opt,name=sizeOp,proto3" json:"sizeOp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInput) Reset()         { *m = ImageInput{} }
func (m *ImageInput) String() string { return proto.CompactTextString(m) }
func (*ImageInput) ProtoMessage()    {}
func (*ImageInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{1}
}

func (m *ImageInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInput.Unmarshal(m, b)
}
func (m *ImageInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInput.Marshal(b, m, deterministic)
}
func (m *ImageInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInput.Merge(m, src)
}
func (m *ImageInput) XXX_Size() int {
	return xxx_messageInfo_ImageInput.Size(m)
}
func (m *ImageInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInput.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInput proto.InternalMessageInfo

func (m *ImageInput) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *ImageInput) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageInput) GetRoof() string {
	if m != nil {
		return m.Roof
	}
	return ""
}

func (m *ImageInput) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ImageInput) GetSizeOp() string {
	if m != nil {
		return m.SizeOp
	}
	return ""
}

type ImageOutput struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	ID                   uint64   `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageOutput) Reset()         { *m = ImageOutput{} }
func (m *ImageOutput) String() string { return proto.CompactTextString(m) }
func (*ImageOutput) ProtoMessage()    {}
func (*ImageOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{2}
}

func (m *ImageOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageOutput.Unmarshal(m, b)
}
func (m *ImageOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageOutput.Marshal(b, m, deterministic)
}
func (m *ImageOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageOutput.Merge(m, src)
}
func (m *ImageOutput) XXX_Size() int {
	return xxx_messageInfo_ImageOutput.Size(m)
}
func (m *ImageOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ImageOutput proto.InternalMessageInfo

func (m *ImageOutput) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ImageOutput) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *ImageOutput) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchInput)(nil), "impb.FetchInput")
	proto.RegisterType((*ImageInput)(nil), "impb.ImageInput")
	proto.RegisterType((*ImageOutput)(nil), "impb.ImageOutput")
}

func init() { proto.RegisterFile("image.proto", fileDescriptor_9624c68e2b547544) }

var fileDescriptor_9624c68e2b547544 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xbf, 0xf4, 0xdf, 0xa7, 0x77, 0x44, 0xc6, 0x20, 0x12, 0x5c, 0x0d, 0x5d, 0x75, 0x95,
	0x85, 0xbe, 0xc1, 0x58, 0x84, 0xe2, 0x62, 0x86, 0xce, 0x13, 0x74, 0xca, 0xad, 0x8d, 0xd8, 0x49,
	0x48, 0x53, 0x71, 0x7c, 0x0c, 0xf1, 0x81, 0x25, 0xb7, 0x1d, 0xc6, 0xc5, 0xe0, 0xca, 0xdd, 0x39,
	0x87, 0x4b, 0xf2, 0xcb, 0xc9, 0x85, 0x99, 0xea, 0xaa, 0x67, 0x94, 0xc6, 0x6a, 0xa7, 0x79, 0xa4,
	0x3a, 0xb3, 0x4d, 0xbf, 0x18, 0xc0, 0x23, 0xba, 0xba, 0x2d, 0x76, 0x66, 0x70, 0xfc, 0x06, 0x92,
	0xca, 0xa8, 0x27, 0xdc, 0x0b, 0xb6, 0x60, 0xd9, 0x79, 0x39, 0x39, 0x3e, 0x87, 0x70, 0xb0, 0x4a,
	0x04, 0x14, 0x7a, 0xc9, 0x05, 0xfc, 0xb7, 0xd8, 0xa0, 0x45, 0x2b, 0x42, 0x4a, 0x0f, 0x96, 0x73,
	0x88, 0xac, 0xd6, 0x8d, 0x88, 0x28, 0x26, 0xed, 0xcf, 0x1d, 0x7a, 0xb4, 0x45, 0x2e, 0xe2, 0x05,
	0xcb, 0xc2, 0x72, 0x72, 0x3e, 0xef, 0xd5, 0x07, 0xae, 0x8c, 0x48, 0xc6, 0xfb, 0x46, 0x97, 0x7e,
	0x32, 0x80, 0xc2, 0xc3, 0xfe, 0x8e, 0x75, 0x0d, 0x31, 0x3d, 0x89, 0xc0, 0x2e, 0xca, 0xd1, 0x78,
	0x80, 0x5d, 0xd5, 0xe1, 0xc4, 0x45, 0xfa, 0x4f, 0xa0, 0x1e, 0x60, 0x46, 0x4c, 0xab, 0xc1, 0x79,
	0x28, 0x0e, 0x91, 0xa9, 0x5c, 0x3b, 0x21, 0x91, 0x3e, 0xd1, 0xd3, 0x25, 0x04, 0x45, 0x4e, 0x28,
	0x51, 0x19, 0x14, 0xf9, 0xdd, 0x0b, 0x9c, 0xd1, 0x21, 0x9b, 0xb7, 0x9a, 0x4b, 0x88, 0xa9, 0x7b,
	0x3e, 0x97, 0xfe, 0x33, 0xe4, 0xf1, 0x23, 0x6e, 0xaf, 0xc6, 0xe4, 0xc7, 0x7d, 0xe9, 0x3f, 0x3f,
	0xbf, 0x71, 0xda, 0xe2, 0x61, 0xfe, 0xd8, 0xd0, 0xc9, 0xf9, 0x65, 0x06, 0xc2, 0x61, 0xdd, 0xca,
	0xa6, 0xdd, 0xbf, 0x4b, 0xf3, 0x5a, 0xb9, 0x46, 0xdb, 0x4e, 0xaa, 0xae, 0x77, 0x7a, 0x39, 0xd6,
	0xbb, 0xf6, 0xab, 0xb0, 0x66, 0xdb, 0x84, 0x76, 0xe2, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x6d,
	0xcc, 0x33, 0xbe, 0x22, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageSvcClient is the client API for ImageSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageSvcClient interface {
	// Fetch
	Fetch(ctx context.Context, in *FetchInput, opts ...grpc.CallOption) (*ImageOutput, error)
	// Store
	Store(ctx context.Context, in *ImageInput, opts ...grpc.CallOption) (*ImageOutput, error)
}

type imageSvcClient struct {
	cc *grpc.ClientConn
}

func NewImageSvcClient(cc *grpc.ClientConn) ImageSvcClient {
	return &imageSvcClient{cc}
}

func (c *imageSvcClient) Fetch(ctx context.Context, in *FetchInput, opts ...grpc.CallOption) (*ImageOutput, error) {
	out := new(ImageOutput)
	err := c.cc.Invoke(ctx, "/impb.ImageSvc/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageSvcClient) Store(ctx context.Context, in *ImageInput, opts ...grpc.CallOption) (*ImageOutput, error) {
	out := new(ImageOutput)
	err := c.cc.Invoke(ctx, "/impb.ImageSvc/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageSvcServer is the server API for ImageSvc service.
type ImageSvcServer interface {
	// Fetch
	Fetch(context.Context, *FetchInput) (*ImageOutput, error)
	// Store
	Store(context.Context, *ImageInput) (*ImageOutput, error)
}

// UnimplementedImageSvcServer can be embedded to have forward compatible implementations.
type UnimplementedImageSvcServer struct {
}

func (*UnimplementedImageSvcServer) Fetch(ctx context.Context, req *FetchInput) (*ImageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedImageSvcServer) Store(ctx context.Context, req *ImageInput) (*ImageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}

func RegisterImageSvcServer(s *grpc.Server, srv ImageSvcServer) {
	s.RegisterService(&_ImageSvc_serviceDesc, srv)
}

func _ImageSvc_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSvcServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impb.ImageSvc/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSvcServer).Fetch(ctx, req.(*FetchInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageSvc_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSvcServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impb.ImageSvc/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSvcServer).Store(ctx, req.(*ImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "impb.ImageSvc",
	HandlerType: (*ImageSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _ImageSvc_Fetch_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _ImageSvc_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}

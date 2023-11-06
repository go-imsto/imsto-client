// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package impb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImageSvcClient is the client API for ImageSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageSvcClient interface {
	// Fetch
	Fetch(ctx context.Context, in *FetchInput, opts ...grpc.CallOption) (*ImageOutput, error)
	// Store
	Store(ctx context.Context, in *ImageInput, opts ...grpc.CallOption) (*ImageOutput, error)
}

type imageSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewImageSvcClient(cc grpc.ClientConnInterface) ImageSvcClient {
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
// All implementations must embed UnimplementedImageSvcServer
// for forward compatibility
type ImageSvcServer interface {
	// Fetch
	Fetch(context.Context, *FetchInput) (*ImageOutput, error)
	// Store
	Store(context.Context, *ImageInput) (*ImageOutput, error)
	mustEmbedUnimplementedImageSvcServer()
}

// UnimplementedImageSvcServer must be embedded to have forward compatible implementations.
type UnimplementedImageSvcServer struct {
}

func (*UnimplementedImageSvcServer) Fetch(context.Context, *FetchInput) (*ImageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedImageSvcServer) Store(context.Context, *ImageInput) (*ImageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedImageSvcServer) mustEmbedUnimplementedImageSvcServer() {}

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
	Metadata: "imsto.proto",
}

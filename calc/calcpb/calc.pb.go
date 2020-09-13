// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: calc/calcpb/calc.proto

package calcpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddendLeft  int32 `protobuf:"varint,1,opt,name=addendLeft,proto3" json:"addendLeft,omitempty"`
	AddendRight int32 `protobuf:"varint,2,opt,name=addendRight,proto3" json:"addendRight,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetAddendLeft() int32 {
	if x != nil {
		return x.AddendLeft
	}
	return 0
}

func (x *CalcRequest) GetAddendRight() int32 {
	if x != nil {
		return x.AddendRight
	}
	return 0
}

type CalcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalcResponse) Reset() {
	*x = CalcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcResponse) ProtoMessage() {}

func (x *CalcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcResponse.ProtoReflect.Descriptor instead.
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return file_calc_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_calc_calcpb_calc_proto protoreflect.FileDescriptor

var file_calc_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x4f,
	0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x52, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x26, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x3e, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x11,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_calcpb_calc_proto_rawDescOnce sync.Once
	file_calc_calcpb_calc_proto_rawDescData = file_calc_calcpb_calc_proto_rawDesc
)

func file_calc_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calc_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calc_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_calcpb_calc_proto_rawDescData)
	})
	return file_calc_calcpb_calc_proto_rawDescData
}

var file_calc_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calc_calcpb_calc_proto_goTypes = []interface{}{
	(*CalcRequest)(nil),  // 0: calc.CalcRequest
	(*CalcResponse)(nil), // 1: calc.CalcResponse
}
var file_calc_calcpb_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalcService.Calc:input_type -> calc.CalcRequest
	1, // 1: calc.CalcService.Calc:output_type -> calc.CalcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_calcpb_calc_proto_init() }
func file_calc_calcpb_calc_proto_init() {
	if File_calc_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calc_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calc_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calc_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calc_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calc_calcpb_calc_proto = out.File
	file_calc_calcpb_calc_proto_rawDesc = nil
	file_calc_calcpb_calc_proto_goTypes = nil
	file_calc_calcpb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	// define unary API
	Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	// define unary API
	Calc(context.Context, *CalcRequest) (*CalcResponse, error)
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Calc(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Calc(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcService_Calc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc/calcpb/calc.proto",
}

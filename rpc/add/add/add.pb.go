// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: add.proto

package add

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

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book  string `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Price int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *AddReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AddResp) Reset() {
	*x = AddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResp) ProtoMessage() {}

func (x *AddResp) ProtoReflect() protoreflect.Message {
	mi := &file_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResp.ProtoReflect.Descriptor instead.
func (*AddResp) Descriptor() ([]byte, []int) {
	return file_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_add_proto protoreflect.FileDescriptor

var file_add_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x64, 0x64,
	0x22, 0x32, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32,
	0x4e, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12,
	0x0b, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x61, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x61,
	0x64, 0x64, 0x2e, 0x61, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x61, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x0c, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x61, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_add_proto_rawDescOnce sync.Once
	file_add_proto_rawDescData = file_add_proto_rawDesc
)

func file_add_proto_rawDescGZIP() []byte {
	file_add_proto_rawDescOnce.Do(func() {
		file_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_add_proto_rawDescData)
	})
	return file_add_proto_rawDescData
}

var file_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),  // 0: add.addReq
	(*AddResp)(nil), // 1: add.addResp
}
var file_add_proto_depIdxs = []int32{
	0, // 0: add.adder.add:input_type -> add.addReq
	0, // 1: add.adder.update:input_type -> add.addReq
	1, // 2: add.adder.add:output_type -> add.addResp
	1, // 3: add.adder.update:output_type -> add.addResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_add_proto_init() }
func file_add_proto_init() {
	if File_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResp); i {
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
			RawDescriptor: file_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_add_proto_goTypes,
		DependencyIndexes: file_add_proto_depIdxs,
		MessageInfos:      file_add_proto_msgTypes,
	}.Build()
	File_add_proto = out.File
	file_add_proto_rawDesc = nil
	file_add_proto_goTypes = nil
	file_add_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdderClient interface {
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error)
	Update(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	out := new(AddResp)
	err := c.cc.Invoke(ctx, "/add.adder/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adderClient) Update(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	out := new(AddResp)
	err := c.cc.Invoke(ctx, "/add.adder/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServer is the server API for Adder service.
type AdderServer interface {
	Add(context.Context, *AddReq) (*AddResp, error)
	Update(context.Context, *AddReq) (*AddResp, error)
}

// UnimplementedAdderServer can be embedded to have forward compatible implementations.
type UnimplementedAdderServer struct {
}

func (*UnimplementedAdderServer) Add(context.Context, *AddReq) (*AddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAdderServer) Update(context.Context, *AddReq) (*AddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterAdderServer(s *grpc.Server, srv AdderServer) {
	s.RegisterService(&_Adder_serviceDesc, srv)
}

func _Adder_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.adder/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adder_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.adder/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Update(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "add.adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _Adder_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Adder_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "add.proto",
}

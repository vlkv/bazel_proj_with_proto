// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/myproto_api.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_myproto_api_proto protoreflect.FileDescriptor

var file_proto_myproto_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x76, 0x6c, 0x6b, 0x76, 0x2e,
	0x62, 0x61, 0x7a, 0x65, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x6e, 0x0a, 0x0a, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x12, 0x60,
	0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x12, 0x29, 0x2e, 0x76, 0x6c, 0x6b, 0x76, 0x2e,
	0x62, 0x61, 0x7a, 0x65, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x2e, 0x76, 0x6c, 0x6b, 0x76, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x00,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x6c, 0x6b, 0x76, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_myproto_api_proto_goTypes = []interface{}{
	(*MyProto)(nil), // 0: vlkv.bazel_proj_with_proto.proto.MyProto
}
var file_proto_myproto_api_proto_depIdxs = []int32{
	0, // 0: vlkv.bazel_proj_with_proto.proto.MyProtoApi.FooBar:input_type -> vlkv.bazel_proj_with_proto.proto.MyProto
	0, // 1: vlkv.bazel_proj_with_proto.proto.MyProtoApi.FooBar:output_type -> vlkv.bazel_proj_with_proto.proto.MyProto
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_myproto_api_proto_init() }
func file_proto_myproto_api_proto_init() {
	if File_proto_myproto_api_proto != nil {
		return
	}
	file_proto_myproto_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_myproto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_myproto_api_proto_goTypes,
		DependencyIndexes: file_proto_myproto_api_proto_depIdxs,
	}.Build()
	File_proto_myproto_api_proto = out.File
	file_proto_myproto_api_proto_rawDesc = nil
	file_proto_myproto_api_proto_goTypes = nil
	file_proto_myproto_api_proto_depIdxs = nil
}
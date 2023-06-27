// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: example/service/v1/service.proto

package exampleservicev1

import (
	_ "github.com/leetm4n/rpc-error-proto-mapper-go/api/proto/rpc/errormapper/v1"
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

type ServiceError int32

const (
	// The default value is reserved for the proto compiler.
	ServiceError_SERVICE_ERROR_UNSPECIFIED ServiceError = 0
	// The entity was not found.
	ServiceError_SERVICE_ERROR_ENTITY_NOT_FOUND ServiceError = 1
	// The user is not eligible for the action.
	ServiceError_SERVICE_ERROR_USER_NOT_ELIGIBLE_FOR_ACTION ServiceError = 2
)

// Enum value maps for ServiceError.
var (
	ServiceError_name = map[int32]string{
		0: "SERVICE_ERROR_UNSPECIFIED",
		1: "SERVICE_ERROR_ENTITY_NOT_FOUND",
		2: "SERVICE_ERROR_USER_NOT_ELIGIBLE_FOR_ACTION",
	}
	ServiceError_value = map[string]int32{
		"SERVICE_ERROR_UNSPECIFIED":                  0,
		"SERVICE_ERROR_ENTITY_NOT_FOUND":             1,
		"SERVICE_ERROR_USER_NOT_ELIGIBLE_FOR_ACTION": 2,
	}
)

func (x ServiceError) Enum() *ServiceError {
	p := new(ServiceError)
	*p = x
	return p
}

func (x ServiceError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceError) Descriptor() protoreflect.EnumDescriptor {
	return file_example_service_v1_service_proto_enumTypes[0].Descriptor()
}

func (ServiceError) Type() protoreflect.EnumType {
	return &file_example_service_v1_service_proto_enumTypes[0]
}

func (x ServiceError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceError.Descriptor instead.
func (ServiceError) EnumDescriptor() ([]byte, []int) {
	return file_example_service_v1_service_proto_rawDescGZIP(), []int{0}
}

var File_example_service_v1_service_proto protoreflect.FileDescriptor

var file_example_service_v1_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xdb, 0x01, 0x0a, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x19,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x1a, 0xda, 0xb3, 0x02,
	0x16, 0xa0, 0xb7, 0x02, 0x02, 0xaa, 0xb7, 0x02, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x1e, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x1a, 0x1a, 0xda, 0xb3, 0x02,
	0x16, 0xa0, 0xb7, 0x02, 0x05, 0xaa, 0xb7, 0x02, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x2a, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x1a, 0x1a, 0xda, 0xb3, 0x02, 0x16, 0xa0, 0xb7, 0x02,
	0x07, 0xaa, 0xb7, 0x02, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x1a, 0x04, 0xc8, 0xb3, 0x02, 0x01, 0x42, 0xec, 0x01, 0x0a, 0x15, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x65, 0x65, 0x74, 0x6d, 0x34, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2d, 0x67,
	0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x45, 0x58, 0xaa, 0x02, 0x11, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_service_v1_service_proto_rawDescOnce sync.Once
	file_example_service_v1_service_proto_rawDescData = file_example_service_v1_service_proto_rawDesc
)

func file_example_service_v1_service_proto_rawDescGZIP() []byte {
	file_example_service_v1_service_proto_rawDescOnce.Do(func() {
		file_example_service_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_service_v1_service_proto_rawDescData)
	})
	return file_example_service_v1_service_proto_rawDescData
}

var file_example_service_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_service_v1_service_proto_goTypes = []interface{}{
	(ServiceError)(0), // 0: example.errors.v1.ServiceError
}
var file_example_service_v1_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_service_v1_service_proto_init() }
func file_example_service_v1_service_proto_init() {
	if File_example_service_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_service_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_service_v1_service_proto_goTypes,
		DependencyIndexes: file_example_service_v1_service_proto_depIdxs,
		EnumInfos:         file_example_service_v1_service_proto_enumTypes,
	}.Build()
	File_example_service_v1_service_proto = out.File
	file_example_service_v1_service_proto_rawDesc = nil
	file_example_service_v1_service_proto_goTypes = nil
	file_example_service_v1_service_proto_depIdxs = nil
}

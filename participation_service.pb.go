// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: participation_service.proto

package songcontestraterprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_participation_service_proto protoreflect.FileDescriptor

var file_participation_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x1a,
	0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xcd, 0x02, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2c, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6e, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_participation_service_proto_goTypes = []any{
	(*emptypb.Empty)(nil),              // 0: google.protobuf.Empty
	(*CreateParticipationRequest)(nil), // 1: songcontestrater.CreateParticipationRequest
	(*DeleteParticipationRequest)(nil), // 2: songcontestrater.DeleteParticipationRequest
	(*ListParticipationsResponse)(nil), // 3: songcontestrater.ListParticipationsResponse
	(*ParticipationResponse)(nil),      // 4: songcontestrater.ParticipationResponse
}
var file_participation_service_proto_depIdxs = []int32{
	0, // 0: songcontestrater.Participation.ListParticipations:input_type -> google.protobuf.Empty
	1, // 1: songcontestrater.Participation.CreateParticipation:input_type -> songcontestrater.CreateParticipationRequest
	2, // 2: songcontestrater.Participation.DeleteParticipation:input_type -> songcontestrater.DeleteParticipationRequest
	3, // 3: songcontestrater.Participation.ListParticipations:output_type -> songcontestrater.ListParticipationsResponse
	4, // 4: songcontestrater.Participation.CreateParticipation:output_type -> songcontestrater.ParticipationResponse
	4, // 5: songcontestrater.Participation.DeleteParticipation:output_type -> songcontestrater.ParticipationResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_participation_service_proto_init() }
func file_participation_service_proto_init() {
	if File_participation_service_proto != nil {
		return
	}
	file_participation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_participation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_participation_service_proto_goTypes,
		DependencyIndexes: file_participation_service_proto_depIdxs,
	}.Build()
	File_participation_service_proto = out.File
	file_participation_service_proto_rawDesc = nil
	file_participation_service_proto_goTypes = nil
	file_participation_service_proto_depIdxs = nil
}

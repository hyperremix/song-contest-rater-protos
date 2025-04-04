// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: stat.proto

package songcontestraterprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CriticType int32

const (
	CriticType_CRITIC_TYPE_UNSPECIFIED       CriticType = 0
	CriticType_CRITIC_TYPE_HARSH             CriticType = 1
	CriticType_CRITIC_TYPE_SLIGHTLY_CRITICAL CriticType = 2
	CriticType_CRITIC_TYPE_BALANCED          CriticType = 3
	CriticType_CRITIC_TYPE_EASY_TO_PLEASE    CriticType = 4
	CriticType_CRITIC_TYPE_GENEROUS          CriticType = 5
)

// Enum value maps for CriticType.
var (
	CriticType_name = map[int32]string{
		0: "CRITIC_TYPE_UNSPECIFIED",
		1: "CRITIC_TYPE_HARSH",
		2: "CRITIC_TYPE_SLIGHTLY_CRITICAL",
		3: "CRITIC_TYPE_BALANCED",
		4: "CRITIC_TYPE_EASY_TO_PLEASE",
		5: "CRITIC_TYPE_GENEROUS",
	}
	CriticType_value = map[string]int32{
		"CRITIC_TYPE_UNSPECIFIED":       0,
		"CRITIC_TYPE_HARSH":             1,
		"CRITIC_TYPE_SLIGHTLY_CRITICAL": 2,
		"CRITIC_TYPE_BALANCED":          3,
		"CRITIC_TYPE_EASY_TO_PLEASE":    4,
		"CRITIC_TYPE_GENEROUS":          5,
	}
)

func (x CriticType) Enum() *CriticType {
	p := new(CriticType)
	*p = x
	return p
}

func (x CriticType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CriticType) Descriptor() protoreflect.EnumDescriptor {
	return file_stat_proto_enumTypes[0].Descriptor()
}

func (CriticType) Type() protoreflect.EnumType {
	return &file_stat_proto_enumTypes[0]
}

func (x CriticType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CriticType.Descriptor instead.
func (CriticType) EnumDescriptor() ([]byte, []int) {
	return file_stat_proto_rawDescGZIP(), []int{0}
}

type ListUserStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*UserStatsResponse `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *ListUserStatsResponse) Reset() {
	*x = ListUserStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserStatsResponse) ProtoMessage() {}

func (x *ListUserStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserStatsResponse.ProtoReflect.Descriptor instead.
func (*ListUserStatsResponse) Descriptor() ([]byte, []int) {
	return file_stat_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserStatsResponse) GetStats() []*UserStatsResponse {
	if x != nil {
		return x.Stats
	}
	return nil
}

type UserStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRatingAvg float64                `protobuf:"fixed64,1,opt,name=user_rating_avg,json=userRatingAvg,proto3" json:"user_rating_avg,omitempty"`
	TotalRatings  int32                  `protobuf:"varint,2,opt,name=total_ratings,json=totalRatings,proto3" json:"total_ratings,omitempty"`
	RatingBias    float64                `protobuf:"fixed64,3,opt,name=rating_bias,json=ratingBias,proto3" json:"rating_bias,omitempty"`
	CriticType    CriticType             `protobuf:"varint,4,opt,name=critic_type,json=criticType,proto3,enum=songcontestrater.CriticType" json:"critic_type,omitempty"`
	User          *UserResponse          `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserStatsResponse) Reset() {
	*x = UserStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatsResponse) ProtoMessage() {}

func (x *UserStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatsResponse.ProtoReflect.Descriptor instead.
func (*UserStatsResponse) Descriptor() ([]byte, []int) {
	return file_stat_proto_rawDescGZIP(), []int{1}
}

func (x *UserStatsResponse) GetUserRatingAvg() float64 {
	if x != nil {
		return x.UserRatingAvg
	}
	return 0
}

func (x *UserStatsResponse) GetTotalRatings() int32 {
	if x != nil {
		return x.TotalRatings
	}
	return 0
}

func (x *UserStatsResponse) GetRatingBias() float64 {
	if x != nil {
		return x.RatingBias
	}
	return 0
}

func (x *UserStatsResponse) GetCriticType() CriticType {
	if x != nil {
		return x.CriticType
	}
	return CriticType_CRITIC_TYPE_UNSPECIFIED
}

func (x *UserStatsResponse) GetUser() *UserResponse {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserStatsResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserStatsResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GlobalStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalRatingAvg float64                `protobuf:"fixed64,1,opt,name=global_rating_avg,json=globalRatingAvg,proto3" json:"global_rating_avg,omitempty"`
	TotalRatings    int32                  `protobuf:"varint,2,opt,name=total_ratings,json=totalRatings,proto3" json:"total_ratings,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GlobalStatsResponse) Reset() {
	*x = GlobalStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalStatsResponse) ProtoMessage() {}

func (x *GlobalStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalStatsResponse.ProtoReflect.Descriptor instead.
func (*GlobalStatsResponse) Descriptor() ([]byte, []int) {
	return file_stat_proto_rawDescGZIP(), []int{2}
}

func (x *GlobalStatsResponse) GetGlobalRatingAvg() float64 {
	if x != nil {
		return x.GlobalRatingAvg
	}
	return 0
}

func (x *GlobalStatsResponse) GetTotalRatings() int32 {
	if x != nil {
		return x.TotalRatings
	}
	return 0
}

func (x *GlobalStatsResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GlobalStatsResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_stat_proto protoreflect.FileDescriptor

var file_stat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6f,
	0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22,
	0xea, 0x02, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x76, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x67, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x69, 0x61,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x42,
	0x69, 0x61, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdc, 0x01, 0x0a,
	0x13, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x76, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x67,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0xb7, 0x01, 0x0a, 0x0a,
	0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x52,
	0x49, 0x54, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x49, 0x54, 0x49,
	0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x52, 0x53, 0x48, 0x10, 0x01, 0x12, 0x21,
	0x0a, 0x1d, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4c,
	0x49, 0x47, 0x48, 0x54, 0x4c, 0x59, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x43,
	0x52, 0x49, 0x54, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x41, 0x53, 0x59, 0x5f,
	0x54, 0x4f, 0x5f, 0x50, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x52, 0x49, 0x54, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x4f, 0x55, 0x53, 0x10, 0x05, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stat_proto_rawDescOnce sync.Once
	file_stat_proto_rawDescData = file_stat_proto_rawDesc
)

func file_stat_proto_rawDescGZIP() []byte {
	file_stat_proto_rawDescOnce.Do(func() {
		file_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_stat_proto_rawDescData)
	})
	return file_stat_proto_rawDescData
}

var file_stat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stat_proto_goTypes = []any{
	(CriticType)(0),               // 0: songcontestrater.CriticType
	(*ListUserStatsResponse)(nil), // 1: songcontestrater.ListUserStatsResponse
	(*UserStatsResponse)(nil),     // 2: songcontestrater.UserStatsResponse
	(*GlobalStatsResponse)(nil),   // 3: songcontestrater.GlobalStatsResponse
	(*UserResponse)(nil),          // 4: songcontestrater.UserResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_stat_proto_depIdxs = []int32{
	2, // 0: songcontestrater.ListUserStatsResponse.stats:type_name -> songcontestrater.UserStatsResponse
	0, // 1: songcontestrater.UserStatsResponse.critic_type:type_name -> songcontestrater.CriticType
	4, // 2: songcontestrater.UserStatsResponse.user:type_name -> songcontestrater.UserResponse
	5, // 3: songcontestrater.UserStatsResponse.created_at:type_name -> google.protobuf.Timestamp
	5, // 4: songcontestrater.UserStatsResponse.updated_at:type_name -> google.protobuf.Timestamp
	5, // 5: songcontestrater.GlobalStatsResponse.created_at:type_name -> google.protobuf.Timestamp
	5, // 6: songcontestrater.GlobalStatsResponse.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_stat_proto_init() }
func file_stat_proto_init() {
	if File_stat_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserStatsResponse); i {
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
		file_stat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserStatsResponse); i {
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
		file_stat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GlobalStatsResponse); i {
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
			RawDescriptor: file_stat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stat_proto_goTypes,
		DependencyIndexes: file_stat_proto_depIdxs,
		EnumInfos:         file_stat_proto_enumTypes,
		MessageInfos:      file_stat_proto_msgTypes,
	}.Build()
	File_stat_proto = out.File
	file_stat_proto_rawDesc = nil
	file_stat_proto_goTypes = nil
	file_stat_proto_depIdxs = nil
}

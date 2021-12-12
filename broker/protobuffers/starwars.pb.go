// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protobuffers/starwars.proto

package __

import (
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

type FulcrumWriteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FulcrumWriteMessage) Reset() {
	*x = FulcrumWriteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_starwars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FulcrumWriteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FulcrumWriteMessage) ProtoMessage() {}

func (x *FulcrumWriteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_starwars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FulcrumWriteMessage.ProtoReflect.Descriptor instead.
func (*FulcrumWriteMessage) Descriptor() ([]byte, []int) {
	return file_protobuffers_starwars_proto_rawDescGZIP(), []int{0}
}

type InformantMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanetName string `protobuf:"bytes,1,opt,name=PlanetName,proto3" json:"PlanetName,omitempty"`
	CityName   string `protobuf:"bytes,2,opt,name=CityName,proto3" json:"CityName,omitempty"`
	NewValue   string `protobuf:"bytes,3,opt,name=newValue,proto3" json:"newValue,omitempty"`
}

func (x *InformantMessage) Reset() {
	*x = InformantMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_starwars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformantMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformantMessage) ProtoMessage() {}

func (x *InformantMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_starwars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformantMessage.ProtoReflect.Descriptor instead.
func (*InformantMessage) Descriptor() ([]byte, []int) {
	return file_protobuffers_starwars_proto_rawDescGZIP(), []int{1}
}

func (x *InformantMessage) GetPlanetName() string {
	if x != nil {
		return x.PlanetName
	}
	return ""
}

func (x *InformantMessage) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *InformantMessage) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

type BrokerWriteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirm bool  `protobuf:"varint,1,opt,name=confirm,proto3" json:"confirm,omitempty"`
	Replica int32 `protobuf:"varint,2,opt,name=replica,proto3" json:"replica,omitempty"`
}

func (x *BrokerWriteMessage) Reset() {
	*x = BrokerWriteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_starwars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerWriteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerWriteMessage) ProtoMessage() {}

func (x *BrokerWriteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_starwars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerWriteMessage.ProtoReflect.Descriptor instead.
func (*BrokerWriteMessage) Descriptor() ([]byte, []int) {
	return file_protobuffers_starwars_proto_rawDescGZIP(), []int{2}
}

func (x *BrokerWriteMessage) GetConfirm() bool {
	if x != nil {
		return x.Confirm
	}
	return false
}

func (x *BrokerWriteMessage) GetReplica() int32 {
	if x != nil {
		return x.Replica
	}
	return 0
}

var File_protobuffers_starwars_proto protoreflect.FileDescriptor

var file_protobuffers_starwars_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x22, 0x15, 0x0a, 0x13,
	0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x6a, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x69, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x48, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x32, 0xde, 0x02, 0x0a, 0x0d, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74,
	0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72,
	0x73, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72,
	0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61,
	0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0xe3, 0x02, 0x0a, 0x0e, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x53, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72,
	0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61,
	0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72,
	0x75, 0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuffers_starwars_proto_rawDescOnce sync.Once
	file_protobuffers_starwars_proto_rawDescData = file_protobuffers_starwars_proto_rawDesc
)

func file_protobuffers_starwars_proto_rawDescGZIP() []byte {
	file_protobuffers_starwars_proto_rawDescOnce.Do(func() {
		file_protobuffers_starwars_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuffers_starwars_proto_rawDescData)
	})
	return file_protobuffers_starwars_proto_rawDescData
}

var file_protobuffers_starwars_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuffers_starwars_proto_goTypes = []interface{}{
	(*FulcrumWriteMessage)(nil), // 0: protoStarWars.FulcrumWriteMessage
	(*InformantMessage)(nil),    // 1: protoStarWars.InformantMessage
	(*BrokerWriteMessage)(nil),  // 2: protoStarWars.BrokerWriteMessage
}
var file_protobuffers_starwars_proto_depIdxs = []int32{
	1, // 0: protoStarWars.BrokerService.AddCity:input_type -> protoStarWars.InformantMessage
	1, // 1: protoStarWars.BrokerService.UpdateName:input_type -> protoStarWars.InformantMessage
	1, // 2: protoStarWars.BrokerService.UpdateNumber:input_type -> protoStarWars.InformantMessage
	1, // 3: protoStarWars.BrokerService.DeleteCity:input_type -> protoStarWars.InformantMessage
	1, // 4: protoStarWars.FulcrumService.AddCity:input_type -> protoStarWars.InformantMessage
	1, // 5: protoStarWars.FulcrumService.UpdateName:input_type -> protoStarWars.InformantMessage
	1, // 6: protoStarWars.FulcrumService.UpdateNumber:input_type -> protoStarWars.InformantMessage
	1, // 7: protoStarWars.FulcrumService.DeleteCity:input_type -> protoStarWars.InformantMessage
	2, // 8: protoStarWars.BrokerService.AddCity:output_type -> protoStarWars.BrokerWriteMessage
	2, // 9: protoStarWars.BrokerService.UpdateName:output_type -> protoStarWars.BrokerWriteMessage
	2, // 10: protoStarWars.BrokerService.UpdateNumber:output_type -> protoStarWars.BrokerWriteMessage
	2, // 11: protoStarWars.BrokerService.DeleteCity:output_type -> protoStarWars.BrokerWriteMessage
	0, // 12: protoStarWars.FulcrumService.AddCity:output_type -> protoStarWars.FulcrumWriteMessage
	0, // 13: protoStarWars.FulcrumService.UpdateName:output_type -> protoStarWars.FulcrumWriteMessage
	0, // 14: protoStarWars.FulcrumService.UpdateNumber:output_type -> protoStarWars.FulcrumWriteMessage
	0, // 15: protoStarWars.FulcrumService.DeleteCity:output_type -> protoStarWars.FulcrumWriteMessage
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuffers_starwars_proto_init() }
func file_protobuffers_starwars_proto_init() {
	if File_protobuffers_starwars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuffers_starwars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FulcrumWriteMessage); i {
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
		file_protobuffers_starwars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformantMessage); i {
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
		file_protobuffers_starwars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerWriteMessage); i {
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
			RawDescriptor: file_protobuffers_starwars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protobuffers_starwars_proto_goTypes,
		DependencyIndexes: file_protobuffers_starwars_proto_depIdxs,
		MessageInfos:      file_protobuffers_starwars_proto_msgTypes,
	}.Build()
	File_protobuffers_starwars_proto = out.File
	file_protobuffers_starwars_proto_rawDesc = nil
	file_protobuffers_starwars_proto_goTypes = nil
	file_protobuffers_starwars_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: tanks.proto

package tankspb

import (
	channeldpb "github.com/channeldorg/channeld/pkg/channeldpb"
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

type TankState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Marks that the state should be removed from tankStates map
	Removed bool  `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
	Health  int32 `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
	IsAI    bool  `protobuf:"varint,3,opt,name=isAI,proto3" json:"isAI,omitempty"`
}

func (x *TankState) Reset() {
	*x = TankState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tanks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TankState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TankState) ProtoMessage() {}

func (x *TankState) ProtoReflect() protoreflect.Message {
	mi := &file_tanks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TankState.ProtoReflect.Descriptor instead.
func (*TankState) Descriptor() ([]byte, []int) {
	return file_tanks_proto_rawDescGZIP(), []int{0}
}

func (x *TankState) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

func (x *TankState) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *TankState) GetIsAI() bool {
	if x != nil {
		return x.IsAI
	}
	return false
}

type TankGameChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransformStates map[uint32]*channeldpb.TransformState `protobuf:"bytes,1,rep,name=transformStates,proto3" json:"transformStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TankStates      map[uint32]*TankState                 `protobuf:"bytes,2,rep,name=tankStates,proto3" json:"tankStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TankGameChannelData) Reset() {
	*x = TankGameChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tanks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TankGameChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TankGameChannelData) ProtoMessage() {}

func (x *TankGameChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_tanks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TankGameChannelData.ProtoReflect.Descriptor instead.
func (*TankGameChannelData) Descriptor() ([]byte, []int) {
	return file_tanks_proto_rawDescGZIP(), []int{1}
}

func (x *TankGameChannelData) GetTransformStates() map[uint32]*channeldpb.TransformState {
	if x != nil {
		return x.TransformStates
	}
	return nil
}

func (x *TankGameChannelData) GetTankStates() map[uint32]*TankState {
	if x != nil {
		return x.TankStates
	}
	return nil
}

var File_tanks_proto protoreflect.FileDescriptor

var file_tanks_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x6e, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74,
	0x61, 0x6e, 0x6b, 0x73, 0x70, 0x62, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x64, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x09, 0x54, 0x61, 0x6e,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x41, 0x49,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x41, 0x49, 0x22, 0xf3, 0x02, 0x0a,
	0x13, 0x54, 0x61, 0x6e, 0x6b, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x5b, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x74, 0x61, 0x6e, 0x6b, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x6e, 0x6b, 0x47, 0x61, 0x6d, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x4c, 0x0a, 0x0a, 0x74, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x61, 0x6e, 0x6b, 0x73, 0x70, 0x62, 0x2e,
	0x54, 0x61, 0x6e, 0x6b, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x74, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x5e, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x64, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x51, 0x0a, 0x0f, 0x54, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x6e, 0x6b, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x61,
	0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x5f, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x64, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x74, 0x61, 0x6e, 0x6b,
	0x73, 0x2f, 0x74, 0x61, 0x6e, 0x6b, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x17, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x64, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x6e, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tanks_proto_rawDescOnce sync.Once
	file_tanks_proto_rawDescData = file_tanks_proto_rawDesc
)

func file_tanks_proto_rawDescGZIP() []byte {
	file_tanks_proto_rawDescOnce.Do(func() {
		file_tanks_proto_rawDescData = protoimpl.X.CompressGZIP(file_tanks_proto_rawDescData)
	})
	return file_tanks_proto_rawDescData
}

var file_tanks_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tanks_proto_goTypes = []interface{}{
	(*TankState)(nil),                 // 0: tankspb.TankState
	(*TankGameChannelData)(nil),       // 1: tankspb.TankGameChannelData
	nil,                               // 2: tankspb.TankGameChannelData.TransformStatesEntry
	nil,                               // 3: tankspb.TankGameChannelData.TankStatesEntry
	(*channeldpb.TransformState)(nil), // 4: channeldpb.TransformState
}
var file_tanks_proto_depIdxs = []int32{
	2, // 0: tankspb.TankGameChannelData.transformStates:type_name -> tankspb.TankGameChannelData.TransformStatesEntry
	3, // 1: tankspb.TankGameChannelData.tankStates:type_name -> tankspb.TankGameChannelData.TankStatesEntry
	4, // 2: tankspb.TankGameChannelData.TransformStatesEntry.value:type_name -> channeldpb.TransformState
	0, // 3: tankspb.TankGameChannelData.TankStatesEntry.value:type_name -> tankspb.TankState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tanks_proto_init() }
func file_tanks_proto_init() {
	if File_tanks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tanks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TankState); i {
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
		file_tanks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TankGameChannelData); i {
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
			RawDescriptor: file_tanks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tanks_proto_goTypes,
		DependencyIndexes: file_tanks_proto_depIdxs,
		MessageInfos:      file_tanks_proto_msgTypes,
	}.Build()
	File_tanks_proto = out.File
	file_tanks_proto_rawDesc = nil
	file_tanks_proto_goTypes = nil
	file_tanks_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: replay.proto

package replaypb

import (
	channeldpb "channeld.clewcat.com/channeld/pkg/channeldpb"
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

type ReplayPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffsetTime int64              `protobuf:"varint,1,opt,name=offsetTime,proto3" json:"offsetTime,omitempty"`
	Packet     *channeldpb.Packet `protobuf:"bytes,2,opt,name=packet,proto3" json:"packet,omitempty"`
}

func (x *ReplayPacket) Reset() {
	*x = ReplayPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplayPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplayPacket) ProtoMessage() {}

func (x *ReplayPacket) ProtoReflect() protoreflect.Message {
	mi := &file_replay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplayPacket.ProtoReflect.Descriptor instead.
func (*ReplayPacket) Descriptor() ([]byte, []int) {
	return file_replay_proto_rawDescGZIP(), []int{0}
}

func (x *ReplayPacket) GetOffsetTime() int64 {
	if x != nil {
		return x.OffsetTime
	}
	return 0
}

func (x *ReplayPacket) GetPacket() *channeldpb.Packet {
	if x != nil {
		return x.Packet
	}
	return nil
}

type ReplaySession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packets []*ReplayPacket `protobuf:"bytes,1,rep,name=packets,proto3" json:"packets,omitempty"`
}

func (x *ReplaySession) Reset() {
	*x = ReplaySession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaySession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaySession) ProtoMessage() {}

func (x *ReplaySession) ProtoReflect() protoreflect.Message {
	mi := &file_replay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaySession.ProtoReflect.Descriptor instead.
func (*ReplaySession) Descriptor() ([]byte, []int) {
	return file_replay_proto_rawDescGZIP(), []int{1}
}

func (x *ReplaySession) GetPackets() []*ReplayPacket {
	if x != nil {
		return x.Packets
	}
	return nil
}

var File_replay_proto protoreflect.FileDescriptor

var file_replay_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x70, 0x62, 0x1a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x64, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x64, 0x2e, 0x63, 0x6c, 0x65, 0x77, 0x63, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_replay_proto_rawDescOnce sync.Once
	file_replay_proto_rawDescData = file_replay_proto_rawDesc
)

func file_replay_proto_rawDescGZIP() []byte {
	file_replay_proto_rawDescOnce.Do(func() {
		file_replay_proto_rawDescData = protoimpl.X.CompressGZIP(file_replay_proto_rawDescData)
	})
	return file_replay_proto_rawDescData
}

var file_replay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_replay_proto_goTypes = []interface{}{
	(*ReplayPacket)(nil),      // 0: replaypb.ReplayPacket
	(*ReplaySession)(nil),     // 1: replaypb.ReplaySession
	(*channeldpb.Packet)(nil), // 2: channeldpb.Packet
}
var file_replay_proto_depIdxs = []int32{
	2, // 0: replaypb.ReplayPacket.packet:type_name -> channeldpb.Packet
	0, // 1: replaypb.ReplaySession.packets:type_name -> replaypb.ReplayPacket
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_replay_proto_init() }
func file_replay_proto_init() {
	if File_replay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_replay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplayPacket); i {
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
		file_replay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaySession); i {
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
			RawDescriptor: file_replay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_replay_proto_goTypes,
		DependencyIndexes: file_replay_proto_depIdxs,
		MessageInfos:      file_replay_proto_msgTypes,
	}.Build()
	File_replay_proto = out.File
	file_replay_proto_rawDesc = nil
	file_replay_proto_goTypes = nil
	file_replay_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: busproto/bus.proto

package busproto

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

// The request message containing the user's name.
type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   *string `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Message []byte  `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_busproto_bus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_busproto_bus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_busproto_bus_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetTopic() string {
	if x != nil && x.Topic != nil {
		return *x.Topic
	}
	return ""
}

func (x *MessageRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

// The response message containing the greetings
type MessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *int64 `protobuf:"varint,1,req,name=ts" json:"ts,omitempty"`
}

func (x *MessageReply) Reset() {
	*x = MessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_busproto_bus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReply) ProtoMessage() {}

func (x *MessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_busproto_bus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReply.ProtoReflect.Descriptor instead.
func (*MessageReply) Descriptor() ([]byte, []int) {
	return file_busproto_bus_proto_rawDescGZIP(), []int{1}
}

func (x *MessageReply) GetTs() int64 {
	if x != nil && x.Ts != nil {
		return *x.Ts
	}
	return 0
}

// subscribe to a topic
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriber *string `protobuf:"bytes,1,req,name=subscriber" json:"subscriber,omitempty"`
	Topic      *string `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_busproto_bus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_busproto_bus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_busproto_bus_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetSubscriber() string {
	if x != nil && x.Subscriber != nil {
		return *x.Subscriber
	}
	return ""
}

func (x *SubscribeRequest) GetTopic() string {
	if x != nil && x.Topic != nil {
		return *x.Topic
	}
	return ""
}

// The response message containing the greetings
type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Done *bool `protobuf:"varint,1,req,name=done" json:"done,omitempty"`
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_busproto_bus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_busproto_bus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReply.ProtoReflect.Descriptor instead.
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return file_busproto_bus_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeReply) GetDone() bool {
	if x != nil && x.Done != nil {
		return *x.Done
	}
	return false
}

var File_busproto_bus_proto protoreflect.FileDescriptor

var file_busproto_bus_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x22, 0x48, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x32, 0x68, 0x0a, 0x03, 0x42, 0x75, 0x73, 0x12, 0x2f, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x08, 0x53, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x15, 0x5a, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2f, 0x62,
	0x75, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_busproto_bus_proto_rawDescOnce sync.Once
	file_busproto_bus_proto_rawDescData = file_busproto_bus_proto_rawDesc
)

func file_busproto_bus_proto_rawDescGZIP() []byte {
	file_busproto_bus_proto_rawDescOnce.Do(func() {
		file_busproto_bus_proto_rawDescData = protoimpl.X.CompressGZIP(file_busproto_bus_proto_rawDescData)
	})
	return file_busproto_bus_proto_rawDescData
}

var file_busproto_bus_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_busproto_bus_proto_goTypes = []interface{}{
	(*MessageRequest)(nil),   // 0: MessageRequest
	(*MessageReply)(nil),     // 1: MessageReply
	(*SubscribeRequest)(nil), // 2: SubscribeRequest
	(*SubscribeReply)(nil),   // 3: SubscribeReply
}
var file_busproto_bus_proto_depIdxs = []int32{
	0, // 0: Bus.SendMessage:input_type -> MessageRequest
	2, // 1: Bus.SubTopic:input_type -> SubscribeRequest
	1, // 2: Bus.SendMessage:output_type -> MessageReply
	3, // 3: Bus.SubTopic:output_type -> SubscribeReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_busproto_bus_proto_init() }
func file_busproto_bus_proto_init() {
	if File_busproto_bus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_busproto_bus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_busproto_bus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReply); i {
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
		file_busproto_bus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_busproto_bus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReply); i {
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
			RawDescriptor: file_busproto_bus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_busproto_bus_proto_goTypes,
		DependencyIndexes: file_busproto_bus_proto_depIdxs,
		MessageInfos:      file_busproto_bus_proto_msgTypes,
	}.Build()
	File_busproto_bus_proto = out.File
	file_busproto_bus_proto_rawDesc = nil
	file_busproto_bus_proto_goTypes = nil
	file_busproto_bus_proto_depIdxs = nil
}
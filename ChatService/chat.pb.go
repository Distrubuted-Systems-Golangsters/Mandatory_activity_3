// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: ChatService/chat.proto

package ChatService

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

type ChatMessageClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int32  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ChatMessageClient) Reset() {
	*x = ChatMessageClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatService_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageClient) ProtoMessage() {}

func (x *ChatMessageClient) ProtoReflect() protoreflect.Message {
	mi := &file_ChatService_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageClient.ProtoReflect.Descriptor instead.
func (*ChatMessageClient) Descriptor() ([]byte, []int) {
	return file_ChatService_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessageClient) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMessageClient) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatMessageClient) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ClientName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Timestamp  int32  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ClientName) Reset() {
	*x = ClientName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatService_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientName) ProtoMessage() {}

func (x *ClientName) ProtoReflect() protoreflect.Message {
	mi := &file_ChatService_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientName.ProtoReflect.Descriptor instead.
func (*ClientName) Descriptor() ([]byte, []int) {
	return file_ChatService_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ClientName) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *ClientName) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int32  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatService_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ChatService_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_ChatService_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ServerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ServerResponse) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatService_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ChatService_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ChatService_chat_proto_rawDescGZIP(), []int{3}
}

var File_ChatService_chat_proto protoreflect.FileDescriptor

var file_ChatService_chat_proto_rawDesc = []byte{
	0x0a, 0x16, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x63, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4a, 0x0a, 0x0a, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x12, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x19, 0x5a, 0x17, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x2d, 0x43, 0x68, 0x61, 0x74,
	0x2f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChatService_chat_proto_rawDescOnce sync.Once
	file_ChatService_chat_proto_rawDescData = file_ChatService_chat_proto_rawDesc
)

func file_ChatService_chat_proto_rawDescGZIP() []byte {
	file_ChatService_chat_proto_rawDescOnce.Do(func() {
		file_ChatService_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChatService_chat_proto_rawDescData)
	})
	return file_ChatService_chat_proto_rawDescData
}

var file_ChatService_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ChatService_chat_proto_goTypes = []any{
	(*ChatMessageClient)(nil), // 0: ChatService.ChatMessageClient
	(*ClientName)(nil),        // 1: ChatService.ClientName
	(*ServerResponse)(nil),    // 2: ChatService.ServerResponse
	(*Empty)(nil),             // 3: ChatService.Empty
}
var file_ChatService_chat_proto_depIdxs = []int32{
	0, // 0: ChatService.ChatService.AddClient:input_type -> ChatService.ChatMessageClient
	1, // 1: ChatService.ChatService.LeaveChat:input_type -> ChatService.ClientName
	2, // 2: ChatService.ChatService.AddClient:output_type -> ChatService.ServerResponse
	3, // 3: ChatService.ChatService.LeaveChat:output_type -> ChatService.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChatService_chat_proto_init() }
func file_ChatService_chat_proto_init() {
	if File_ChatService_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChatService_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChatMessageClient); i {
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
		file_ChatService_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ClientName); i {
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
		file_ChatService_chat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ServerResponse); i {
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
		file_ChatService_chat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_ChatService_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ChatService_chat_proto_goTypes,
		DependencyIndexes: file_ChatService_chat_proto_depIdxs,
		MessageInfos:      file_ChatService_chat_proto_msgTypes,
	}.Build()
	File_ChatService_chat_proto = out.File
	file_ChatService_chat_proto_rawDesc = nil
	file_ChatService_chat_proto_goTypes = nil
	file_ChatService_chat_proto_depIdxs = nil
}

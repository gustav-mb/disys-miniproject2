// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: chatpb/chat.proto

package chatpb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active bool  `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Connect) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Connect) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{3}
}

type Done struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Done) Reset() {
	*x = Done{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Done) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Done) ProtoMessage() {}

func (x *Done) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Done.ProtoReflect.Descriptor instead.
func (*Done) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{4}
}

var File_chatpb_chat_proto protoreflect.FileDescriptor

var file_chatpb_chat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x22,
	0x46, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x67, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x47, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x32, 0xd1, 0x02, 0x0a, 0x09, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x1a, 0x13, 0x2e,
	0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x10, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e,
	0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x11, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x43, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79,
	0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x43,
	0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_chatpb_chat_proto_rawDescOnce sync.Once
	file_chatpb_chat_proto_rawDescData = file_chatpb_chat_proto_rawDesc
)

func file_chatpb_chat_proto_rawDescGZIP() []byte {
	file_chatpb_chat_proto_rawDescOnce.Do(func() {
		file_chatpb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatpb_chat_proto_rawDescData)
	})
	return file_chatpb_chat_proto_rawDescData
}

var file_chatpb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chatpb_chat_proto_goTypes = []interface{}{
	(*User)(nil),    // 0: ChittyChat.User
	(*Message)(nil), // 1: ChittyChat.Message
	(*Connect)(nil), // 2: ChittyChat.Connect
	(*Close)(nil),   // 3: ChittyChat.Close
	(*Done)(nil),    // 4: ChittyChat.Done
}
var file_chatpb_chat_proto_depIdxs = []int32{
	0, // 0: ChittyChat.Message.user:type_name -> ChittyChat.User
	0, // 1: ChittyChat.Connect.user:type_name -> ChittyChat.User
	2, // 2: ChittyChat.Broadcast.CreateStream:input_type -> ChittyChat.Connect
	0, // 3: ChittyChat.Broadcast.DeleteStream:input_type -> ChittyChat.User
	0, // 4: ChittyChat.Broadcast.DisconnectStream:input_type -> ChittyChat.User
	0, // 5: ChittyChat.Broadcast.ConnectStream:input_type -> ChittyChat.User
	1, // 6: ChittyChat.Broadcast.broadcast:input_type -> ChittyChat.Message
	1, // 7: ChittyChat.Broadcast.Publish:input_type -> ChittyChat.Message
	1, // 8: ChittyChat.Broadcast.CreateStream:output_type -> ChittyChat.Message
	3, // 9: ChittyChat.Broadcast.DeleteStream:output_type -> ChittyChat.Close
	3, // 10: ChittyChat.Broadcast.DisconnectStream:output_type -> ChittyChat.Close
	4, // 11: ChittyChat.Broadcast.ConnectStream:output_type -> ChittyChat.Done
	3, // 12: ChittyChat.Broadcast.broadcast:output_type -> ChittyChat.Close
	4, // 13: ChittyChat.Broadcast.Publish:output_type -> ChittyChat.Done
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chatpb_chat_proto_init() }
func file_chatpb_chat_proto_init() {
	if File_chatpb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatpb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_chatpb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chatpb_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
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
		file_chatpb_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
		file_chatpb_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Done); i {
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
			RawDescriptor: file_chatpb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatpb_chat_proto_goTypes,
		DependencyIndexes: file_chatpb_chat_proto_depIdxs,
		MessageInfos:      file_chatpb_chat_proto_msgTypes,
	}.Build()
	File_chatpb_chat_proto = out.File
	file_chatpb_chat_proto_rawDesc = nil
	file_chatpb_chat_proto_goTypes = nil
	file_chatpb_chat_proto_depIdxs = nil
}

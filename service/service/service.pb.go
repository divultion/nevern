// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0--rc3
// source: service/service.proto

package service

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[0]
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
	return file_service_service_proto_rawDescGZIP(), []int{0}
}

type WriteInputByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataWritten int64 `protobuf:"varint,1,opt,name=dataWritten,proto3" json:"dataWritten,omitempty"`
}

func (x *WriteInputByIdResponse) Reset() {
	*x = WriteInputByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteInputByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteInputByIdResponse) ProtoMessage() {}

func (x *WriteInputByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteInputByIdResponse.ProtoReflect.Descriptor instead.
func (*WriteInputByIdResponse) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *WriteInputByIdResponse) GetDataWritten() int64 {
	if x != nil {
		return x.DataWritten
	}
	return 0
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Id   *ConnectionId `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{2}
}

func (x *Input) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Input) GetId() *ConnectionId {
	if x != nil {
		return x.Id
	}
	return nil
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Ok   bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{3}
}

func (x *Output) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Output) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ConnectionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                *ConnectionId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address           string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	MessagesAvailable uint32        `protobuf:"varint,3,opt,name=messagesAvailable,proto3" json:"messagesAvailable,omitempty"`
	Connected         bool          `protobuf:"varint,4,opt,name=connected,proto3" json:"connected,omitempty"`
}

func (x *ConnectionData) Reset() {
	*x = ConnectionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionData) ProtoMessage() {}

func (x *ConnectionData) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionData.ProtoReflect.Descriptor instead.
func (*ConnectionData) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectionData) GetId() *ConnectionId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ConnectionData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectionData) GetMessagesAvailable() uint32 {
	if x != nil {
		return x.MessagesAvailable
	}
	return 0
}

func (x *ConnectionData) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

type ConnectionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawId []byte `protobuf:"bytes,1,opt,name=rawId,proto3" json:"rawId,omitempty"`
}

func (x *ConnectionId) Reset() {
	*x = ConnectionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionId) ProtoMessage() {}

func (x *ConnectionId) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionId.ProtoReflect.Descriptor instead.
func (*ConnectionId) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectionId) GetRawId() []byte {
	if x != nil {
		return x.RawId
	}
	return nil
}

var File_service_service_proto protoreflect.FileDescriptor

var file_service_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3a, 0x0a, 0x16, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x57, 0x72, 0x69, 0x74, 0x74,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x57, 0x72,
	0x69, 0x74, 0x74, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x25, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x06, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x77, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x49, 0x64, 0x32, 0xc7, 0x02,
	0x0a, 0x0d, 0x4e, 0x65, 0x76, 0x65, 0x72, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3d, 0x0a, 0x11, 0x54, 0x72, 0x79, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x78, 0x6f, 0x69, 0x64, 0x2f, 0x6e, 0x65, 0x76, 0x65,
	0x72, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_service_proto_rawDescOnce sync.Once
	file_service_service_proto_rawDescData = file_service_service_proto_rawDesc
)

func file_service_service_proto_rawDescGZIP() []byte {
	file_service_service_proto_rawDescOnce.Do(func() {
		file_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_service_proto_rawDescData)
	})
	return file_service_service_proto_rawDescData
}

var file_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_service_proto_goTypes = []any{
	(*Empty)(nil),                  // 0: service.Empty
	(*WriteInputByIdResponse)(nil), // 1: service.WriteInputByIdResponse
	(*Input)(nil),                  // 2: service.Input
	(*Output)(nil),                 // 3: service.Output
	(*ConnectionData)(nil),         // 4: service.ConnectionData
	(*ConnectionId)(nil),           // 5: service.ConnectionId
}
var file_service_service_proto_depIdxs = []int32{
	5, // 0: service.Input.id:type_name -> service.ConnectionId
	5, // 1: service.ConnectionData.id:type_name -> service.ConnectionId
	0, // 2: service.NevernService.ListConnectionIds:input_type -> service.Empty
	5, // 3: service.NevernService.TryReadOutputById:input_type -> service.ConnectionId
	2, // 4: service.NevernService.WriteInputById:input_type -> service.Input
	5, // 5: service.NevernService.ForgetById:input_type -> service.ConnectionId
	5, // 6: service.NevernService.DisconnectById:input_type -> service.ConnectionId
	4, // 7: service.NevernService.ListConnectionIds:output_type -> service.ConnectionData
	3, // 8: service.NevernService.TryReadOutputById:output_type -> service.Output
	1, // 9: service.NevernService.WriteInputById:output_type -> service.WriteInputByIdResponse
	0, // 10: service.NevernService.ForgetById:output_type -> service.Empty
	0, // 11: service.NevernService.DisconnectById:output_type -> service.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_service_proto_init() }
func file_service_service_proto_init() {
	if File_service_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_service_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WriteInputByIdResponse); i {
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
		file_service_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Input); i {
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
		file_service_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Output); i {
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
		file_service_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionData); i {
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
		file_service_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionId); i {
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
			RawDescriptor: file_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_service_proto_goTypes,
		DependencyIndexes: file_service_service_proto_depIdxs,
		MessageInfos:      file_service_service_proto_msgTypes,
	}.Build()
	File_service_service_proto = out.File
	file_service_service_proto_rawDesc = nil
	file_service_service_proto_goTypes = nil
	file_service_service_proto_depIdxs = nil
}
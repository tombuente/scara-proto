// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: scara.proto

package scara_proto

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

type OperationState int32

const (
	OperationState_QUEUED     OperationState = 0
	OperationState_PROCESSING OperationState = 1
	OperationState_DONE       OperationState = 2
)

// Enum value maps for OperationState.
var (
	OperationState_name = map[int32]string{
		0: "QUEUED",
		1: "PROCESSING",
		2: "DONE",
	}
	OperationState_value = map[string]int32{
		"QUEUED":     0,
		"PROCESSING": 1,
		"DONE":       2,
	}
)

func (x OperationState) Enum() *OperationState {
	p := new(OperationState)
	*p = x
	return p
}

func (x OperationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationState) Descriptor() protoreflect.EnumDescriptor {
	return file_scara_proto_enumTypes[0].Descriptor()
}

func (OperationState) Type() protoreflect.EnumType {
	return &file_scara_proto_enumTypes[0]
}

func (x OperationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationState.Descriptor instead.
func (OperationState) EnumDescriptor() ([]byte, []int) {
	return file_scara_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scara_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_scara_proto_msgTypes[0]
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
	return file_scara_proto_rawDescGZIP(), []int{0}
}

type ExecCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *ExecCommandRequest) Reset() {
	*x = ExecCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scara_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandRequest) ProtoMessage() {}

func (x *ExecCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scara_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandRequest.ProtoReflect.Descriptor instead.
func (*ExecCommandRequest) Descriptor() ([]byte, []int) {
	return file_scara_proto_rawDescGZIP(), []int{1}
}

func (x *ExecCommandRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type ExecCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExecCommandResponse) Reset() {
	*x = ExecCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scara_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandResponse) ProtoMessage() {}

func (x *ExecCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scara_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandResponse.ProtoReflect.Descriptor instead.
func (*ExecCommandResponse) Descriptor() ([]byte, []int) {
	return file_scara_proto_rawDescGZIP(), []int{2}
}

func (x *ExecCommandResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommandInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	State OperationState `protobuf:"varint,2,opt,name=state,proto3,enum=scaraproto.OperationState" json:"state,omitempty"`
}

func (x *CommandInfo) Reset() {
	*x = CommandInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scara_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandInfo) ProtoMessage() {}

func (x *CommandInfo) ProtoReflect() protoreflect.Message {
	mi := &file_scara_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandInfo.ProtoReflect.Descriptor instead.
func (*CommandInfo) Descriptor() ([]byte, []int) {
	return file_scara_proto_rawDescGZIP(), []int{3}
}

func (x *CommandInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommandInfo) GetState() OperationState {
	if x != nil {
		return x.State
	}
	return OperationState_QUEUED
}

var File_scara_proto protoreflect.FileDescriptor

var file_scara_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x61, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x63, 0x61, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x2e, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x36, 0x0a, 0x0e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06,
	0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45,
	0x10, 0x02, 0x32, 0x99, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x50, 0x0a, 0x0b,
	0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x63,
	0x61, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x63,
	0x61, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x11,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d,
	0x62, 0x75, 0x65, 0x6e, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x61, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scara_proto_rawDescOnce sync.Once
	file_scara_proto_rawDescData = file_scara_proto_rawDesc
)

func file_scara_proto_rawDescGZIP() []byte {
	file_scara_proto_rawDescOnce.Do(func() {
		file_scara_proto_rawDescData = protoimpl.X.CompressGZIP(file_scara_proto_rawDescData)
	})
	return file_scara_proto_rawDescData
}

var file_scara_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scara_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_scara_proto_goTypes = []interface{}{
	(OperationState)(0),         // 0: scaraproto.OperationState
	(*Empty)(nil),               // 1: scaraproto.Empty
	(*ExecCommandRequest)(nil),  // 2: scaraproto.ExecCommandRequest
	(*ExecCommandResponse)(nil), // 3: scaraproto.ExecCommandResponse
	(*CommandInfo)(nil),         // 4: scaraproto.CommandInfo
}
var file_scara_proto_depIdxs = []int32{
	0, // 0: scaraproto.CommandInfo.state:type_name -> scaraproto.OperationState
	2, // 1: scaraproto.Robot.ExecCommand:input_type -> scaraproto.ExecCommandRequest
	1, // 2: scaraproto.Robot.QueueUpdates:input_type -> scaraproto.Empty
	3, // 3: scaraproto.Robot.ExecCommand:output_type -> scaraproto.ExecCommandResponse
	4, // 4: scaraproto.Robot.QueueUpdates:output_type -> scaraproto.CommandInfo
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scara_proto_init() }
func file_scara_proto_init() {
	if File_scara_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scara_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_scara_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandRequest); i {
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
		file_scara_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandResponse); i {
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
		file_scara_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandInfo); i {
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
			RawDescriptor: file_scara_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scara_proto_goTypes,
		DependencyIndexes: file_scara_proto_depIdxs,
		EnumInfos:         file_scara_proto_enumTypes,
		MessageInfos:      file_scara_proto_msgTypes,
	}.Build()
	File_scara_proto = out.File
	file_scara_proto_rawDesc = nil
	file_scara_proto_goTypes = nil
	file_scara_proto_depIdxs = nil
}

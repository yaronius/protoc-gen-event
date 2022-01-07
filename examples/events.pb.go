// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: examples/events.proto

package examples

import (
	_ "github.com/voi-oss/protoc-gen-event/pkg/options"
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

type NotifyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string            `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Data      map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NotifyEvent) Reset() {
	*x = NotifyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyEvent) ProtoMessage() {}

func (x *NotifyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_examples_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyEvent.ProtoReflect.Descriptor instead.
func (*NotifyEvent) Descriptor() ([]byte, []int) {
	return file_examples_events_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyEvent) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *NotifyEvent) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type CustomTopicEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string            `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Data      map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CustomTopicEvent) Reset() {
	*x = CustomTopicEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTopicEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTopicEvent) ProtoMessage() {}

func (x *CustomTopicEvent) ProtoReflect() protoreflect.Message {
	mi := &file_examples_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTopicEvent.ProtoReflect.Descriptor instead.
func (*CustomTopicEvent) Descriptor() ([]byte, []int) {
	return file_examples_events_proto_rawDescGZIP(), []int{1}
}

func (x *CustomTopicEvent) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *CustomTopicEvent) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type NotAnEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string            `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Data      map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NotAnEvent) Reset() {
	*x = NotAnEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotAnEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotAnEvent) ProtoMessage() {}

func (x *NotAnEvent) ProtoReflect() protoreflect.Message {
	mi := &file_examples_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotAnEvent.ProtoReflect.Descriptor instead.
func (*NotAnEvent) Descriptor() ([]byte, []int) {
	return file_examples_events_proto_rawDescGZIP(), []int{2}
}

func (x *NotAnEvent) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *NotAnEvent) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_examples_events_proto protoreflect.FileDescriptor

var file_examples_events_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x76, 0x6f, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x67, 0x65, 0x6e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x1a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x46, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x76, 0x6f, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x67, 0x65, 0x6e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xc8, 0x01, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x4b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x76, 0x6f, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x67,
	0x65, 0x6e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x10, 0xf2, 0xbd, 0x18, 0x0c,
	0x0a, 0x0a, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0xb2, 0x01, 0x0a,
	0x0a, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x76, 0x6f, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x67, 0x65, 0x6e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x06, 0xf2, 0xbd, 0x18, 0x02, 0x10,
	0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x6f, 0x69, 0x2d, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_events_proto_rawDescOnce sync.Once
	file_examples_events_proto_rawDescData = file_examples_events_proto_rawDesc
)

func file_examples_events_proto_rawDescGZIP() []byte {
	file_examples_events_proto_rawDescOnce.Do(func() {
		file_examples_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_events_proto_rawDescData)
	})
	return file_examples_events_proto_rawDescData
}

var file_examples_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_examples_events_proto_goTypes = []interface{}{
	(*NotifyEvent)(nil),      // 0: voi.protocgenevent.examples.NotifyEvent
	(*CustomTopicEvent)(nil), // 1: voi.protocgenevent.examples.CustomTopicEvent
	(*NotAnEvent)(nil),       // 2: voi.protocgenevent.examples.NotAnEvent
	nil,                      // 3: voi.protocgenevent.examples.NotifyEvent.DataEntry
	nil,                      // 4: voi.protocgenevent.examples.CustomTopicEvent.DataEntry
	nil,                      // 5: voi.protocgenevent.examples.NotAnEvent.DataEntry
}
var file_examples_events_proto_depIdxs = []int32{
	3, // 0: voi.protocgenevent.examples.NotifyEvent.data:type_name -> voi.protocgenevent.examples.NotifyEvent.DataEntry
	4, // 1: voi.protocgenevent.examples.CustomTopicEvent.data:type_name -> voi.protocgenevent.examples.CustomTopicEvent.DataEntry
	5, // 2: voi.protocgenevent.examples.NotAnEvent.data:type_name -> voi.protocgenevent.examples.NotAnEvent.DataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_examples_events_proto_init() }
func file_examples_events_proto_init() {
	if File_examples_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyEvent); i {
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
		file_examples_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTopicEvent); i {
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
		file_examples_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotAnEvent); i {
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
			RawDescriptor: file_examples_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_events_proto_goTypes,
		DependencyIndexes: file_examples_events_proto_depIdxs,
		MessageInfos:      file_examples_events_proto_msgTypes,
	}.Build()
	File_examples_events_proto = out.File
	file_examples_events_proto_rawDesc = nil
	file_examples_events_proto_goTypes = nil
	file_examples_events_proto_depIdxs = nil
}

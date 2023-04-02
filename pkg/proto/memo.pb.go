// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: memo.proto

package proto

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

type Memosphere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memos []*Memosphere_Memo `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
}

func (x *Memosphere) Reset() {
	*x = Memosphere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memosphere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memosphere) ProtoMessage() {}

func (x *Memosphere) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memosphere.ProtoReflect.Descriptor instead.
func (*Memosphere) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0}
}

func (x *Memosphere) GetMemos() []*Memosphere_Memo {
	if x != nil {
		return x.Memos
	}
	return nil
}

type Memosphere_Memo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *Thunk             `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Calls  []*Memosphere_Call `protobuf:"bytes,2,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *Memosphere_Memo) Reset() {
	*x = Memosphere_Memo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memosphere_Memo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memosphere_Memo) ProtoMessage() {}

func (x *Memosphere_Memo) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memosphere_Memo.ProtoReflect.Descriptor instead.
func (*Memosphere_Memo) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Memosphere_Memo) GetModule() *Thunk {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *Memosphere_Memo) GetCalls() []*Memosphere_Call {
	if x != nil {
		return x.Calls
	}
	return nil
}

type Memosphere_Call struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binding string               `protobuf:"bytes,1,opt,name=binding,proto3" json:"binding,omitempty"`
	Results []*Memosphere_Result `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Memosphere_Call) Reset() {
	*x = Memosphere_Call{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memosphere_Call) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memosphere_Call) ProtoMessage() {}

func (x *Memosphere_Call) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memosphere_Call.ProtoReflect.Descriptor instead.
func (*Memosphere_Call) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Memosphere_Call) GetBinding() string {
	if x != nil {
		return x.Binding
	}
	return ""
}

func (x *Memosphere_Call) GetResults() []*Memosphere_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Memosphere_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  *Value `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output *Value `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Memosphere_Result) Reset() {
	*x = Memosphere_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memosphere_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memosphere_Result) ProtoMessage() {}

func (x *Memosphere_Result) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memosphere_Result.ProtoReflect.Descriptor instead.
func (*Memosphere_Result) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Memosphere_Result) GetInput() *Value {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Memosphere_Result) GetOutput() *Value {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_memo_proto protoreflect.FileDescriptor

var file_memo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x73, 0x73, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba,
	0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x12, 0x2b, 0x0a,
	0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62,
	0x61, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x4d,
	0x65, 0x6d, 0x6f, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x1a, 0x58, 0x0a, 0x04, 0x4d, 0x65,
	0x6d, 0x6f, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x4d, 0x65,
	0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x63,
	0x61, 0x6c, 0x6c, 0x73, 0x1a, 0x53, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x4d,
	0x65, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memo_proto_rawDescOnce sync.Once
	file_memo_proto_rawDescData = file_memo_proto_rawDesc
)

func file_memo_proto_rawDescGZIP() []byte {
	file_memo_proto_rawDescOnce.Do(func() {
		file_memo_proto_rawDescData = protoimpl.X.CompressGZIP(file_memo_proto_rawDescData)
	})
	return file_memo_proto_rawDescData
}

var file_memo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memo_proto_goTypes = []interface{}{
	(*Memosphere)(nil),        // 0: bass.Memosphere
	(*Memosphere_Memo)(nil),   // 1: bass.Memosphere.Memo
	(*Memosphere_Call)(nil),   // 2: bass.Memosphere.Call
	(*Memosphere_Result)(nil), // 3: bass.Memosphere.Result
	(*Thunk)(nil),             // 4: bass.Thunk
	(*Value)(nil),             // 5: bass.Value
}
var file_memo_proto_depIdxs = []int32{
	1, // 0: bass.Memosphere.memos:type_name -> bass.Memosphere.Memo
	4, // 1: bass.Memosphere.Memo.module:type_name -> bass.Thunk
	2, // 2: bass.Memosphere.Memo.calls:type_name -> bass.Memosphere.Call
	3, // 3: bass.Memosphere.Call.results:type_name -> bass.Memosphere.Result
	5, // 4: bass.Memosphere.Result.input:type_name -> bass.Value
	5, // 5: bass.Memosphere.Result.output:type_name -> bass.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_memo_proto_init() }
func file_memo_proto_init() {
	if File_memo_proto != nil {
		return
	}
	file_bass_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memosphere); i {
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
		file_memo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memosphere_Memo); i {
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
		file_memo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memosphere_Call); i {
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
		file_memo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memosphere_Result); i {
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
			RawDescriptor: file_memo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_memo_proto_goTypes,
		DependencyIndexes: file_memo_proto_depIdxs,
		MessageInfos:      file_memo_proto_msgTypes,
	}.Build()
	File_memo_proto = out.File
	file_memo_proto_rawDesc = nil
	file_memo_proto_goTypes = nil
	file_memo_proto_depIdxs = nil
}

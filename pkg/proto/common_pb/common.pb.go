// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/common/v1/common.proto

package common_pb

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

type Observation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time         int64              `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Measurements map[string]float64 `protobuf:"bytes,2,rep,name=measurements,proto3" json:"measurements,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Categories   map[string]string  `protobuf:"bytes,3,rep,name=categories,proto3" json:"categories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags         []string           `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Observation) Reset() {
	*x = Observation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observation.ProtoReflect.Descriptor instead.
func (*Observation) Descriptor() ([]byte, []int) {
	return file_proto_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Observation) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Observation) GetMeasurements() map[string]float64 {
	if x != nil {
		return x.Measurements
	}
	return nil
}

func (x *Observation) GetCategories() map[string]string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Observation) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Interpretation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start   int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End     int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Actions []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Tags    []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Interpretation) Reset() {
	*x = Interpretation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interpretation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interpretation) ProtoMessage() {}

func (x *Interpretation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interpretation.ProtoReflect.Descriptor instead.
func (*Interpretation) Descriptor() ([]byte, []int) {
	return file_proto_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Interpretation) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Interpretation) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Interpretation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Interpretation) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Interpretation) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type InterpretationIndices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indicies []uint32 `protobuf:"varint,1,rep,packed,name=indicies,proto3" json:"indicies,omitempty"`
}

func (x *InterpretationIndices) Reset() {
	*x = InterpretationIndices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterpretationIndices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterpretationIndices) ProtoMessage() {}

func (x *InterpretationIndices) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterpretationIndices.ProtoReflect.Descriptor instead.
func (*InterpretationIndices) Descriptor() ([]byte, []int) {
	return file_proto_common_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *InterpretationIndices) GetIndicies() []uint32 {
	if x != nil {
		return x.Indicies
	}
	return nil
}

type IndexedInterpretations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interpretations []*Interpretation                `protobuf:"bytes,1,rep,name=interpretations,proto3" json:"interpretations,omitempty"`
	Index           map[int64]*InterpretationIndices `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IndexedInterpretations) Reset() {
	*x = IndexedInterpretations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexedInterpretations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexedInterpretations) ProtoMessage() {}

func (x *IndexedInterpretations) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexedInterpretations.ProtoReflect.Descriptor instead.
func (*IndexedInterpretations) Descriptor() ([]byte, []int) {
	return file_proto_common_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *IndexedInterpretations) GetInterpretations() []*Interpretation {
	if x != nil {
		return x.Interpretations
	}
	return nil
}

func (x *IndexedInterpretations) GetIndex() map[int64]*InterpretationIndices {
	if x != nil {
		return x.Index
	}
	return nil
}

var File_proto_common_v1_common_proto protoreflect.FileDescriptor

var file_proto_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xc5, 0x02, 0x0a, 0x0b, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x3f,
	0x0a, 0x11, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3d, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7a,
	0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x37, 0x0a, 0x15, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40,
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3f, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0x57, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x63, 0x65, 0x61, 0x69,
	0x2f, 0x73, 0x70, 0x69, 0x63, 0x65, 0x61, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_v1_common_proto_rawDescOnce sync.Once
	file_proto_common_v1_common_proto_rawDescData = file_proto_common_v1_common_proto_rawDesc
)

func file_proto_common_v1_common_proto_rawDescGZIP() []byte {
	file_proto_common_v1_common_proto_rawDescOnce.Do(func() {
		file_proto_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_v1_common_proto_rawDescData)
	})
	return file_proto_common_v1_common_proto_rawDescData
}

var file_proto_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_common_v1_common_proto_goTypes = []interface{}{
	(*Observation)(nil),            // 0: common.Observation
	(*Interpretation)(nil),         // 1: common.Interpretation
	(*InterpretationIndices)(nil),  // 2: common.InterpretationIndices
	(*IndexedInterpretations)(nil), // 3: common.IndexedInterpretations
	nil,                            // 4: common.Observation.MeasurementsEntry
	nil,                            // 5: common.Observation.CategoriesEntry
	nil,                            // 6: common.IndexedInterpretations.IndexEntry
}
var file_proto_common_v1_common_proto_depIdxs = []int32{
	4, // 0: common.Observation.measurements:type_name -> common.Observation.MeasurementsEntry
	5, // 1: common.Observation.categories:type_name -> common.Observation.CategoriesEntry
	1, // 2: common.IndexedInterpretations.interpretations:type_name -> common.Interpretation
	6, // 3: common.IndexedInterpretations.index:type_name -> common.IndexedInterpretations.IndexEntry
	2, // 4: common.IndexedInterpretations.IndexEntry.value:type_name -> common.InterpretationIndices
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_common_v1_common_proto_init() }
func file_proto_common_v1_common_proto_init() {
	if File_proto_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Observation); i {
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
		file_proto_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interpretation); i {
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
		file_proto_common_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterpretationIndices); i {
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
		file_proto_common_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexedInterpretations); i {
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
			RawDescriptor: file_proto_common_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_v1_common_proto_goTypes,
		DependencyIndexes: file_proto_common_v1_common_proto_depIdxs,
		MessageInfos:      file_proto_common_v1_common_proto_msgTypes,
	}.Build()
	File_proto_common_v1_common_proto = out.File
	file_proto_common_v1_common_proto_rawDesc = nil
	file_proto_common_v1_common_proto_goTypes = nil
	file_proto_common_v1_common_proto_depIdxs = nil
}

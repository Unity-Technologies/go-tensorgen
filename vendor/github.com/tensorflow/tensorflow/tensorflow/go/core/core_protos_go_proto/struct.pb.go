// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/struct.proto

package core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TypeSpecProto_TypeSpecClass int32

const (
	TypeSpecProto_UNKNOWN             TypeSpecProto_TypeSpecClass = 0
	TypeSpecProto_SPARSE_TENSOR_SPEC  TypeSpecProto_TypeSpecClass = 1
	TypeSpecProto_INDEXED_SLICES_SPEC TypeSpecProto_TypeSpecClass = 2
	TypeSpecProto_RAGGED_TENSOR_SPEC  TypeSpecProto_TypeSpecClass = 3
	TypeSpecProto_TENSOR_ARRAY_SPEC   TypeSpecProto_TypeSpecClass = 4
	TypeSpecProto_DATA_DATASET_SPEC   TypeSpecProto_TypeSpecClass = 5
	TypeSpecProto_DATA_ITERATOR_SPEC  TypeSpecProto_TypeSpecClass = 6
	TypeSpecProto_OPTIONAL_SPEC       TypeSpecProto_TypeSpecClass = 7
	TypeSpecProto_PER_REPLICA_SPEC    TypeSpecProto_TypeSpecClass = 8
	TypeSpecProto_VARIABLE_SPEC       TypeSpecProto_TypeSpecClass = 9
)

var TypeSpecProto_TypeSpecClass_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPARSE_TENSOR_SPEC",
	2: "INDEXED_SLICES_SPEC",
	3: "RAGGED_TENSOR_SPEC",
	4: "TENSOR_ARRAY_SPEC",
	5: "DATA_DATASET_SPEC",
	6: "DATA_ITERATOR_SPEC",
	7: "OPTIONAL_SPEC",
	8: "PER_REPLICA_SPEC",
	9: "VARIABLE_SPEC",
}

var TypeSpecProto_TypeSpecClass_value = map[string]int32{
	"UNKNOWN":             0,
	"SPARSE_TENSOR_SPEC":  1,
	"INDEXED_SLICES_SPEC": 2,
	"RAGGED_TENSOR_SPEC":  3,
	"TENSOR_ARRAY_SPEC":   4,
	"DATA_DATASET_SPEC":   5,
	"DATA_ITERATOR_SPEC":  6,
	"OPTIONAL_SPEC":       7,
	"PER_REPLICA_SPEC":    8,
	"VARIABLE_SPEC":       9,
}

func (x TypeSpecProto_TypeSpecClass) String() string {
	return proto.EnumName(TypeSpecProto_TypeSpecClass_name, int32(x))
}

func (TypeSpecProto_TypeSpecClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{8, 0}
}

// `StructuredValue` represents a dynamically typed value representing various
// data structures that are inspired by Python data structures typically used in
// TensorFlow functions as inputs and outputs.
//
// For example when saving a Layer there may be a `training` argument. If the
// user passes a boolean True/False, that switches between two concrete
// TensorFlow functions. In order to switch between them in the same way after
// loading the SavedModel, we need to represent "True" and "False".
//
// A more advanced example might be a function which takes a list of
// dictionaries mapping from strings to Tensors. In order to map from
// user-specified arguments `[{"a": tf.constant(1.)}, {"q": tf.constant(3.)}]`
// after load to the right saved TensorFlow function, we need to represent the
// nested structure and the strings, recording that we have a trace for anything
// matching `[{"a": tf.TensorSpec(None, tf.float32)}, {"q": tf.TensorSpec([],
// tf.float64)}]` as an example.
//
// Likewise functions may return nested structures of Tensors, for example
// returning a dictionary mapping from strings to Tensors. In order for the
// loaded function to return the same structure we need to serialize it.
//
// This is an ergonomic aid for working with loaded SavedModels, not a promise
// to serialize all possible function signatures. For example we do not expect
// to pickle generic Python objects, and ideally we'd stay language-agnostic.
type StructuredValue struct {
	// The kind of value.
	//
	// Types that are valid to be assigned to Kind:
	//	*StructuredValue_NoneValue
	//	*StructuredValue_Float64Value
	//	*StructuredValue_Int64Value
	//	*StructuredValue_StringValue
	//	*StructuredValue_BoolValue
	//	*StructuredValue_TensorShapeValue
	//	*StructuredValue_TensorDtypeValue
	//	*StructuredValue_TensorSpecValue
	//	*StructuredValue_TypeSpecValue
	//	*StructuredValue_ListValue
	//	*StructuredValue_TupleValue
	//	*StructuredValue_DictValue
	//	*StructuredValue_NamedTupleValue
	Kind                 isStructuredValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StructuredValue) Reset()         { *m = StructuredValue{} }
func (m *StructuredValue) String() string { return proto.CompactTextString(m) }
func (*StructuredValue) ProtoMessage()    {}
func (*StructuredValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{0}
}

func (m *StructuredValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredValue.Unmarshal(m, b)
}
func (m *StructuredValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredValue.Marshal(b, m, deterministic)
}
func (m *StructuredValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredValue.Merge(m, src)
}
func (m *StructuredValue) XXX_Size() int {
	return xxx_messageInfo_StructuredValue.Size(m)
}
func (m *StructuredValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredValue.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredValue proto.InternalMessageInfo

type isStructuredValue_Kind interface {
	isStructuredValue_Kind()
}

type StructuredValue_NoneValue struct {
	NoneValue *NoneValue `protobuf:"bytes,1,opt,name=none_value,json=noneValue,proto3,oneof"`
}

type StructuredValue_Float64Value struct {
	Float64Value float64 `protobuf:"fixed64,11,opt,name=float64_value,json=float64Value,proto3,oneof"`
}

type StructuredValue_Int64Value struct {
	Int64Value int64 `protobuf:"zigzag64,12,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type StructuredValue_StringValue struct {
	StringValue string `protobuf:"bytes,13,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type StructuredValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,14,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type StructuredValue_TensorShapeValue struct {
	TensorShapeValue *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,31,opt,name=tensor_shape_value,json=tensorShapeValue,proto3,oneof"`
}

type StructuredValue_TensorDtypeValue struct {
	TensorDtypeValue types_go_proto.DataType `protobuf:"varint,32,opt,name=tensor_dtype_value,json=tensorDtypeValue,proto3,enum=tensorflow.DataType,oneof"`
}

type StructuredValue_TensorSpecValue struct {
	TensorSpecValue *TensorSpecProto `protobuf:"bytes,33,opt,name=tensor_spec_value,json=tensorSpecValue,proto3,oneof"`
}

type StructuredValue_TypeSpecValue struct {
	TypeSpecValue *TypeSpecProto `protobuf:"bytes,34,opt,name=type_spec_value,json=typeSpecValue,proto3,oneof"`
}

type StructuredValue_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,51,opt,name=list_value,json=listValue,proto3,oneof"`
}

type StructuredValue_TupleValue struct {
	TupleValue *TupleValue `protobuf:"bytes,52,opt,name=tuple_value,json=tupleValue,proto3,oneof"`
}

type StructuredValue_DictValue struct {
	DictValue *DictValue `protobuf:"bytes,53,opt,name=dict_value,json=dictValue,proto3,oneof"`
}

type StructuredValue_NamedTupleValue struct {
	NamedTupleValue *NamedTupleValue `protobuf:"bytes,54,opt,name=named_tuple_value,json=namedTupleValue,proto3,oneof"`
}

func (*StructuredValue_NoneValue) isStructuredValue_Kind() {}

func (*StructuredValue_Float64Value) isStructuredValue_Kind() {}

func (*StructuredValue_Int64Value) isStructuredValue_Kind() {}

func (*StructuredValue_StringValue) isStructuredValue_Kind() {}

func (*StructuredValue_BoolValue) isStructuredValue_Kind() {}

func (*StructuredValue_TensorShapeValue) isStructuredValue_Kind() {}

func (*StructuredValue_TensorDtypeValue) isStructuredValue_Kind() {}

func (*StructuredValue_TensorSpecValue) isStructuredValue_Kind() {}

func (*StructuredValue_TypeSpecValue) isStructuredValue_Kind() {}

func (*StructuredValue_ListValue) isStructuredValue_Kind() {}

func (*StructuredValue_TupleValue) isStructuredValue_Kind() {}

func (*StructuredValue_DictValue) isStructuredValue_Kind() {}

func (*StructuredValue_NamedTupleValue) isStructuredValue_Kind() {}

func (m *StructuredValue) GetKind() isStructuredValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *StructuredValue) GetNoneValue() *NoneValue {
	if x, ok := m.GetKind().(*StructuredValue_NoneValue); ok {
		return x.NoneValue
	}
	return nil
}

func (m *StructuredValue) GetFloat64Value() float64 {
	if x, ok := m.GetKind().(*StructuredValue_Float64Value); ok {
		return x.Float64Value
	}
	return 0
}

func (m *StructuredValue) GetInt64Value() int64 {
	if x, ok := m.GetKind().(*StructuredValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *StructuredValue) GetStringValue() string {
	if x, ok := m.GetKind().(*StructuredValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *StructuredValue) GetBoolValue() bool {
	if x, ok := m.GetKind().(*StructuredValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *StructuredValue) GetTensorShapeValue() *tensor_shape_go_proto.TensorShapeProto {
	if x, ok := m.GetKind().(*StructuredValue_TensorShapeValue); ok {
		return x.TensorShapeValue
	}
	return nil
}

func (m *StructuredValue) GetTensorDtypeValue() types_go_proto.DataType {
	if x, ok := m.GetKind().(*StructuredValue_TensorDtypeValue); ok {
		return x.TensorDtypeValue
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *StructuredValue) GetTensorSpecValue() *TensorSpecProto {
	if x, ok := m.GetKind().(*StructuredValue_TensorSpecValue); ok {
		return x.TensorSpecValue
	}
	return nil
}

func (m *StructuredValue) GetTypeSpecValue() *TypeSpecProto {
	if x, ok := m.GetKind().(*StructuredValue_TypeSpecValue); ok {
		return x.TypeSpecValue
	}
	return nil
}

func (m *StructuredValue) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*StructuredValue_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (m *StructuredValue) GetTupleValue() *TupleValue {
	if x, ok := m.GetKind().(*StructuredValue_TupleValue); ok {
		return x.TupleValue
	}
	return nil
}

func (m *StructuredValue) GetDictValue() *DictValue {
	if x, ok := m.GetKind().(*StructuredValue_DictValue); ok {
		return x.DictValue
	}
	return nil
}

func (m *StructuredValue) GetNamedTupleValue() *NamedTupleValue {
	if x, ok := m.GetKind().(*StructuredValue_NamedTupleValue); ok {
		return x.NamedTupleValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StructuredValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StructuredValue_NoneValue)(nil),
		(*StructuredValue_Float64Value)(nil),
		(*StructuredValue_Int64Value)(nil),
		(*StructuredValue_StringValue)(nil),
		(*StructuredValue_BoolValue)(nil),
		(*StructuredValue_TensorShapeValue)(nil),
		(*StructuredValue_TensorDtypeValue)(nil),
		(*StructuredValue_TensorSpecValue)(nil),
		(*StructuredValue_TypeSpecValue)(nil),
		(*StructuredValue_ListValue)(nil),
		(*StructuredValue_TupleValue)(nil),
		(*StructuredValue_DictValue)(nil),
		(*StructuredValue_NamedTupleValue)(nil),
	}
}

// Represents None.
type NoneValue struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoneValue) Reset()         { *m = NoneValue{} }
func (m *NoneValue) String() string { return proto.CompactTextString(m) }
func (*NoneValue) ProtoMessage()    {}
func (*NoneValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{1}
}

func (m *NoneValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoneValue.Unmarshal(m, b)
}
func (m *NoneValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoneValue.Marshal(b, m, deterministic)
}
func (m *NoneValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoneValue.Merge(m, src)
}
func (m *NoneValue) XXX_Size() int {
	return xxx_messageInfo_NoneValue.Size(m)
}
func (m *NoneValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NoneValue.DiscardUnknown(m)
}

var xxx_messageInfo_NoneValue proto.InternalMessageInfo

// Represents a Python list.
type ListValue struct {
	Values               []*StructuredValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{2}
}

func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListValue.Unmarshal(m, b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
}
func (m *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(m, src)
}
func (m *ListValue) XXX_Size() int {
	return xxx_messageInfo_ListValue.Size(m)
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

func (m *ListValue) GetValues() []*StructuredValue {
	if m != nil {
		return m.Values
	}
	return nil
}

// Represents a Python tuple.
type TupleValue struct {
	Values               []*StructuredValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TupleValue) Reset()         { *m = TupleValue{} }
func (m *TupleValue) String() string { return proto.CompactTextString(m) }
func (*TupleValue) ProtoMessage()    {}
func (*TupleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{3}
}

func (m *TupleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TupleValue.Unmarshal(m, b)
}
func (m *TupleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TupleValue.Marshal(b, m, deterministic)
}
func (m *TupleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TupleValue.Merge(m, src)
}
func (m *TupleValue) XXX_Size() int {
	return xxx_messageInfo_TupleValue.Size(m)
}
func (m *TupleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TupleValue.DiscardUnknown(m)
}

var xxx_messageInfo_TupleValue proto.InternalMessageInfo

func (m *TupleValue) GetValues() []*StructuredValue {
	if m != nil {
		return m.Values
	}
	return nil
}

// Represents a Python dict keyed by `str`.
// The comment on Unicode from Value.string_value applies analogously.
type DictValue struct {
	Fields               map[string]*StructuredValue `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DictValue) Reset()         { *m = DictValue{} }
func (m *DictValue) String() string { return proto.CompactTextString(m) }
func (*DictValue) ProtoMessage()    {}
func (*DictValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{4}
}

func (m *DictValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictValue.Unmarshal(m, b)
}
func (m *DictValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictValue.Marshal(b, m, deterministic)
}
func (m *DictValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictValue.Merge(m, src)
}
func (m *DictValue) XXX_Size() int {
	return xxx_messageInfo_DictValue.Size(m)
}
func (m *DictValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DictValue.DiscardUnknown(m)
}

var xxx_messageInfo_DictValue proto.InternalMessageInfo

func (m *DictValue) GetFields() map[string]*StructuredValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Represents a (key, value) pair.
type PairValue struct {
	Key                  string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *StructuredValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PairValue) Reset()         { *m = PairValue{} }
func (m *PairValue) String() string { return proto.CompactTextString(m) }
func (*PairValue) ProtoMessage()    {}
func (*PairValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{5}
}

func (m *PairValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairValue.Unmarshal(m, b)
}
func (m *PairValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairValue.Marshal(b, m, deterministic)
}
func (m *PairValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairValue.Merge(m, src)
}
func (m *PairValue) XXX_Size() int {
	return xxx_messageInfo_PairValue.Size(m)
}
func (m *PairValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PairValue.DiscardUnknown(m)
}

var xxx_messageInfo_PairValue proto.InternalMessageInfo

func (m *PairValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PairValue) GetValue() *StructuredValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// Represents Python's namedtuple.
type NamedTupleValue struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values               []*PairValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NamedTupleValue) Reset()         { *m = NamedTupleValue{} }
func (m *NamedTupleValue) String() string { return proto.CompactTextString(m) }
func (*NamedTupleValue) ProtoMessage()    {}
func (*NamedTupleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{6}
}

func (m *NamedTupleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedTupleValue.Unmarshal(m, b)
}
func (m *NamedTupleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedTupleValue.Marshal(b, m, deterministic)
}
func (m *NamedTupleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedTupleValue.Merge(m, src)
}
func (m *NamedTupleValue) XXX_Size() int {
	return xxx_messageInfo_NamedTupleValue.Size(m)
}
func (m *NamedTupleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedTupleValue.DiscardUnknown(m)
}

var xxx_messageInfo_NamedTupleValue proto.InternalMessageInfo

func (m *NamedTupleValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedTupleValue) GetValues() []*PairValue {
	if m != nil {
		return m.Values
	}
	return nil
}

// A protobuf to tf.TensorSpec.
type TensorSpecProto struct {
	Name                 string                                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shape                *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype                types_go_proto.DataType                 `protobuf:"varint,3,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *TensorSpecProto) Reset()         { *m = TensorSpecProto{} }
func (m *TensorSpecProto) String() string { return proto.CompactTextString(m) }
func (*TensorSpecProto) ProtoMessage()    {}
func (*TensorSpecProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{7}
}

func (m *TensorSpecProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorSpecProto.Unmarshal(m, b)
}
func (m *TensorSpecProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorSpecProto.Marshal(b, m, deterministic)
}
func (m *TensorSpecProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorSpecProto.Merge(m, src)
}
func (m *TensorSpecProto) XXX_Size() int {
	return xxx_messageInfo_TensorSpecProto.Size(m)
}
func (m *TensorSpecProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorSpecProto.DiscardUnknown(m)
}

var xxx_messageInfo_TensorSpecProto proto.InternalMessageInfo

func (m *TensorSpecProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TensorSpecProto) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *TensorSpecProto) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

// Represents a tf.TypeSpec
type TypeSpecProto struct {
	TypeSpecClass TypeSpecProto_TypeSpecClass `protobuf:"varint,1,opt,name=type_spec_class,json=typeSpecClass,proto3,enum=tensorflow.TypeSpecProto_TypeSpecClass" json:"type_spec_class,omitempty"`
	// The value returned by TypeSpec._serialize().
	TypeState *StructuredValue `protobuf:"bytes,2,opt,name=type_state,json=typeState,proto3" json:"type_state,omitempty"`
	// This is currently redundant with the type_spec_class enum, and is only
	// used for error reporting.  In particular, if you use an older binary to
	// load a newer model, and the model uses a TypeSpecClass that the older
	// binary doesn't support, then this lets us display a useful error message.
	TypeSpecClassName    string   `protobuf:"bytes,3,opt,name=type_spec_class_name,json=typeSpecClassName,proto3" json:"type_spec_class_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeSpecProto) Reset()         { *m = TypeSpecProto{} }
func (m *TypeSpecProto) String() string { return proto.CompactTextString(m) }
func (*TypeSpecProto) ProtoMessage()    {}
func (*TypeSpecProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6f8fd91d5fa722, []int{8}
}

func (m *TypeSpecProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeSpecProto.Unmarshal(m, b)
}
func (m *TypeSpecProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeSpecProto.Marshal(b, m, deterministic)
}
func (m *TypeSpecProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeSpecProto.Merge(m, src)
}
func (m *TypeSpecProto) XXX_Size() int {
	return xxx_messageInfo_TypeSpecProto.Size(m)
}
func (m *TypeSpecProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeSpecProto.DiscardUnknown(m)
}

var xxx_messageInfo_TypeSpecProto proto.InternalMessageInfo

func (m *TypeSpecProto) GetTypeSpecClass() TypeSpecProto_TypeSpecClass {
	if m != nil {
		return m.TypeSpecClass
	}
	return TypeSpecProto_UNKNOWN
}

func (m *TypeSpecProto) GetTypeState() *StructuredValue {
	if m != nil {
		return m.TypeState
	}
	return nil
}

func (m *TypeSpecProto) GetTypeSpecClassName() string {
	if m != nil {
		return m.TypeSpecClassName
	}
	return ""
}

func init() {
	proto.RegisterEnum("tensorflow.TypeSpecProto_TypeSpecClass", TypeSpecProto_TypeSpecClass_name, TypeSpecProto_TypeSpecClass_value)
	proto.RegisterType((*StructuredValue)(nil), "tensorflow.StructuredValue")
	proto.RegisterType((*NoneValue)(nil), "tensorflow.NoneValue")
	proto.RegisterType((*ListValue)(nil), "tensorflow.ListValue")
	proto.RegisterType((*TupleValue)(nil), "tensorflow.TupleValue")
	proto.RegisterType((*DictValue)(nil), "tensorflow.DictValue")
	proto.RegisterMapType((map[string]*StructuredValue)(nil), "tensorflow.DictValue.FieldsEntry")
	proto.RegisterType((*PairValue)(nil), "tensorflow.PairValue")
	proto.RegisterType((*NamedTupleValue)(nil), "tensorflow.NamedTupleValue")
	proto.RegisterType((*TensorSpecProto)(nil), "tensorflow.TensorSpecProto")
	proto.RegisterType((*TypeSpecProto)(nil), "tensorflow.TypeSpecProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/struct.proto", fileDescriptor_8f6f8fd91d5fa722)
}

var fileDescriptor_8f6f8fd91d5fa722 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x9b, 0x26, 0xbb, 0x3e, 0x69, 0x9a, 0x74, 0xe8, 0x2e, 0xa5, 0x20, 0x6d, 0x6a, 0x54,
	0x11, 0x21, 0x48, 0x44, 0xbb, 0x54, 0xec, 0x5e, 0xe1, 0xc6, 0x66, 0x13, 0x88, 0x1c, 0x6b, 0x6c,
	0xca, 0xcf, 0x8d, 0xe5, 0x3a, 0x93, 0xae, 0x55, 0xd7, 0xb6, 0xec, 0x09, 0xab, 0x3c, 0x00, 0xaf,
	0xc1, 0x6b, 0xf0, 0x48, 0x3c, 0x04, 0x37, 0x68, 0x66, 0xfc, 0x9b, 0x76, 0x2b, 0x84, 0xb8, 0x49,
	0xe6, 0x7c, 0xfe, 0xce, 0x77, 0xce, 0x7c, 0xe3, 0xe3, 0x81, 0x53, 0x4a, 0xc2, 0x34, 0x4a, 0x56,
	0x41, 0xf4, 0x6e, 0xec, 0x45, 0x09, 0x19, 0xc7, 0x49, 0x44, 0xa3, 0xeb, 0xf5, 0x6a, 0x9c, 0xd2,
	0x64, 0xed, 0xd1, 0x11, 0x8f, 0x11, 0x94, 0xb4, 0xe3, 0x2f, 0xb6, 0x53, 0x56, 0x89, 0x7b, 0x47,
	0xde, 0x45, 0xc9, 0xed, 0x58, 0x3c, 0x71, 0xd2, 0xb7, 0x6e, 0x4c, 0x44, 0xe6, 0xf1, 0xe9, 0x23,
	0xec, 0x4d, 0x4c, 0x52, 0x41, 0x53, 0xfe, 0x6e, 0x41, 0xcf, 0xe2, 0x15, 0xd7, 0x09, 0x59, 0x5e,
	0xb9, 0xc1, 0x9a, 0xa0, 0x0b, 0x80, 0x30, 0x0a, 0x89, 0xf3, 0x1b, 0x8b, 0x8e, 0xa4, 0x81, 0x34,
	0xec, 0x9c, 0x3d, 0x1b, 0x95, 0x7a, 0x23, 0x23, 0x0a, 0x09, 0xa7, 0x4e, 0x1b, 0x58, 0x0e, 0xf3,
	0x00, 0x9d, 0x42, 0x77, 0x15, 0x44, 0x2e, 0xbd, 0x78, 0x99, 0xa5, 0x76, 0x06, 0xd2, 0x50, 0x9a,
	0x36, 0xf0, 0x5e, 0x06, 0x0b, 0xda, 0x09, 0x74, 0xfc, 0xb0, 0x24, 0xed, 0x0d, 0xa4, 0x21, 0x9a,
	0x36, 0x30, 0x70, 0x50, 0x50, 0x3e, 0x85, 0xbd, 0x94, 0x26, 0x7e, 0x78, 0x93, 0x71, 0xba, 0x03,
	0x69, 0x28, 0x4f, 0x1b, 0xb8, 0x23, 0x50, 0x41, 0x7a, 0x01, 0x70, 0x1d, 0x45, 0x41, 0x46, 0xd9,
	0x1f, 0x48, 0xc3, 0xa7, 0xac, 0x1f, 0x86, 0x09, 0xc2, 0x1c, 0x50, 0xd5, 0x98, 0x8c, 0xf8, 0x82,
	0xef, 0xe7, 0x93, 0xea, 0x7e, 0x6c, 0xbe, 0xb4, 0x18, 0xc9, 0x64, 0xae, 0x4c, 0x1b, 0xb8, 0x4f,
	0x4b, 0x4c, 0xa8, 0x69, 0x85, 0xda, 0x92, 0x19, 0x98, 0xa9, 0x0d, 0x06, 0xd2, 0x70, 0xff, 0xec,
	0xb0, 0xaa, 0xa6, 0xb9, 0xd4, 0xb5, 0x37, 0x31, 0x29, 0x55, 0x34, 0x96, 0x20, 0x54, 0x66, 0x70,
	0x90, 0xf7, 0x14, 0x13, 0x2f, 0x13, 0x39, 0xe1, 0x2d, 0x7d, 0xfc, 0x40, 0x4b, 0x31, 0xf1, 0xf2,
	0x8e, 0x7a, 0xb4, 0x80, 0x84, 0xd4, 0x04, 0x7a, 0xbc, 0x91, 0x8a, 0x90, 0xc2, 0x85, 0x3e, 0xaa,
	0x09, 0x6d, 0x62, 0x52, 0x95, 0xe9, 0xd2, 0x0c, 0x28, 0xce, 0x3a, 0xf0, 0x53, 0x9a, 0xe5, 0x9f,
	0xdf, 0x3f, 0xeb, 0xb9, 0x9f, 0xd2, 0xe2, 0xac, 0x83, 0x3c, 0x40, 0xaf, 0xa0, 0x43, 0xd7, 0x71,
	0x90, 0xdb, 0xf0, 0x92, 0x27, 0x3e, 0xaf, 0x15, 0x66, 0x8f, 0xf3, 0x4c, 0xa0, 0x45, 0xc4, 0x4a,
	0x2e, 0x7d, 0x2f, 0x2f, 0xf9, 0xf5, 0xfd, 0x92, 0x9a, 0xef, 0x95, 0x25, 0x97, 0x79, 0xc0, 0xac,
	0x0b, 0xdd, 0x3b, 0xb2, 0x74, 0xaa, 0x85, 0x2f, 0xee, 0x5b, 0x67, 0x30, 0x52, 0xad, 0x7a, 0x2f,
	0xac, 0x43, 0x97, 0x6d, 0xd8, 0xbd, 0xf5, 0xc3, 0xa5, 0xd2, 0x01, 0xb9, 0x78, 0x97, 0x95, 0x6f,
	0x41, 0x2e, 0x36, 0x8b, 0xce, 0xa1, 0xcd, 0x0b, 0xa4, 0x47, 0xd2, 0xa0, 0xb9, 0x5d, 0x61, 0x6b,
	0x60, 0x70, 0x46, 0x55, 0x54, 0x80, 0xb2, 0xc8, 0x7f, 0x93, 0xf8, 0x43, 0x02, 0xb9, 0xd8, 0x3f,
	0x7a, 0x05, 0xed, 0x95, 0x4f, 0x82, 0x65, 0x2e, 0x71, 0xf2, 0xa0, 0x4d, 0xa3, 0xef, 0x38, 0x47,
	0x0f, 0x69, 0xb2, 0xc1, 0x59, 0xc2, 0xf1, 0x15, 0x74, 0x2a, 0x30, 0xea, 0x43, 0xf3, 0x96, 0x6c,
	0xf8, 0x30, 0xcb, 0x98, 0x2d, 0xd1, 0x57, 0xd0, 0x12, 0x16, 0xee, 0xdc, 0xb7, 0x70, 0xbb, 0x3b,
	0xc1, 0x7c, 0xbd, 0xf3, 0x8d, 0xa4, 0x98, 0x20, 0x9b, 0xae, 0x9f, 0x88, 0xfe, 0xfe, 0x0f, 0x55,
	0xc5, 0x86, 0xde, 0xd6, 0x91, 0x21, 0x04, 0xbb, 0xec, 0xc8, 0x32, 0x61, 0xbe, 0x46, 0x5f, 0x16,
	0x76, 0xee, 0x70, 0x2f, 0x6a, 0xaf, 0x4c, 0xd1, 0x52, 0x61, 0xe4, 0xef, 0x12, 0xf4, 0xb6, 0x86,
	0xe8, 0x41, 0xd9, 0x33, 0x68, 0xf1, 0xaf, 0x43, 0xd6, 0xf0, 0xa3, 0xdf, 0x05, 0x2c, 0xa8, 0xe8,
	0x73, 0x68, 0xf1, 0x6f, 0xc0, 0x51, 0xf3, 0xfd, 0xd3, 0x8f, 0x05, 0x45, 0xf9, 0xb3, 0x09, 0xdd,
	0xda, 0x0c, 0xa2, 0x45, 0x75, 0x6e, 0xbd, 0xc0, 0x4d, 0x53, 0xde, 0xd0, 0xfe, 0xd9, 0x67, 0xef,
	0x9d, 0xdb, 0x22, 0x9a, 0x30, 0x7a, 0x39, 0xc3, 0x3c, 0x44, 0xaf, 0x01, 0x84, 0x20, 0x75, 0xe9,
	0xbf, 0x32, 0x5e, 0xe6, 0xf9, 0x8c, 0x8d, 0xc6, 0x70, 0xb8, 0xd5, 0x8c, 0xc3, 0x2d, 0x6a, 0x72,
	0x8b, 0x0e, 0x6a, 0x85, 0xd8, 0x29, 0x29, 0x7f, 0x49, 0xe5, 0x7e, 0x44, 0xf9, 0x0e, 0x3c, 0xf9,
	0xd1, 0xf8, 0xc1, 0x58, 0xfc, 0x64, 0xf4, 0x1b, 0xe8, 0x39, 0x20, 0xcb, 0x54, 0xb1, 0xa5, 0x3b,
	0xb6, 0x6e, 0x58, 0x0b, 0xec, 0x58, 0xa6, 0x3e, 0xe9, 0x4b, 0xe8, 0x43, 0xf8, 0x60, 0x66, 0x68,
	0xfa, 0xcf, 0xba, 0xe6, 0x58, 0xf3, 0xd9, 0x44, 0xb7, 0xc4, 0x83, 0x1d, 0x96, 0x80, 0xd5, 0x37,
	0x6f, 0x74, 0xad, 0x96, 0xd0, 0x44, 0xcf, 0xe0, 0x20, 0x03, 0x54, 0x8c, 0xd5, 0x5f, 0x04, 0xbc,
	0xcb, 0x60, 0x4d, 0xb5, 0x55, 0x87, 0xfd, 0x58, 0xba, 0x2d, 0xe0, 0x16, 0x53, 0xe1, 0xf0, 0xcc,
	0xd6, 0xb1, 0x6a, 0xe7, 0x2a, 0x6d, 0x74, 0x00, 0xdd, 0x85, 0x69, 0xcf, 0x16, 0x86, 0x3a, 0x17,
	0xd0, 0x13, 0x74, 0x08, 0x7d, 0x53, 0xc7, 0x0e, 0xd6, 0xcd, 0xf9, 0x6c, 0xa2, 0x0a, 0xf4, 0x29,
	0x23, 0x5e, 0xa9, 0x78, 0xa6, 0x5e, 0xce, 0x75, 0x01, 0xc9, 0x97, 0xdf, 0xff, 0x3a, 0xbd, 0xf1,
	0xe9, 0xdb, 0xf5, 0xf5, 0xc8, 0x8b, 0xee, 0xc6, 0x95, 0xeb, 0xf4, 0xe1, 0xe5, 0x4d, 0x24, 0xee,
	0x59, 0xf6, 0xe3, 0xf0, 0xcb, 0x35, 0x75, 0x6e, 0x22, 0xb1, 0xba, 0x6e, 0xf3, 0xbf, 0xf3, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xda, 0xf9, 0xf8, 0x72, 0xf7, 0x07, 0x00, 0x00,
}

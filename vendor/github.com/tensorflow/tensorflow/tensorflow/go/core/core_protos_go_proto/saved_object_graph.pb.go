// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/saved_object_graph.proto

package core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	variable_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/variable_go_proto"
	versions_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto"
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

type SavedObjectGraph struct {
	// Flattened list of objects in the object graph.
	//
	// The position of the object in this list indicates its id.
	// Nodes[0] is considered the root node.
	Nodes []*SavedObject `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// Information about captures and output structures in concrete functions.
	// Referenced from SavedBareConcreteFunction and SavedFunction.
	ConcreteFunctions    map[string]*SavedConcreteFunction `protobuf:"bytes,2,rep,name=concrete_functions,json=concreteFunctions,proto3" json:"concrete_functions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SavedObjectGraph) Reset()         { *m = SavedObjectGraph{} }
func (m *SavedObjectGraph) String() string { return proto.CompactTextString(m) }
func (*SavedObjectGraph) ProtoMessage()    {}
func (*SavedObjectGraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{0}
}

func (m *SavedObjectGraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedObjectGraph.Unmarshal(m, b)
}
func (m *SavedObjectGraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedObjectGraph.Marshal(b, m, deterministic)
}
func (m *SavedObjectGraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedObjectGraph.Merge(m, src)
}
func (m *SavedObjectGraph) XXX_Size() int {
	return xxx_messageInfo_SavedObjectGraph.Size(m)
}
func (m *SavedObjectGraph) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedObjectGraph.DiscardUnknown(m)
}

var xxx_messageInfo_SavedObjectGraph proto.InternalMessageInfo

func (m *SavedObjectGraph) GetNodes() []*SavedObject {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *SavedObjectGraph) GetConcreteFunctions() map[string]*SavedConcreteFunction {
	if m != nil {
		return m.ConcreteFunctions
	}
	return nil
}

type SavedObject struct {
	// Objects which this object depends on: named edges in the dependency
	// graph.
	//
	// Note: currently only valid if kind == "user_object".
	Children []*TrackableObjectGraph_TrackableObject_ObjectReference `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
	// Slot variables owned by this object. This describes the three-way
	// (optimizer, variable, slot variable) relationship; none of the three
	// depend on the others directly.
	//
	// Note: currently only valid if kind == "user_object".
	SlotVariables []*TrackableObjectGraph_TrackableObject_SlotVariableReference `protobuf:"bytes,3,rep,name=slot_variables,json=slotVariables,proto3" json:"slot_variables,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*SavedObject_UserObject
	//	*SavedObject_Asset
	//	*SavedObject_Function
	//	*SavedObject_Variable
	//	*SavedObject_BareConcreteFunction
	//	*SavedObject_Constant
	//	*SavedObject_Resource
	Kind                 isSavedObject_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SavedObject) Reset()         { *m = SavedObject{} }
func (m *SavedObject) String() string { return proto.CompactTextString(m) }
func (*SavedObject) ProtoMessage()    {}
func (*SavedObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{1}
}

func (m *SavedObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedObject.Unmarshal(m, b)
}
func (m *SavedObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedObject.Marshal(b, m, deterministic)
}
func (m *SavedObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedObject.Merge(m, src)
}
func (m *SavedObject) XXX_Size() int {
	return xxx_messageInfo_SavedObject.Size(m)
}
func (m *SavedObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedObject.DiscardUnknown(m)
}

var xxx_messageInfo_SavedObject proto.InternalMessageInfo

func (m *SavedObject) GetChildren() []*TrackableObjectGraph_TrackableObject_ObjectReference {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *SavedObject) GetSlotVariables() []*TrackableObjectGraph_TrackableObject_SlotVariableReference {
	if m != nil {
		return m.SlotVariables
	}
	return nil
}

type isSavedObject_Kind interface {
	isSavedObject_Kind()
}

type SavedObject_UserObject struct {
	UserObject *SavedUserObject `protobuf:"bytes,4,opt,name=user_object,json=userObject,proto3,oneof"`
}

type SavedObject_Asset struct {
	Asset *SavedAsset `protobuf:"bytes,5,opt,name=asset,proto3,oneof"`
}

type SavedObject_Function struct {
	Function *SavedFunction `protobuf:"bytes,6,opt,name=function,proto3,oneof"`
}

type SavedObject_Variable struct {
	Variable *SavedVariable `protobuf:"bytes,7,opt,name=variable,proto3,oneof"`
}

type SavedObject_BareConcreteFunction struct {
	BareConcreteFunction *SavedBareConcreteFunction `protobuf:"bytes,8,opt,name=bare_concrete_function,json=bareConcreteFunction,proto3,oneof"`
}

type SavedObject_Constant struct {
	Constant *SavedConstant `protobuf:"bytes,9,opt,name=constant,proto3,oneof"`
}

type SavedObject_Resource struct {
	Resource *SavedResource `protobuf:"bytes,10,opt,name=resource,proto3,oneof"`
}

func (*SavedObject_UserObject) isSavedObject_Kind() {}

func (*SavedObject_Asset) isSavedObject_Kind() {}

func (*SavedObject_Function) isSavedObject_Kind() {}

func (*SavedObject_Variable) isSavedObject_Kind() {}

func (*SavedObject_BareConcreteFunction) isSavedObject_Kind() {}

func (*SavedObject_Constant) isSavedObject_Kind() {}

func (*SavedObject_Resource) isSavedObject_Kind() {}

func (m *SavedObject) GetKind() isSavedObject_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *SavedObject) GetUserObject() *SavedUserObject {
	if x, ok := m.GetKind().(*SavedObject_UserObject); ok {
		return x.UserObject
	}
	return nil
}

func (m *SavedObject) GetAsset() *SavedAsset {
	if x, ok := m.GetKind().(*SavedObject_Asset); ok {
		return x.Asset
	}
	return nil
}

func (m *SavedObject) GetFunction() *SavedFunction {
	if x, ok := m.GetKind().(*SavedObject_Function); ok {
		return x.Function
	}
	return nil
}

func (m *SavedObject) GetVariable() *SavedVariable {
	if x, ok := m.GetKind().(*SavedObject_Variable); ok {
		return x.Variable
	}
	return nil
}

func (m *SavedObject) GetBareConcreteFunction() *SavedBareConcreteFunction {
	if x, ok := m.GetKind().(*SavedObject_BareConcreteFunction); ok {
		return x.BareConcreteFunction
	}
	return nil
}

func (m *SavedObject) GetConstant() *SavedConstant {
	if x, ok := m.GetKind().(*SavedObject_Constant); ok {
		return x.Constant
	}
	return nil
}

func (m *SavedObject) GetResource() *SavedResource {
	if x, ok := m.GetKind().(*SavedObject_Resource); ok {
		return x.Resource
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SavedObject) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SavedObject_UserObject)(nil),
		(*SavedObject_Asset)(nil),
		(*SavedObject_Function)(nil),
		(*SavedObject_Variable)(nil),
		(*SavedObject_BareConcreteFunction)(nil),
		(*SavedObject_Constant)(nil),
		(*SavedObject_Resource)(nil),
	}
}

// A SavedUserObject is an object (in the object-oriented language of the
// TensorFlow program) of some user- or framework-defined class other than
// those handled specifically by the other kinds of SavedObjects.
//
// This object cannot be evaluated as a tensor, and therefore cannot be bound
// to an input of a function.
type SavedUserObject struct {
	// Corresponds to a registration of the type to use in the loading program.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Version information from the producer of this SavedUserObject.
	Version *versions_go_proto.VersionDef `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Initialization-related metadata.
	Metadata             string   `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedUserObject) Reset()         { *m = SavedUserObject{} }
func (m *SavedUserObject) String() string { return proto.CompactTextString(m) }
func (*SavedUserObject) ProtoMessage()    {}
func (*SavedUserObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{2}
}

func (m *SavedUserObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedUserObject.Unmarshal(m, b)
}
func (m *SavedUserObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedUserObject.Marshal(b, m, deterministic)
}
func (m *SavedUserObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedUserObject.Merge(m, src)
}
func (m *SavedUserObject) XXX_Size() int {
	return xxx_messageInfo_SavedUserObject.Size(m)
}
func (m *SavedUserObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedUserObject.DiscardUnknown(m)
}

var xxx_messageInfo_SavedUserObject proto.InternalMessageInfo

func (m *SavedUserObject) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *SavedUserObject) GetVersion() *versions_go_proto.VersionDef {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *SavedUserObject) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// A SavedAsset points to an asset in the MetaGraph.
//
// When bound to a function this object evaluates to a tensor with the absolute
// filename. Users should not depend on a particular part of the filename to
// remain stable (e.g. basename could be changed).
type SavedAsset struct {
	// Index into `MetaGraphDef.asset_file_def[]` that describes the Asset.
	//
	// Only the field `AssetFileDef.filename` is used. Other fields, such as
	// `AssetFileDef.tensor_info`, MUST be ignored.
	AssetFileDefIndex    int32    `protobuf:"varint,1,opt,name=asset_file_def_index,json=assetFileDefIndex,proto3" json:"asset_file_def_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedAsset) Reset()         { *m = SavedAsset{} }
func (m *SavedAsset) String() string { return proto.CompactTextString(m) }
func (*SavedAsset) ProtoMessage()    {}
func (*SavedAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{3}
}

func (m *SavedAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedAsset.Unmarshal(m, b)
}
func (m *SavedAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedAsset.Marshal(b, m, deterministic)
}
func (m *SavedAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedAsset.Merge(m, src)
}
func (m *SavedAsset) XXX_Size() int {
	return xxx_messageInfo_SavedAsset.Size(m)
}
func (m *SavedAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SavedAsset proto.InternalMessageInfo

func (m *SavedAsset) GetAssetFileDefIndex() int32 {
	if m != nil {
		return m.AssetFileDefIndex
	}
	return 0
}

// A function with multiple signatures, possibly with non-Tensor arguments.
type SavedFunction struct {
	ConcreteFunctions    []string      `protobuf:"bytes,1,rep,name=concrete_functions,json=concreteFunctions,proto3" json:"concrete_functions,omitempty"`
	FunctionSpec         *FunctionSpec `protobuf:"bytes,2,opt,name=function_spec,json=functionSpec,proto3" json:"function_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SavedFunction) Reset()         { *m = SavedFunction{} }
func (m *SavedFunction) String() string { return proto.CompactTextString(m) }
func (*SavedFunction) ProtoMessage()    {}
func (*SavedFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{4}
}

func (m *SavedFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedFunction.Unmarshal(m, b)
}
func (m *SavedFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedFunction.Marshal(b, m, deterministic)
}
func (m *SavedFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedFunction.Merge(m, src)
}
func (m *SavedFunction) XXX_Size() int {
	return xxx_messageInfo_SavedFunction.Size(m)
}
func (m *SavedFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedFunction.DiscardUnknown(m)
}

var xxx_messageInfo_SavedFunction proto.InternalMessageInfo

func (m *SavedFunction) GetConcreteFunctions() []string {
	if m != nil {
		return m.ConcreteFunctions
	}
	return nil
}

func (m *SavedFunction) GetFunctionSpec() *FunctionSpec {
	if m != nil {
		return m.FunctionSpec
	}
	return nil
}

// Stores low-level information about a concrete function. Referenced in either
// a SavedFunction or a SavedBareConcreteFunction.
type SavedConcreteFunction struct {
	// Bound inputs to the function. The SavedObjects identified by the node ids
	// given here are appended as extra inputs to the caller-supplied inputs.
	// The only types of SavedObjects valid here are SavedVariable, SavedResource
	// and SavedAsset.
	BoundInputs []int32 `protobuf:"varint,2,rep,packed,name=bound_inputs,json=boundInputs,proto3" json:"bound_inputs,omitempty"`
	// Input in canonicalized form that was received to create this concrete
	// function.
	CanonicalizedInputSignature *StructuredValue `protobuf:"bytes,3,opt,name=canonicalized_input_signature,json=canonicalizedInputSignature,proto3" json:"canonicalized_input_signature,omitempty"`
	// Output that was the return value of this function after replacing all
	// Tensors with TensorSpecs. This can be an arbitrary nested function and will
	// be used to reconstruct the full structure from pure tensors.
	OutputSignature      *StructuredValue `protobuf:"bytes,4,opt,name=output_signature,json=outputSignature,proto3" json:"output_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SavedConcreteFunction) Reset()         { *m = SavedConcreteFunction{} }
func (m *SavedConcreteFunction) String() string { return proto.CompactTextString(m) }
func (*SavedConcreteFunction) ProtoMessage()    {}
func (*SavedConcreteFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{5}
}

func (m *SavedConcreteFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedConcreteFunction.Unmarshal(m, b)
}
func (m *SavedConcreteFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedConcreteFunction.Marshal(b, m, deterministic)
}
func (m *SavedConcreteFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedConcreteFunction.Merge(m, src)
}
func (m *SavedConcreteFunction) XXX_Size() int {
	return xxx_messageInfo_SavedConcreteFunction.Size(m)
}
func (m *SavedConcreteFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedConcreteFunction.DiscardUnknown(m)
}

var xxx_messageInfo_SavedConcreteFunction proto.InternalMessageInfo

func (m *SavedConcreteFunction) GetBoundInputs() []int32 {
	if m != nil {
		return m.BoundInputs
	}
	return nil
}

func (m *SavedConcreteFunction) GetCanonicalizedInputSignature() *StructuredValue {
	if m != nil {
		return m.CanonicalizedInputSignature
	}
	return nil
}

func (m *SavedConcreteFunction) GetOutputSignature() *StructuredValue {
	if m != nil {
		return m.OutputSignature
	}
	return nil
}

type SavedBareConcreteFunction struct {
	// Identifies a SavedConcreteFunction.
	ConcreteFunctionName string `protobuf:"bytes,1,opt,name=concrete_function_name,json=concreteFunctionName,proto3" json:"concrete_function_name,omitempty"`
	// A sequence of unique strings, one per Tensor argument.
	ArgumentKeywords []string `protobuf:"bytes,2,rep,name=argument_keywords,json=argumentKeywords,proto3" json:"argument_keywords,omitempty"`
	// The prefix of `argument_keywords` which may be identified by position.
	AllowedPositionalArguments int64    `protobuf:"varint,3,opt,name=allowed_positional_arguments,json=allowedPositionalArguments,proto3" json:"allowed_positional_arguments,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *SavedBareConcreteFunction) Reset()         { *m = SavedBareConcreteFunction{} }
func (m *SavedBareConcreteFunction) String() string { return proto.CompactTextString(m) }
func (*SavedBareConcreteFunction) ProtoMessage()    {}
func (*SavedBareConcreteFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{6}
}

func (m *SavedBareConcreteFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedBareConcreteFunction.Unmarshal(m, b)
}
func (m *SavedBareConcreteFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedBareConcreteFunction.Marshal(b, m, deterministic)
}
func (m *SavedBareConcreteFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedBareConcreteFunction.Merge(m, src)
}
func (m *SavedBareConcreteFunction) XXX_Size() int {
	return xxx_messageInfo_SavedBareConcreteFunction.Size(m)
}
func (m *SavedBareConcreteFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedBareConcreteFunction.DiscardUnknown(m)
}

var xxx_messageInfo_SavedBareConcreteFunction proto.InternalMessageInfo

func (m *SavedBareConcreteFunction) GetConcreteFunctionName() string {
	if m != nil {
		return m.ConcreteFunctionName
	}
	return ""
}

func (m *SavedBareConcreteFunction) GetArgumentKeywords() []string {
	if m != nil {
		return m.ArgumentKeywords
	}
	return nil
}

func (m *SavedBareConcreteFunction) GetAllowedPositionalArguments() int64 {
	if m != nil {
		return m.AllowedPositionalArguments
	}
	return 0
}

type SavedConstant struct {
	// An Operation name for a ConstantOp in this SavedObjectGraph's MetaGraph.
	Operation            string   `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedConstant) Reset()         { *m = SavedConstant{} }
func (m *SavedConstant) String() string { return proto.CompactTextString(m) }
func (*SavedConstant) ProtoMessage()    {}
func (*SavedConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{7}
}

func (m *SavedConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedConstant.Unmarshal(m, b)
}
func (m *SavedConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedConstant.Marshal(b, m, deterministic)
}
func (m *SavedConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedConstant.Merge(m, src)
}
func (m *SavedConstant) XXX_Size() int {
	return xxx_messageInfo_SavedConstant.Size(m)
}
func (m *SavedConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedConstant.DiscardUnknown(m)
}

var xxx_messageInfo_SavedConstant proto.InternalMessageInfo

func (m *SavedConstant) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

// Represents a Variable that is initialized by loading the contents from the
// checkpoint.
type SavedVariable struct {
	Dtype                types_go_proto.DataType                   `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *tensor_shape_go_proto.TensorShapeProto   `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	Trainable            bool                                      `protobuf:"varint,3,opt,name=trainable,proto3" json:"trainable,omitempty"`
	Synchronization      variable_go_proto.VariableSynchronization `protobuf:"varint,4,opt,name=synchronization,proto3,enum=tensorflow.VariableSynchronization" json:"synchronization,omitempty"`
	Aggregation          variable_go_proto.VariableAggregation     `protobuf:"varint,5,opt,name=aggregation,proto3,enum=tensorflow.VariableAggregation" json:"aggregation,omitempty"`
	Name                 string                                    `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *SavedVariable) Reset()         { *m = SavedVariable{} }
func (m *SavedVariable) String() string { return proto.CompactTextString(m) }
func (*SavedVariable) ProtoMessage()    {}
func (*SavedVariable) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{8}
}

func (m *SavedVariable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedVariable.Unmarshal(m, b)
}
func (m *SavedVariable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedVariable.Marshal(b, m, deterministic)
}
func (m *SavedVariable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedVariable.Merge(m, src)
}
func (m *SavedVariable) XXX_Size() int {
	return xxx_messageInfo_SavedVariable.Size(m)
}
func (m *SavedVariable) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedVariable.DiscardUnknown(m)
}

var xxx_messageInfo_SavedVariable proto.InternalMessageInfo

func (m *SavedVariable) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *SavedVariable) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *SavedVariable) GetTrainable() bool {
	if m != nil {
		return m.Trainable
	}
	return false
}

func (m *SavedVariable) GetSynchronization() variable_go_proto.VariableSynchronization {
	if m != nil {
		return m.Synchronization
	}
	return variable_go_proto.VariableSynchronization_VARIABLE_SYNCHRONIZATION_AUTO
}

func (m *SavedVariable) GetAggregation() variable_go_proto.VariableAggregation {
	if m != nil {
		return m.Aggregation
	}
	return variable_go_proto.VariableAggregation_VARIABLE_AGGREGATION_NONE
}

func (m *SavedVariable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Represents `FunctionSpec` used in `Function`. This represents a
// function that has been wrapped as a TensorFlow `Function`.
type FunctionSpec struct {
	// Full arg spec from inspect.getfullargspec().
	Fullargspec *StructuredValue `protobuf:"bytes,1,opt,name=fullargspec,proto3" json:"fullargspec,omitempty"`
	// Whether this represents a class method.
	IsMethod bool `protobuf:"varint,2,opt,name=is_method,json=isMethod,proto3" json:"is_method,omitempty"`
	// The input signature, if specified.
	InputSignature       *StructuredValue `protobuf:"bytes,5,opt,name=input_signature,json=inputSignature,proto3" json:"input_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FunctionSpec) Reset()         { *m = FunctionSpec{} }
func (m *FunctionSpec) String() string { return proto.CompactTextString(m) }
func (*FunctionSpec) ProtoMessage()    {}
func (*FunctionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{9}
}

func (m *FunctionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionSpec.Unmarshal(m, b)
}
func (m *FunctionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionSpec.Marshal(b, m, deterministic)
}
func (m *FunctionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionSpec.Merge(m, src)
}
func (m *FunctionSpec) XXX_Size() int {
	return xxx_messageInfo_FunctionSpec.Size(m)
}
func (m *FunctionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionSpec proto.InternalMessageInfo

func (m *FunctionSpec) GetFullargspec() *StructuredValue {
	if m != nil {
		return m.Fullargspec
	}
	return nil
}

func (m *FunctionSpec) GetIsMethod() bool {
	if m != nil {
		return m.IsMethod
	}
	return false
}

func (m *FunctionSpec) GetInputSignature() *StructuredValue {
	if m != nil {
		return m.InputSignature
	}
	return nil
}

// A SavedResource represents a TF object that holds state during its lifetime.
// An object of this type can have a reference to a:
// create_resource() and an initialize() function.
type SavedResource struct {
	// A device specification indicating a required placement for the resource
	// creation function, e.g. "CPU". An empty string allows the user to select a
	// device.
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedResource) Reset()         { *m = SavedResource{} }
func (m *SavedResource) String() string { return proto.CompactTextString(m) }
func (*SavedResource) ProtoMessage()    {}
func (*SavedResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f63c49021beb5aa, []int{10}
}

func (m *SavedResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedResource.Unmarshal(m, b)
}
func (m *SavedResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedResource.Marshal(b, m, deterministic)
}
func (m *SavedResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedResource.Merge(m, src)
}
func (m *SavedResource) XXX_Size() int {
	return xxx_messageInfo_SavedResource.Size(m)
}
func (m *SavedResource) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedResource.DiscardUnknown(m)
}

var xxx_messageInfo_SavedResource proto.InternalMessageInfo

func (m *SavedResource) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func init() {
	proto.RegisterType((*SavedObjectGraph)(nil), "tensorflow.SavedObjectGraph")
	proto.RegisterMapType((map[string]*SavedConcreteFunction)(nil), "tensorflow.SavedObjectGraph.ConcreteFunctionsEntry")
	proto.RegisterType((*SavedObject)(nil), "tensorflow.SavedObject")
	proto.RegisterType((*SavedUserObject)(nil), "tensorflow.SavedUserObject")
	proto.RegisterType((*SavedAsset)(nil), "tensorflow.SavedAsset")
	proto.RegisterType((*SavedFunction)(nil), "tensorflow.SavedFunction")
	proto.RegisterType((*SavedConcreteFunction)(nil), "tensorflow.SavedConcreteFunction")
	proto.RegisterType((*SavedBareConcreteFunction)(nil), "tensorflow.SavedBareConcreteFunction")
	proto.RegisterType((*SavedConstant)(nil), "tensorflow.SavedConstant")
	proto.RegisterType((*SavedVariable)(nil), "tensorflow.SavedVariable")
	proto.RegisterType((*FunctionSpec)(nil), "tensorflow.FunctionSpec")
	proto.RegisterType((*SavedResource)(nil), "tensorflow.SavedResource")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/saved_object_graph.proto", fileDescriptor_4f63c49021beb5aa)
}

var fileDescriptor_4f63c49021beb5aa = []byte{
	// 1091 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0x8e, 0xff, 0x6c, 0x6a, 0x1f, 0xb7, 0x8d, 0x3b, 0xca, 0x2f, 0xbf, 0x6d, 0x1a, 0x20, 0x35,
	0xaa, 0x88, 0x80, 0xd8, 0x90, 0x82, 0x8a, 0x90, 0x82, 0x9a, 0x34, 0x84, 0x34, 0x28, 0x50, 0x8d,
	0x4b, 0x2f, 0x10, 0x68, 0x35, 0xde, 0x3d, 0x6b, 0x2f, 0x59, 0xcf, 0x58, 0x33, 0xb3, 0x09, 0xae,
	0x84, 0x78, 0x03, 0x5e, 0x83, 0x07, 0xe0, 0x96, 0x0b, 0xde, 0x85, 0x97, 0xe0, 0x12, 0xed, 0xec,
	0xac, 0xbd, 0xfe, 0x47, 0xe0, 0xc6, 0x9a, 0x39, 0xe7, 0xfb, 0xbe, 0x39, 0x73, 0xe6, 0x9c, 0xe3,
	0x85, 0x0f, 0x35, 0x72, 0x25, 0x64, 0x18, 0x8b, 0xeb, 0x8e, 0x2f, 0x24, 0x76, 0x46, 0x52, 0x68,
	0xd1, 0x4b, 0xc2, 0x8e, 0x62, 0x57, 0x18, 0x78, 0xa2, 0xf7, 0x03, 0xfa, 0xda, 0xeb, 0x4b, 0x36,
	0x1a, 0xb4, 0x8d, 0x8f, 0xc0, 0x94, 0xb2, 0xfd, 0xfe, 0x3c, 0x3d, 0x94, 0x6c, 0x88, 0xd7, 0x42,
	0x5e, 0x76, 0x32, 0x8f, 0xa7, 0x06, 0x6c, 0x84, 0x19, 0x73, 0xfb, 0xd1, 0x3f, 0xa0, 0xc7, 0x23,
	0x54, 0x16, 0xb6, 0xb7, 0x1a, 0x76, 0xc5, 0x64, 0xc4, 0x7a, 0x31, 0xfe, 0x0b, 0x24, 0x4a, 0x15,
	0x09, 0xae, 0x56, 0x1d, 0x3d, 0xbd, 0xa7, 0x96, 0x89, 0xaf, 0x2d, 0xec, 0xe3, 0x95, 0x30, 0x2d,
	0x99, 0x7f, 0x99, 0x1e, 0xbd, 0x24, 0x25, 0xad, 0x5f, 0xca, 0xd0, 0xec, 0xa6, 0xf9, 0xfa, 0xda,
	0xf8, 0xbe, 0x48, 0x5d, 0x64, 0x1f, 0x1c, 0x2e, 0x02, 0x54, 0x6e, 0x69, 0xb7, 0xb2, 0xd7, 0x38,
	0xf8, 0x7f, 0x7b, 0xaa, 0xdd, 0x2e, 0x80, 0x69, 0x86, 0x22, 0x3d, 0x20, 0xbe, 0xe0, 0xbe, 0x44,
	0x8d, 0x5e, 0x98, 0x70, 0x5f, 0xa7, 0xd1, 0xbb, 0x65, 0xc3, 0x7d, 0xbc, 0x82, 0x6b, 0x0e, 0x6a,
	0x3f, 0xb3, 0xb4, 0xd3, 0x9c, 0xf5, 0x39, 0xd7, 0x72, 0x4c, 0xef, 0xf9, 0xf3, 0xf6, 0xed, 0x3e,
	0x6c, 0x2d, 0x07, 0x93, 0x26, 0x54, 0x2e, 0x71, 0xec, 0x96, 0x76, 0x4b, 0x7b, 0x75, 0x9a, 0x2e,
	0xc9, 0x13, 0x70, 0xae, 0x58, 0x9c, 0xa0, 0x5b, 0xde, 0x2d, 0xed, 0x35, 0x0e, 0x1e, 0x2e, 0x84,
	0x30, 0xaf, 0x44, 0x33, 0xfc, 0xa7, 0xe5, 0x4f, 0x4a, 0xad, 0x5f, 0x1d, 0x68, 0x14, 0xe2, 0x24,
	0xdf, 0x41, 0xcd, 0x1f, 0x44, 0x71, 0x20, 0x91, 0xdb, 0x74, 0x3c, 0x2d, 0xea, 0xbd, 0xcc, 0x93,
	0x5b, 0xbc, 0xd6, 0x9c, 0xb1, 0x6d, 0xd3, 0x85, 0x21, 0x4a, 0xe4, 0x3e, 0xd2, 0x89, 0x22, 0x19,
	0xc2, 0x5d, 0x15, 0x0b, 0xed, 0xe5, 0xd5, 0xa1, 0xdc, 0x8a, 0x39, 0xe3, 0xf4, 0x3f, 0x9f, 0xd1,
	0x8d, 0x85, 0x7e, 0x65, 0x55, 0xa6, 0x27, 0xdd, 0x51, 0x05, 0xb3, 0x22, 0x9f, 0x41, 0x23, 0x51,
	0x28, 0x6d, 0x21, 0xb8, 0x55, 0x93, 0x9f, 0x07, 0x0b, 0xf9, 0xf9, 0x46, 0xa1, 0xcc, 0x64, 0xcf,
	0xd6, 0x28, 0x24, 0x93, 0x1d, 0x69, 0x83, 0xc3, 0x94, 0x42, 0xed, 0x3a, 0x86, 0xb9, 0xb5, 0xc0,
	0x3c, 0x4a, 0xbd, 0x67, 0x6b, 0x34, 0x83, 0x91, 0x27, 0x50, 0xcb, 0x0b, 0xc2, 0x5d, 0x37, 0x94,
	0xfb, 0x0b, 0x94, 0xfc, 0x11, 0xce, 0xd6, 0xe8, 0x04, 0x9c, 0x12, 0xf3, 0x94, 0xb8, 0xb7, 0x56,
	0x10, 0xf3, 0x6b, 0xa5, 0xc4, 0x1c, 0x4c, 0xbe, 0x87, 0xad, 0x1e, 0x93, 0xe8, 0x2d, 0x14, 0xa4,
	0x5b, 0x33, 0x32, 0x8f, 0x16, 0x64, 0x8e, 0x99, 0xc4, 0xf9, 0x82, 0x38, 0x5b, 0xa3, 0x9b, 0xbd,
	0x25, 0xf6, 0x34, 0x2e, 0x5f, 0x70, 0xa5, 0x19, 0xd7, 0x6e, 0x7d, 0x45, 0x5c, 0xcf, 0x2c, 0x20,
	0x8d, 0x2b, 0x07, 0xa7, 0x44, 0x89, 0x4a, 0x24, 0xd2, 0x47, 0x17, 0x56, 0x10, 0xa9, 0x05, 0xa4,
	0xc4, 0x1c, 0x7c, 0xbc, 0x0e, 0xd5, 0xcb, 0x88, 0x07, 0xe7, 0xd5, 0x5a, 0xb9, 0x59, 0xa1, 0xc0,
	0xb4, 0x96, 0x51, 0x2f, 0xd1, 0xa8, 0x5a, 0x3f, 0xc3, 0xc6, 0xdc, 0x6b, 0x91, 0x37, 0x01, 0xa2,
	0x00, 0xb9, 0x8e, 0xc2, 0x08, 0xa5, 0x6d, 0x89, 0x82, 0x85, 0x7c, 0x00, 0xb7, 0xec, 0x74, 0xb1,
	0xbd, 0x31, 0xf3, 0x82, 0xaf, 0x32, 0xd7, 0x09, 0x86, 0x34, 0x87, 0x91, 0x6d, 0xa8, 0x0d, 0x51,
	0xb3, 0x80, 0x69, 0xe6, 0x56, 0x8c, 0xde, 0x64, 0xdf, 0x3a, 0x04, 0x98, 0x3e, 0x3a, 0xe9, 0xc0,
	0xa6, 0x79, 0x74, 0x2f, 0x8c, 0x62, 0xf4, 0x02, 0x0c, 0xbd, 0x88, 0x07, 0xf8, 0xa3, 0x89, 0xc2,
	0xa1, 0xf7, 0x8c, 0xef, 0x34, 0x8a, 0xf1, 0x04, 0xc3, 0xe7, 0xa9, 0xa3, 0xf5, 0x13, 0xdc, 0x99,
	0x29, 0x00, 0xb2, 0xbf, 0x74, 0x8e, 0xa4, 0x4d, 0x57, 0x5f, 0x32, 0x12, 0xc8, 0x21, 0xdc, 0xc9,
	0x51, 0x9e, 0x1a, 0xa1, 0x6f, 0xaf, 0xe4, 0x16, 0xaf, 0x94, 0xa3, 0xbb, 0x23, 0xf4, 0xe9, 0xed,
	0xb0, 0xb0, 0x6b, 0xfd, 0x59, 0x82, 0xff, 0x2d, 0x9d, 0x06, 0xe4, 0x21, 0xdc, 0xee, 0x89, 0x84,
	0x07, 0x5e, 0xc4, 0x47, 0x89, 0xce, 0x26, 0x99, 0x43, 0x1b, 0xc6, 0xf6, 0xdc, 0x98, 0x88, 0x07,
	0x6f, 0xf8, 0x8c, 0x0b, 0x1e, 0xf9, 0x2c, 0x8e, 0x5e, 0xa3, 0x85, 0x7a, 0x2a, 0xea, 0x73, 0xa6,
	0x13, 0x89, 0x26, 0x57, 0xf3, 0xad, 0x65, 0xc6, 0x75, 0x22, 0xd3, 0xca, 0x8d, 0x13, 0xa4, 0x0f,
	0x66, 0x14, 0x8c, 0x70, 0x37, 0xe7, 0x93, 0x53, 0x68, 0x8a, 0x44, 0xcf, 0x6a, 0x56, 0x6f, 0xd6,
	0xdc, 0xc8, 0x48, 0x13, 0x9d, 0xd6, 0x1f, 0x25, 0xb8, 0xbf, 0xb2, 0xcc, 0xc9, 0x47, 0xb0, 0xb5,
	0x90, 0x71, 0x8f, 0xb3, 0x21, 0xda, 0xda, 0xd9, 0x9c, 0xcf, 0xfa, 0x57, 0x6c, 0x88, 0xe4, 0x3d,
	0xb8, 0xc7, 0x64, 0x3f, 0x19, 0x22, 0xd7, 0xde, 0x25, 0x8e, 0xaf, 0x85, 0x0c, 0xb2, 0x24, 0xd5,
	0x69, 0x33, 0x77, 0x7c, 0x69, 0xed, 0xe4, 0x29, 0xec, 0xb0, 0x38, 0x16, 0xd7, 0x18, 0x78, 0x23,
	0xa1, 0xa2, 0x54, 0x84, 0xc5, 0x5e, 0x0e, 0x53, 0x26, 0x51, 0x15, 0xba, 0x6d, 0x31, 0x2f, 0x26,
	0x90, 0xa3, 0x1c, 0xd1, 0xda, 0xb7, 0x75, 0x92, 0xf7, 0x15, 0xd9, 0x81, 0xba, 0x18, 0xa1, 0x64,
	0xa6, 0xad, 0xb3, 0x40, 0xa7, 0x86, 0xd6, 0x6f, 0x65, 0x8b, 0xcf, 0xe7, 0x03, 0x79, 0x17, 0x9c,
	0x20, 0xfd, 0x97, 0x36, 0xd8, 0xbb, 0x07, 0x9b, 0xc5, 0x04, 0x9e, 0x30, 0xcd, 0x5e, 0x8e, 0x47,
	0x48, 0x33, 0x08, 0x39, 0x00, 0xc7, 0xfc, 0xef, 0xdb, 0x62, 0xda, 0x99, 0x99, 0xc3, 0x66, 0xd9,
	0x4d, 0xdd, 0x2f, 0xd2, 0x3f, 0x4f, 0x9a, 0x41, 0xd3, 0x78, 0xb4, 0x64, 0x11, 0x37, 0xd3, 0x2a,
	0xbd, 0x4f, 0x8d, 0x4e, 0x0d, 0xe4, 0x02, 0x36, 0xd4, 0x98, 0xfb, 0x03, 0x29, 0x78, 0xf4, 0x3a,
	0x8b, 0xb9, 0x6a, 0xe2, 0x78, 0x7b, 0xa6, 0xf7, 0x6c, 0xb0, 0xdd, 0x59, 0x28, 0x9d, 0xe7, 0x92,
	0x23, 0x68, 0xb0, 0x7e, 0x5f, 0x62, 0x3f, 0x93, 0x72, 0x8c, 0xd4, 0x5b, 0xcb, 0xa4, 0x8e, 0xa6,
	0x30, 0x5a, 0xe4, 0x10, 0x02, 0x55, 0xf3, 0xc6, 0xeb, 0x26, 0x75, 0x66, 0xdd, 0xfa, 0xbd, 0x04,
	0xb7, 0x8b, 0xcd, 0x42, 0x0e, 0xa1, 0x11, 0x26, 0x71, 0xcc, 0x64, 0xdf, 0xf4, 0x56, 0xe9, 0xe6,
	0xda, 0x2b, 0xe2, 0xc9, 0x03, 0xa8, 0x47, 0xca, 0x1b, 0xa2, 0x1e, 0x88, 0xc0, 0xe4, 0xb2, 0x46,
	0x6b, 0x91, 0xba, 0x30, 0x7b, 0x72, 0x02, 0x1b, 0xf3, 0xfd, 0xe2, 0xdc, 0xac, 0x7f, 0x37, 0x9a,
	0x69, 0x91, 0xf3, 0x6a, 0xad, 0xd2, 0xac, 0x9e, 0x57, 0x6b, 0xd5, 0xa6, 0xd3, 0x7a, 0xc7, 0xbe,
	0x79, 0x3e, 0x42, 0xc9, 0x16, 0xac, 0x07, 0x78, 0x15, 0xf9, 0x79, 0x25, 0xdb, 0xdd, 0xf1, 0xc5,
	0xb7, 0x67, 0xfd, 0x48, 0x0f, 0x92, 0x5e, 0xdb, 0x17, 0xc3, 0x4e, 0xe1, 0x9b, 0x69, 0xf9, 0xb2,
	0x2f, 0xb2, 0x8f, 0xa9, 0xf4, 0xc7, 0x33, 0x5f, 0x4c, 0xca, 0xeb, 0x8b, 0x6c, 0xf5, 0x57, 0xa9,
	0xd4, 0x5b, 0x37, 0xab, 0xc7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x10, 0xc2, 0x84, 0x45, 0x8d,
	0x0a, 0x00, 0x00,
}
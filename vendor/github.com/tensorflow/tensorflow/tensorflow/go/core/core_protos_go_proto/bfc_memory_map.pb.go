// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/bfc_memory_map.proto

package core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Some of the data from AllocatorStats
type MemAllocatorStats struct {
	NumAllocs            int64    `protobuf:"varint,1,opt,name=num_allocs,json=numAllocs,proto3" json:"num_allocs,omitempty"`
	BytesInUse           int64    `protobuf:"varint,2,opt,name=bytes_in_use,json=bytesInUse,proto3" json:"bytes_in_use,omitempty"`
	PeakBytesInUse       int64    `protobuf:"varint,3,opt,name=peak_bytes_in_use,json=peakBytesInUse,proto3" json:"peak_bytes_in_use,omitempty"`
	LargestAllocSize     int64    `protobuf:"varint,4,opt,name=largest_alloc_size,json=largestAllocSize,proto3" json:"largest_alloc_size,omitempty"`
	FragmentationMetric  float32  `protobuf:"fixed32,5,opt,name=fragmentation_metric,json=fragmentationMetric,proto3" json:"fragmentation_metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemAllocatorStats) Reset()         { *m = MemAllocatorStats{} }
func (m *MemAllocatorStats) String() string { return proto.CompactTextString(m) }
func (*MemAllocatorStats) ProtoMessage()    {}
func (*MemAllocatorStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{0}
}

func (m *MemAllocatorStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemAllocatorStats.Unmarshal(m, b)
}
func (m *MemAllocatorStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemAllocatorStats.Marshal(b, m, deterministic)
}
func (m *MemAllocatorStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemAllocatorStats.Merge(m, src)
}
func (m *MemAllocatorStats) XXX_Size() int {
	return xxx_messageInfo_MemAllocatorStats.Size(m)
}
func (m *MemAllocatorStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MemAllocatorStats.DiscardUnknown(m)
}

var xxx_messageInfo_MemAllocatorStats proto.InternalMessageInfo

func (m *MemAllocatorStats) GetNumAllocs() int64 {
	if m != nil {
		return m.NumAllocs
	}
	return 0
}

func (m *MemAllocatorStats) GetBytesInUse() int64 {
	if m != nil {
		return m.BytesInUse
	}
	return 0
}

func (m *MemAllocatorStats) GetPeakBytesInUse() int64 {
	if m != nil {
		return m.PeakBytesInUse
	}
	return 0
}

func (m *MemAllocatorStats) GetLargestAllocSize() int64 {
	if m != nil {
		return m.LargestAllocSize
	}
	return 0
}

func (m *MemAllocatorStats) GetFragmentationMetric() float32 {
	if m != nil {
		return m.FragmentationMetric
	}
	return 0
}

type MemChunk struct {
	Address              uint64   `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	RequestedSize        int64    `protobuf:"varint,3,opt,name=requested_size,json=requestedSize,proto3" json:"requested_size,omitempty"`
	Bin                  int32    `protobuf:"varint,4,opt,name=bin,proto3" json:"bin,omitempty"`
	OpName               string   `protobuf:"bytes,5,opt,name=op_name,json=opName,proto3" json:"op_name,omitempty"`
	FreedAtCount         uint64   `protobuf:"varint,6,opt,name=freed_at_count,json=freedAtCount,proto3" json:"freed_at_count,omitempty"`
	ActionCount          uint64   `protobuf:"varint,7,opt,name=action_count,json=actionCount,proto3" json:"action_count,omitempty"`
	InUse                bool     `protobuf:"varint,8,opt,name=in_use,json=inUse,proto3" json:"in_use,omitempty"`
	StepId               uint64   `protobuf:"varint,9,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemChunk) Reset()         { *m = MemChunk{} }
func (m *MemChunk) String() string { return proto.CompactTextString(m) }
func (*MemChunk) ProtoMessage()    {}
func (*MemChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{1}
}

func (m *MemChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemChunk.Unmarshal(m, b)
}
func (m *MemChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemChunk.Marshal(b, m, deterministic)
}
func (m *MemChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemChunk.Merge(m, src)
}
func (m *MemChunk) XXX_Size() int {
	return xxx_messageInfo_MemChunk.Size(m)
}
func (m *MemChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_MemChunk.DiscardUnknown(m)
}

var xxx_messageInfo_MemChunk proto.InternalMessageInfo

func (m *MemChunk) GetAddress() uint64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *MemChunk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MemChunk) GetRequestedSize() int64 {
	if m != nil {
		return m.RequestedSize
	}
	return 0
}

func (m *MemChunk) GetBin() int32 {
	if m != nil {
		return m.Bin
	}
	return 0
}

func (m *MemChunk) GetOpName() string {
	if m != nil {
		return m.OpName
	}
	return ""
}

func (m *MemChunk) GetFreedAtCount() uint64 {
	if m != nil {
		return m.FreedAtCount
	}
	return 0
}

func (m *MemChunk) GetActionCount() uint64 {
	if m != nil {
		return m.ActionCount
	}
	return 0
}

func (m *MemChunk) GetInUse() bool {
	if m != nil {
		return m.InUse
	}
	return false
}

func (m *MemChunk) GetStepId() uint64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

type BinSummary struct {
	Bin                  int32    `protobuf:"varint,1,opt,name=bin,proto3" json:"bin,omitempty"`
	TotalBytesInUse      int64    `protobuf:"varint,2,opt,name=total_bytes_in_use,json=totalBytesInUse,proto3" json:"total_bytes_in_use,omitempty"`
	TotalBytesInBin      int64    `protobuf:"varint,3,opt,name=total_bytes_in_bin,json=totalBytesInBin,proto3" json:"total_bytes_in_bin,omitempty"`
	TotalChunksInUse     int64    `protobuf:"varint,4,opt,name=total_chunks_in_use,json=totalChunksInUse,proto3" json:"total_chunks_in_use,omitempty"`
	TotalChunksInBin     int64    `protobuf:"varint,5,opt,name=total_chunks_in_bin,json=totalChunksInBin,proto3" json:"total_chunks_in_bin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinSummary) Reset()         { *m = BinSummary{} }
func (m *BinSummary) String() string { return proto.CompactTextString(m) }
func (*BinSummary) ProtoMessage()    {}
func (*BinSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{2}
}

func (m *BinSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinSummary.Unmarshal(m, b)
}
func (m *BinSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinSummary.Marshal(b, m, deterministic)
}
func (m *BinSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinSummary.Merge(m, src)
}
func (m *BinSummary) XXX_Size() int {
	return xxx_messageInfo_BinSummary.Size(m)
}
func (m *BinSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BinSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BinSummary proto.InternalMessageInfo

func (m *BinSummary) GetBin() int32 {
	if m != nil {
		return m.Bin
	}
	return 0
}

func (m *BinSummary) GetTotalBytesInUse() int64 {
	if m != nil {
		return m.TotalBytesInUse
	}
	return 0
}

func (m *BinSummary) GetTotalBytesInBin() int64 {
	if m != nil {
		return m.TotalBytesInBin
	}
	return 0
}

func (m *BinSummary) GetTotalChunksInUse() int64 {
	if m != nil {
		return m.TotalChunksInUse
	}
	return 0
}

func (m *BinSummary) GetTotalChunksInBin() int64 {
	if m != nil {
		return m.TotalChunksInBin
	}
	return 0
}

type SnapShot struct {
	ActionCount          uint64   `protobuf:"varint,1,opt,name=action_count,json=actionCount,proto3" json:"action_count,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapShot) Reset()         { *m = SnapShot{} }
func (m *SnapShot) String() string { return proto.CompactTextString(m) }
func (*SnapShot) ProtoMessage()    {}
func (*SnapShot) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{3}
}

func (m *SnapShot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapShot.Unmarshal(m, b)
}
func (m *SnapShot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapShot.Marshal(b, m, deterministic)
}
func (m *SnapShot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapShot.Merge(m, src)
}
func (m *SnapShot) XXX_Size() int {
	return xxx_messageInfo_SnapShot.Size(m)
}
func (m *SnapShot) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapShot.DiscardUnknown(m)
}

var xxx_messageInfo_SnapShot proto.InternalMessageInfo

func (m *SnapShot) GetActionCount() uint64 {
	if m != nil {
		return m.ActionCount
	}
	return 0
}

func (m *SnapShot) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MemoryDump struct {
	AllocatorName        string             `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	BinSummary           []*BinSummary      `protobuf:"bytes,2,rep,name=bin_summary,json=binSummary,proto3" json:"bin_summary,omitempty"`
	Chunk                []*MemChunk        `protobuf:"bytes,3,rep,name=chunk,proto3" json:"chunk,omitempty"`
	SnapShot             []*SnapShot        `protobuf:"bytes,4,rep,name=snap_shot,json=snapShot,proto3" json:"snap_shot,omitempty"`
	Stats                *MemAllocatorStats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MemoryDump) Reset()         { *m = MemoryDump{} }
func (m *MemoryDump) String() string { return proto.CompactTextString(m) }
func (*MemoryDump) ProtoMessage()    {}
func (*MemoryDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{4}
}

func (m *MemoryDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryDump.Unmarshal(m, b)
}
func (m *MemoryDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryDump.Marshal(b, m, deterministic)
}
func (m *MemoryDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryDump.Merge(m, src)
}
func (m *MemoryDump) XXX_Size() int {
	return xxx_messageInfo_MemoryDump.Size(m)
}
func (m *MemoryDump) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryDump.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryDump proto.InternalMessageInfo

func (m *MemoryDump) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *MemoryDump) GetBinSummary() []*BinSummary {
	if m != nil {
		return m.BinSummary
	}
	return nil
}

func (m *MemoryDump) GetChunk() []*MemChunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *MemoryDump) GetSnapShot() []*SnapShot {
	if m != nil {
		return m.SnapShot
	}
	return nil
}

func (m *MemoryDump) GetStats() *MemAllocatorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*MemAllocatorStats)(nil), "tensorflow.MemAllocatorStats")
	proto.RegisterType((*MemChunk)(nil), "tensorflow.MemChunk")
	proto.RegisterType((*BinSummary)(nil), "tensorflow.BinSummary")
	proto.RegisterType((*SnapShot)(nil), "tensorflow.SnapShot")
	proto.RegisterType((*MemoryDump)(nil), "tensorflow.MemoryDump")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/bfc_memory_map.proto", fileDescriptor_fdf22777007c1f3d)
}

var fileDescriptor_fdf22777007c1f3d = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x3d, 0x6f, 0xdb, 0x3c,
	0x10, 0xc7, 0x21, 0xdb, 0xf2, 0xcb, 0x39, 0xc9, 0x93, 0x30, 0x79, 0x1a, 0x2d, 0x01, 0x5c, 0xa3,
	0x05, 0xdc, 0x97, 0xd8, 0x48, 0x32, 0x74, 0x8e, 0xd3, 0xa1, 0x29, 0xe0, 0x0e, 0x32, 0xba, 0x74,
	0x21, 0x28, 0x89, 0xb6, 0x89, 0x98, 0xa4, 0x4a, 0x52, 0x28, 0x92, 0xa9, 0x40, 0xbf, 0x63, 0xe7,
	0x7e, 0x94, 0x42, 0x47, 0x2b, 0x4a, 0x6a, 0x2f, 0x36, 0x79, 0xf7, 0xe3, 0xf1, 0xee, 0xfe, 0x3a,
	0xc2, 0xb9, 0xe3, 0xca, 0x6a, 0xb3, 0x58, 0xeb, 0x1f, 0x93, 0x54, 0x1b, 0x3e, 0xc9, 0x8d, 0x76,
	0x3a, 0x29, 0x16, 0x93, 0x64, 0x91, 0x52, 0xc9, 0xa5, 0x36, 0xf7, 0x54, 0xb2, 0x7c, 0x8c, 0x76,
	0x02, 0x35, 0x3e, 0xfc, 0x13, 0xc0, 0xd1, 0x8c, 0xcb, 0xeb, 0xf5, 0x5a, 0xa7, 0xcc, 0x69, 0x33,
	0x77, 0xcc, 0x59, 0x72, 0x06, 0xa0, 0x0a, 0x49, 0x59, 0x69, 0xb5, 0x51, 0x30, 0x08, 0x46, 0xcd,
	0xb8, 0xa7, 0x0a, 0x8f, 0x59, 0x32, 0x80, 0xbd, 0xe4, 0xde, 0x71, 0x4b, 0x85, 0xa2, 0x85, 0xe5,
	0x51, 0x03, 0x01, 0x40, 0xdb, 0xad, 0xfa, 0x6a, 0x39, 0x79, 0x03, 0x47, 0x39, 0x67, 0x77, 0xf4,
	0x19, 0xd6, 0x44, 0xec, 0xa0, 0x74, 0x4c, 0x6b, 0xf4, 0x3d, 0x90, 0x35, 0x33, 0x4b, 0x6e, 0x9d,
	0xbf, 0x8f, 0x5a, 0xf1, 0xc0, 0xa3, 0x16, 0xb2, 0x87, 0x1b, 0x0f, 0xde, 0x3b, 0x17, 0x0f, 0x9c,
	0x5c, 0xc0, 0xc9, 0xc2, 0xb0, 0xa5, 0xe4, 0xca, 0x31, 0x27, 0xb4, 0xa2, 0x92, 0x3b, 0x23, 0xd2,
	0x28, 0x1c, 0x04, 0xa3, 0x46, 0x7c, 0xfc, 0xcc, 0x37, 0x43, 0xd7, 0xf0, 0x57, 0x03, 0xba, 0x33,
	0x2e, 0x6f, 0x56, 0x85, 0xba, 0x23, 0x11, 0x74, 0x58, 0x96, 0x19, 0x6e, 0x7d, 0x59, 0xad, 0xb8,
	0xda, 0x12, 0x02, 0x2d, 0xbc, 0xd9, 0x17, 0x83, 0x6b, 0xf2, 0x1a, 0x0e, 0x0c, 0xff, 0x5e, 0x70,
	0xeb, 0x78, 0xe6, 0xf3, 0xf2, 0x35, 0xec, 0x3f, 0x5a, 0x31, 0xa9, 0x43, 0x68, 0x26, 0x42, 0x61,
	0xce, 0x61, 0x5c, 0x2e, 0xc9, 0x29, 0x74, 0x74, 0x4e, 0x15, 0x93, 0x1c, 0x33, 0xeb, 0xc5, 0x6d,
	0x9d, 0x7f, 0x61, 0x92, 0x93, 0x57, 0x70, 0xb0, 0x30, 0x9c, 0x67, 0x94, 0x39, 0x9a, 0xea, 0x42,
	0xb9, 0xa8, 0x8d, 0x69, 0xec, 0xa1, 0xf5, 0xda, 0xdd, 0x94, 0x36, 0xf2, 0x12, 0xf6, 0x58, 0x8a,
	0xe5, 0x79, 0xa6, 0x83, 0x4c, 0xdf, 0xdb, 0x3c, 0xf2, 0x3f, 0xb4, 0x37, 0x6d, 0xed, 0x0e, 0x82,
	0x51, 0x37, 0x0e, 0x05, 0x76, 0xf3, 0x14, 0x3a, 0xd6, 0xf1, 0x9c, 0x8a, 0x2c, 0xea, 0xe1, 0xa1,
	0x76, 0xb9, 0xbd, 0xcd, 0x86, 0xbf, 0x03, 0x80, 0xa9, 0x50, 0xf3, 0x42, 0x4a, 0x66, 0xee, 0xab,
	0x94, 0x83, 0x3a, 0xe5, 0x77, 0x40, 0x9c, 0x76, 0x6c, 0x4d, 0x77, 0x48, 0xfb, 0x1f, 0x7a, 0x9e,
	0x88, 0xb6, 0x0d, 0x97, 0xd1, 0x9a, 0xdb, 0xf0, 0x54, 0x28, 0x72, 0x0e, 0xc7, 0x1e, 0x4e, 0x4b,
	0x09, 0x1e, 0x43, 0x6f, 0x24, 0x46, 0x17, 0x8a, 0xb3, 0x89, 0xbd, 0x03, 0x2f, 0x83, 0x87, 0x3b,
	0xf0, 0xa9, 0x50, 0xc3, 0x6b, 0xe8, 0xce, 0x15, 0xcb, 0xe7, 0x2b, 0xbd, 0xdd, 0xb7, 0x60, 0xbb,
	0x6f, 0x3b, 0x64, 0x1e, 0xfe, 0x6c, 0x00, 0xcc, 0x70, 0x4a, 0x3e, 0x16, 0x32, 0x2f, 0x55, 0x67,
	0xd5, 0x3c, 0x78, 0x0d, 0x03, 0xd4, 0x70, 0xff, 0xd1, 0x8a, 0x52, 0x7e, 0x80, 0x7e, 0x22, 0x14,
	0xb5, 0xbe, 0xa3, 0x51, 0x63, 0xd0, 0x1c, 0xf5, 0x2f, 0x5f, 0x8c, 0xeb, 0xe1, 0x1a, 0xd7, 0xfd,
	0x8e, 0x21, 0xa9, 0x7b, 0xff, 0x16, 0x42, 0x2c, 0x2d, 0x6a, 0xe2, 0x91, 0x93, 0xa7, 0x47, 0xaa,
	0x0f, 0x35, 0xf6, 0x08, 0xb9, 0x80, 0x9e, 0x55, 0x2c, 0xa7, 0x76, 0xa5, 0x5d, 0xd4, 0xda, 0xe6,
	0xab, 0xd2, 0xe3, 0xae, 0xad, 0x9a, 0x70, 0x05, 0xa1, 0x2d, 0xa7, 0x18, 0x3b, 0xd6, 0xbf, 0x3c,
	0xfb, 0x27, 0xfc, 0xf3, 0x51, 0x8f, 0x3d, 0x3b, 0xfd, 0xfc, 0xed, 0xd3, 0x52, 0xb8, 0x55, 0x91,
	0x8c, 0x53, 0x2d, 0x27, 0x4f, 0xde, 0x93, 0xdd, 0xcb, 0xa5, 0xf6, 0x0f, 0x4d, 0xf9, 0x43, 0xf1,
	0x55, 0xb1, 0x74, 0xa9, 0xfd, 0x2a, 0x69, 0xe3, 0xdf, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x26, 0xd5, 0xae, 0x9b, 0x97, 0x04, 0x00, 0x00,
}
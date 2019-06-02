// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/state/object.proto

package gvisor_state_statefile

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

type Slice struct {
	Length               uint32   `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Capacity             uint32   `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	RefValue             uint64   `protobuf:"varint,3,opt,name=ref_value,json=refValue,proto3" json:"ref_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slice) Reset()         { *m = Slice{} }
func (m *Slice) String() string { return proto.CompactTextString(m) }
func (*Slice) ProtoMessage()    {}
func (*Slice) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{0}
}

func (m *Slice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slice.Unmarshal(m, b)
}
func (m *Slice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slice.Marshal(b, m, deterministic)
}
func (m *Slice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slice.Merge(m, src)
}
func (m *Slice) XXX_Size() int {
	return xxx_messageInfo_Slice.Size(m)
}
func (m *Slice) XXX_DiscardUnknown() {
	xxx_messageInfo_Slice.DiscardUnknown(m)
}

var xxx_messageInfo_Slice proto.InternalMessageInfo

func (m *Slice) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Slice) GetCapacity() uint32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Slice) GetRefValue() uint64 {
	if m != nil {
		return m.RefValue
	}
	return 0
}

type Array struct {
	Contents             []*Object `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Array) Reset()         { *m = Array{} }
func (m *Array) String() string { return proto.CompactTextString(m) }
func (*Array) ProtoMessage()    {}
func (*Array) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{1}
}

func (m *Array) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Array.Unmarshal(m, b)
}
func (m *Array) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Array.Marshal(b, m, deterministic)
}
func (m *Array) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Array.Merge(m, src)
}
func (m *Array) XXX_Size() int {
	return xxx_messageInfo_Array.Size(m)
}
func (m *Array) XXX_DiscardUnknown() {
	xxx_messageInfo_Array.DiscardUnknown(m)
}

var xxx_messageInfo_Array proto.InternalMessageInfo

func (m *Array) GetContents() []*Object {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Map struct {
	Keys                 []*Object `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Values               []*Object `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}
func (*Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{2}
}

func (m *Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map.Unmarshal(m, b)
}
func (m *Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map.Marshal(b, m, deterministic)
}
func (m *Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map.Merge(m, src)
}
func (m *Map) XXX_Size() int {
	return xxx_messageInfo_Map.Size(m)
}
func (m *Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Map proto.InternalMessageInfo

func (m *Map) GetKeys() []*Object {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Map) GetValues() []*Object {
	if m != nil {
		return m.Values
	}
	return nil
}

type Interface struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                *Object  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Interface) Reset()         { *m = Interface{} }
func (m *Interface) String() string { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()    {}
func (*Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{3}
}

func (m *Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interface.Unmarshal(m, b)
}
func (m *Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interface.Marshal(b, m, deterministic)
}
func (m *Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interface.Merge(m, src)
}
func (m *Interface) XXX_Size() int {
	return xxx_messageInfo_Interface.Size(m)
}
func (m *Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Interface proto.InternalMessageInfo

func (m *Interface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Interface) GetValue() *Object {
	if m != nil {
		return m.Value
	}
	return nil
}

type Struct struct {
	Fields               []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Struct) Reset()         { *m = Struct{} }
func (m *Struct) String() string { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()    {}
func (*Struct) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{4}
}

func (m *Struct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Struct.Unmarshal(m, b)
}
func (m *Struct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Struct.Marshal(b, m, deterministic)
}
func (m *Struct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Struct.Merge(m, src)
}
func (m *Struct) XXX_Size() int {
	return xxx_messageInfo_Struct.Size(m)
}
func (m *Struct) XXX_DiscardUnknown() {
	xxx_messageInfo_Struct.DiscardUnknown(m)
}

var xxx_messageInfo_Struct proto.InternalMessageInfo

func (m *Struct) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Field struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                *Object  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{5}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetValue() *Object {
	if m != nil {
		return m.Value
	}
	return nil
}

type Uint16S struct {
	Values               []uint32 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint16S) Reset()         { *m = Uint16S{} }
func (m *Uint16S) String() string { return proto.CompactTextString(m) }
func (*Uint16S) ProtoMessage()    {}
func (*Uint16S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{6}
}

func (m *Uint16S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint16S.Unmarshal(m, b)
}
func (m *Uint16S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint16S.Marshal(b, m, deterministic)
}
func (m *Uint16S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint16S.Merge(m, src)
}
func (m *Uint16S) XXX_Size() int {
	return xxx_messageInfo_Uint16S.Size(m)
}
func (m *Uint16S) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint16S.DiscardUnknown(m)
}

var xxx_messageInfo_Uint16S proto.InternalMessageInfo

func (m *Uint16S) GetValues() []uint32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Uint32S struct {
	Values               []uint32 `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint32S) Reset()         { *m = Uint32S{} }
func (m *Uint32S) String() string { return proto.CompactTextString(m) }
func (*Uint32S) ProtoMessage()    {}
func (*Uint32S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{7}
}

func (m *Uint32S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint32S.Unmarshal(m, b)
}
func (m *Uint32S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint32S.Marshal(b, m, deterministic)
}
func (m *Uint32S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint32S.Merge(m, src)
}
func (m *Uint32S) XXX_Size() int {
	return xxx_messageInfo_Uint32S.Size(m)
}
func (m *Uint32S) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint32S.DiscardUnknown(m)
}

var xxx_messageInfo_Uint32S proto.InternalMessageInfo

func (m *Uint32S) GetValues() []uint32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Uint64S struct {
	Values               []uint64 `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint64S) Reset()         { *m = Uint64S{} }
func (m *Uint64S) String() string { return proto.CompactTextString(m) }
func (*Uint64S) ProtoMessage()    {}
func (*Uint64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{8}
}

func (m *Uint64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint64S.Unmarshal(m, b)
}
func (m *Uint64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint64S.Marshal(b, m, deterministic)
}
func (m *Uint64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint64S.Merge(m, src)
}
func (m *Uint64S) XXX_Size() int {
	return xxx_messageInfo_Uint64S.Size(m)
}
func (m *Uint64S) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint64S.DiscardUnknown(m)
}

var xxx_messageInfo_Uint64S proto.InternalMessageInfo

func (m *Uint64S) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Uintptrs struct {
	Values               []uint64 `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uintptrs) Reset()         { *m = Uintptrs{} }
func (m *Uintptrs) String() string { return proto.CompactTextString(m) }
func (*Uintptrs) ProtoMessage()    {}
func (*Uintptrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{9}
}

func (m *Uintptrs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uintptrs.Unmarshal(m, b)
}
func (m *Uintptrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uintptrs.Marshal(b, m, deterministic)
}
func (m *Uintptrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uintptrs.Merge(m, src)
}
func (m *Uintptrs) XXX_Size() int {
	return xxx_messageInfo_Uintptrs.Size(m)
}
func (m *Uintptrs) XXX_DiscardUnknown() {
	xxx_messageInfo_Uintptrs.DiscardUnknown(m)
}

var xxx_messageInfo_Uintptrs proto.InternalMessageInfo

func (m *Uintptrs) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Int8S struct {
	Values               []byte   `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int8S) Reset()         { *m = Int8S{} }
func (m *Int8S) String() string { return proto.CompactTextString(m) }
func (*Int8S) ProtoMessage()    {}
func (*Int8S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{10}
}

func (m *Int8S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int8S.Unmarshal(m, b)
}
func (m *Int8S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int8S.Marshal(b, m, deterministic)
}
func (m *Int8S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int8S.Merge(m, src)
}
func (m *Int8S) XXX_Size() int {
	return xxx_messageInfo_Int8S.Size(m)
}
func (m *Int8S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int8S.DiscardUnknown(m)
}

var xxx_messageInfo_Int8S proto.InternalMessageInfo

func (m *Int8S) GetValues() []byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type Int16S struct {
	Values               []int32  `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int16S) Reset()         { *m = Int16S{} }
func (m *Int16S) String() string { return proto.CompactTextString(m) }
func (*Int16S) ProtoMessage()    {}
func (*Int16S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{11}
}

func (m *Int16S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int16S.Unmarshal(m, b)
}
func (m *Int16S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int16S.Marshal(b, m, deterministic)
}
func (m *Int16S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int16S.Merge(m, src)
}
func (m *Int16S) XXX_Size() int {
	return xxx_messageInfo_Int16S.Size(m)
}
func (m *Int16S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int16S.DiscardUnknown(m)
}

var xxx_messageInfo_Int16S proto.InternalMessageInfo

func (m *Int16S) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Int32S struct {
	Values               []int32  `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32S) Reset()         { *m = Int32S{} }
func (m *Int32S) String() string { return proto.CompactTextString(m) }
func (*Int32S) ProtoMessage()    {}
func (*Int32S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{12}
}

func (m *Int32S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32S.Unmarshal(m, b)
}
func (m *Int32S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32S.Marshal(b, m, deterministic)
}
func (m *Int32S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32S.Merge(m, src)
}
func (m *Int32S) XXX_Size() int {
	return xxx_messageInfo_Int32S.Size(m)
}
func (m *Int32S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32S.DiscardUnknown(m)
}

var xxx_messageInfo_Int32S proto.InternalMessageInfo

func (m *Int32S) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Int64S struct {
	Values               []int64  `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64S) Reset()         { *m = Int64S{} }
func (m *Int64S) String() string { return proto.CompactTextString(m) }
func (*Int64S) ProtoMessage()    {}
func (*Int64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{13}
}

func (m *Int64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64S.Unmarshal(m, b)
}
func (m *Int64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64S.Marshal(b, m, deterministic)
}
func (m *Int64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64S.Merge(m, src)
}
func (m *Int64S) XXX_Size() int {
	return xxx_messageInfo_Int64S.Size(m)
}
func (m *Int64S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64S.DiscardUnknown(m)
}

var xxx_messageInfo_Int64S proto.InternalMessageInfo

func (m *Int64S) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Bools struct {
	Values               []bool   `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bools) Reset()         { *m = Bools{} }
func (m *Bools) String() string { return proto.CompactTextString(m) }
func (*Bools) ProtoMessage()    {}
func (*Bools) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{14}
}

func (m *Bools) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bools.Unmarshal(m, b)
}
func (m *Bools) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bools.Marshal(b, m, deterministic)
}
func (m *Bools) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bools.Merge(m, src)
}
func (m *Bools) XXX_Size() int {
	return xxx_messageInfo_Bools.Size(m)
}
func (m *Bools) XXX_DiscardUnknown() {
	xxx_messageInfo_Bools.DiscardUnknown(m)
}

var xxx_messageInfo_Bools proto.InternalMessageInfo

func (m *Bools) GetValues() []bool {
	if m != nil {
		return m.Values
	}
	return nil
}

type Float64S struct {
	Values               []float64 `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Float64S) Reset()         { *m = Float64S{} }
func (m *Float64S) String() string { return proto.CompactTextString(m) }
func (*Float64S) ProtoMessage()    {}
func (*Float64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{15}
}

func (m *Float64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float64S.Unmarshal(m, b)
}
func (m *Float64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float64S.Marshal(b, m, deterministic)
}
func (m *Float64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float64S.Merge(m, src)
}
func (m *Float64S) XXX_Size() int {
	return xxx_messageInfo_Float64S.Size(m)
}
func (m *Float64S) XXX_DiscardUnknown() {
	xxx_messageInfo_Float64S.DiscardUnknown(m)
}

var xxx_messageInfo_Float64S proto.InternalMessageInfo

func (m *Float64S) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Float32S struct {
	Values               []float32 `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Float32S) Reset()         { *m = Float32S{} }
func (m *Float32S) String() string { return proto.CompactTextString(m) }
func (*Float32S) ProtoMessage()    {}
func (*Float32S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{16}
}

func (m *Float32S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float32S.Unmarshal(m, b)
}
func (m *Float32S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float32S.Marshal(b, m, deterministic)
}
func (m *Float32S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float32S.Merge(m, src)
}
func (m *Float32S) XXX_Size() int {
	return xxx_messageInfo_Float32S.Size(m)
}
func (m *Float32S) XXX_DiscardUnknown() {
	xxx_messageInfo_Float32S.DiscardUnknown(m)
}

var xxx_messageInfo_Float32S proto.InternalMessageInfo

func (m *Float32S) GetValues() []float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Object struct {
	// Types that are valid to be assigned to Value:
	//	*Object_BoolValue
	//	*Object_StringValue
	//	*Object_Int64Value
	//	*Object_Uint64Value
	//	*Object_DoubleValue
	//	*Object_RefValue
	//	*Object_SliceValue
	//	*Object_ArrayValue
	//	*Object_InterfaceValue
	//	*Object_StructValue
	//	*Object_MapValue
	//	*Object_ByteArrayValue
	//	*Object_Uint16ArrayValue
	//	*Object_Uint32ArrayValue
	//	*Object_Uint64ArrayValue
	//	*Object_UintptrArrayValue
	//	*Object_Int8ArrayValue
	//	*Object_Int16ArrayValue
	//	*Object_Int32ArrayValue
	//	*Object_Int64ArrayValue
	//	*Object_BoolArrayValue
	//	*Object_Float64ArrayValue
	//	*Object_Float32ArrayValue
	Value                isObject_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dee2c1912d4d62d, []int{17}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

type isObject_Value interface {
	isObject_Value()
}

type Object_BoolValue struct {
	BoolValue bool `protobuf:"varint,1,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Object_StringValue struct {
	StringValue []byte `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Object_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Object_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,4,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type Object_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Object_RefValue struct {
	RefValue uint64 `protobuf:"varint,6,opt,name=ref_value,json=refValue,proto3,oneof"`
}

type Object_SliceValue struct {
	SliceValue *Slice `protobuf:"bytes,7,opt,name=slice_value,json=sliceValue,proto3,oneof"`
}

type Object_ArrayValue struct {
	ArrayValue *Array `protobuf:"bytes,8,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

type Object_InterfaceValue struct {
	InterfaceValue *Interface `protobuf:"bytes,9,opt,name=interface_value,json=interfaceValue,proto3,oneof"`
}

type Object_StructValue struct {
	StructValue *Struct `protobuf:"bytes,10,opt,name=struct_value,json=structValue,proto3,oneof"`
}

type Object_MapValue struct {
	MapValue *Map `protobuf:"bytes,11,opt,name=map_value,json=mapValue,proto3,oneof"`
}

type Object_ByteArrayValue struct {
	ByteArrayValue []byte `protobuf:"bytes,12,opt,name=byte_array_value,json=byteArrayValue,proto3,oneof"`
}

type Object_Uint16ArrayValue struct {
	Uint16ArrayValue *Uint16S `protobuf:"bytes,13,opt,name=uint16_array_value,json=uint16ArrayValue,proto3,oneof"`
}

type Object_Uint32ArrayValue struct {
	Uint32ArrayValue *Uint32S `protobuf:"bytes,14,opt,name=uint32_array_value,json=uint32ArrayValue,proto3,oneof"`
}

type Object_Uint64ArrayValue struct {
	Uint64ArrayValue *Uint64S `protobuf:"bytes,15,opt,name=uint64_array_value,json=uint64ArrayValue,proto3,oneof"`
}

type Object_UintptrArrayValue struct {
	UintptrArrayValue *Uintptrs `protobuf:"bytes,16,opt,name=uintptr_array_value,json=uintptrArrayValue,proto3,oneof"`
}

type Object_Int8ArrayValue struct {
	Int8ArrayValue *Int8S `protobuf:"bytes,17,opt,name=int8_array_value,json=int8ArrayValue,proto3,oneof"`
}

type Object_Int16ArrayValue struct {
	Int16ArrayValue *Int16S `protobuf:"bytes,18,opt,name=int16_array_value,json=int16ArrayValue,proto3,oneof"`
}

type Object_Int32ArrayValue struct {
	Int32ArrayValue *Int32S `protobuf:"bytes,19,opt,name=int32_array_value,json=int32ArrayValue,proto3,oneof"`
}

type Object_Int64ArrayValue struct {
	Int64ArrayValue *Int64S `protobuf:"bytes,20,opt,name=int64_array_value,json=int64ArrayValue,proto3,oneof"`
}

type Object_BoolArrayValue struct {
	BoolArrayValue *Bools `protobuf:"bytes,21,opt,name=bool_array_value,json=boolArrayValue,proto3,oneof"`
}

type Object_Float64ArrayValue struct {
	Float64ArrayValue *Float64S `protobuf:"bytes,22,opt,name=float64_array_value,json=float64ArrayValue,proto3,oneof"`
}

type Object_Float32ArrayValue struct {
	Float32ArrayValue *Float32S `protobuf:"bytes,23,opt,name=float32_array_value,json=float32ArrayValue,proto3,oneof"`
}

func (*Object_BoolValue) isObject_Value() {}

func (*Object_StringValue) isObject_Value() {}

func (*Object_Int64Value) isObject_Value() {}

func (*Object_Uint64Value) isObject_Value() {}

func (*Object_DoubleValue) isObject_Value() {}

func (*Object_RefValue) isObject_Value() {}

func (*Object_SliceValue) isObject_Value() {}

func (*Object_ArrayValue) isObject_Value() {}

func (*Object_InterfaceValue) isObject_Value() {}

func (*Object_StructValue) isObject_Value() {}

func (*Object_MapValue) isObject_Value() {}

func (*Object_ByteArrayValue) isObject_Value() {}

func (*Object_Uint16ArrayValue) isObject_Value() {}

func (*Object_Uint32ArrayValue) isObject_Value() {}

func (*Object_Uint64ArrayValue) isObject_Value() {}

func (*Object_UintptrArrayValue) isObject_Value() {}

func (*Object_Int8ArrayValue) isObject_Value() {}

func (*Object_Int16ArrayValue) isObject_Value() {}

func (*Object_Int32ArrayValue) isObject_Value() {}

func (*Object_Int64ArrayValue) isObject_Value() {}

func (*Object_BoolArrayValue) isObject_Value() {}

func (*Object_Float64ArrayValue) isObject_Value() {}

func (*Object_Float32ArrayValue) isObject_Value() {}

func (m *Object) GetValue() isObject_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Object) GetBoolValue() bool {
	if x, ok := m.GetValue().(*Object_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Object) GetStringValue() []byte {
	if x, ok := m.GetValue().(*Object_StringValue); ok {
		return x.StringValue
	}
	return nil
}

func (m *Object) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Object_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Object) GetUint64Value() uint64 {
	if x, ok := m.GetValue().(*Object_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *Object) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Object_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Object) GetRefValue() uint64 {
	if x, ok := m.GetValue().(*Object_RefValue); ok {
		return x.RefValue
	}
	return 0
}

func (m *Object) GetSliceValue() *Slice {
	if x, ok := m.GetValue().(*Object_SliceValue); ok {
		return x.SliceValue
	}
	return nil
}

func (m *Object) GetArrayValue() *Array {
	if x, ok := m.GetValue().(*Object_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (m *Object) GetInterfaceValue() *Interface {
	if x, ok := m.GetValue().(*Object_InterfaceValue); ok {
		return x.InterfaceValue
	}
	return nil
}

func (m *Object) GetStructValue() *Struct {
	if x, ok := m.GetValue().(*Object_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (m *Object) GetMapValue() *Map {
	if x, ok := m.GetValue().(*Object_MapValue); ok {
		return x.MapValue
	}
	return nil
}

func (m *Object) GetByteArrayValue() []byte {
	if x, ok := m.GetValue().(*Object_ByteArrayValue); ok {
		return x.ByteArrayValue
	}
	return nil
}

func (m *Object) GetUint16ArrayValue() *Uint16S {
	if x, ok := m.GetValue().(*Object_Uint16ArrayValue); ok {
		return x.Uint16ArrayValue
	}
	return nil
}

func (m *Object) GetUint32ArrayValue() *Uint32S {
	if x, ok := m.GetValue().(*Object_Uint32ArrayValue); ok {
		return x.Uint32ArrayValue
	}
	return nil
}

func (m *Object) GetUint64ArrayValue() *Uint64S {
	if x, ok := m.GetValue().(*Object_Uint64ArrayValue); ok {
		return x.Uint64ArrayValue
	}
	return nil
}

func (m *Object) GetUintptrArrayValue() *Uintptrs {
	if x, ok := m.GetValue().(*Object_UintptrArrayValue); ok {
		return x.UintptrArrayValue
	}
	return nil
}

func (m *Object) GetInt8ArrayValue() *Int8S {
	if x, ok := m.GetValue().(*Object_Int8ArrayValue); ok {
		return x.Int8ArrayValue
	}
	return nil
}

func (m *Object) GetInt16ArrayValue() *Int16S {
	if x, ok := m.GetValue().(*Object_Int16ArrayValue); ok {
		return x.Int16ArrayValue
	}
	return nil
}

func (m *Object) GetInt32ArrayValue() *Int32S {
	if x, ok := m.GetValue().(*Object_Int32ArrayValue); ok {
		return x.Int32ArrayValue
	}
	return nil
}

func (m *Object) GetInt64ArrayValue() *Int64S {
	if x, ok := m.GetValue().(*Object_Int64ArrayValue); ok {
		return x.Int64ArrayValue
	}
	return nil
}

func (m *Object) GetBoolArrayValue() *Bools {
	if x, ok := m.GetValue().(*Object_BoolArrayValue); ok {
		return x.BoolArrayValue
	}
	return nil
}

func (m *Object) GetFloat64ArrayValue() *Float64S {
	if x, ok := m.GetValue().(*Object_Float64ArrayValue); ok {
		return x.Float64ArrayValue
	}
	return nil
}

func (m *Object) GetFloat32ArrayValue() *Float32S {
	if x, ok := m.GetValue().(*Object_Float32ArrayValue); ok {
		return x.Float32ArrayValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Object) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Object_BoolValue)(nil),
		(*Object_StringValue)(nil),
		(*Object_Int64Value)(nil),
		(*Object_Uint64Value)(nil),
		(*Object_DoubleValue)(nil),
		(*Object_RefValue)(nil),
		(*Object_SliceValue)(nil),
		(*Object_ArrayValue)(nil),
		(*Object_InterfaceValue)(nil),
		(*Object_StructValue)(nil),
		(*Object_MapValue)(nil),
		(*Object_ByteArrayValue)(nil),
		(*Object_Uint16ArrayValue)(nil),
		(*Object_Uint32ArrayValue)(nil),
		(*Object_Uint64ArrayValue)(nil),
		(*Object_UintptrArrayValue)(nil),
		(*Object_Int8ArrayValue)(nil),
		(*Object_Int16ArrayValue)(nil),
		(*Object_Int32ArrayValue)(nil),
		(*Object_Int64ArrayValue)(nil),
		(*Object_BoolArrayValue)(nil),
		(*Object_Float64ArrayValue)(nil),
		(*Object_Float32ArrayValue)(nil),
	}
}

func init() {
	proto.RegisterType((*Slice)(nil), "gvisor.state.statefile.Slice")
	proto.RegisterType((*Array)(nil), "gvisor.state.statefile.Array")
	proto.RegisterType((*Map)(nil), "gvisor.state.statefile.Map")
	proto.RegisterType((*Interface)(nil), "gvisor.state.statefile.Interface")
	proto.RegisterType((*Struct)(nil), "gvisor.state.statefile.Struct")
	proto.RegisterType((*Field)(nil), "gvisor.state.statefile.Field")
	proto.RegisterType((*Uint16S)(nil), "gvisor.state.statefile.Uint16s")
	proto.RegisterType((*Uint32S)(nil), "gvisor.state.statefile.Uint32s")
	proto.RegisterType((*Uint64S)(nil), "gvisor.state.statefile.Uint64s")
	proto.RegisterType((*Uintptrs)(nil), "gvisor.state.statefile.Uintptrs")
	proto.RegisterType((*Int8S)(nil), "gvisor.state.statefile.Int8s")
	proto.RegisterType((*Int16S)(nil), "gvisor.state.statefile.Int16s")
	proto.RegisterType((*Int32S)(nil), "gvisor.state.statefile.Int32s")
	proto.RegisterType((*Int64S)(nil), "gvisor.state.statefile.Int64s")
	proto.RegisterType((*Bools)(nil), "gvisor.state.statefile.Bools")
	proto.RegisterType((*Float64S)(nil), "gvisor.state.statefile.Float64s")
	proto.RegisterType((*Float32S)(nil), "gvisor.state.statefile.Float32s")
	proto.RegisterType((*Object)(nil), "gvisor.state.statefile.Object")
}

func init() { proto.RegisterFile("pkg/state/object.proto", fileDescriptor_3dee2c1912d4d62d) }

var fileDescriptor_3dee2c1912d4d62d = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0x6f, 0x4f, 0xda, 0x5e,
	0x14, 0xc7, 0xa9, 0x40, 0x29, 0x07, 0x14, 0xb8, 0xfe, 0x7e, 0x8c, 0xcc, 0x38, 0xb1, 0x7b, 0x42,
	0xf6, 0x00, 0x33, 0x60, 0xc4, 0xf8, 0x64, 0x53, 0x13, 0x03, 0xc9, 0x8c, 0x59, 0x8d, 0xcb, 0x9e,
	0x99, 0x52, 0x2f, 0xac, 0xb3, 0xb6, 0x5d, 0x7b, 0x6b, 0xc2, 0xcb, 0xdc, 0x3b, 0x5a, 0xee, 0x1f,
	0xae, 0xfd, 0x03, 0xc5, 0xec, 0x89, 0xa1, 0xb7, 0xdf, 0xf3, 0xe1, 0xdc, 0xf3, 0x3d, 0xe7, 0x08,
	0xb4, 0xfd, 0xc7, 0xc5, 0x49, 0x48, 0x4c, 0x82, 0x4f, 0xbc, 0xd9, 0x2f, 0x6c, 0x91, 0xbe, 0x1f,
	0x78, 0xc4, 0x43, 0xed, 0xc5, 0xb3, 0x1d, 0x7a, 0x41, 0x9f, 0xbd, 0xe2, 0x7f, 0xe7, 0xb6, 0x83,
	0xf5, 0x1f, 0x50, 0xbe, 0x75, 0x6c, 0x0b, 0xa3, 0x36, 0xa8, 0x0e, 0x76, 0x17, 0xe4, 0x67, 0x47,
	0xe9, 0x2a, 0xbd, 0x5d, 0x43, 0x3c, 0xa1, 0xb7, 0xa0, 0x59, 0xa6, 0x6f, 0x5a, 0x36, 0x59, 0x76,
	0x76, 0xd8, 0x1b, 0xf9, 0x8c, 0x0e, 0xa0, 0x1a, 0xe0, 0xf9, 0xfd, 0xb3, 0xe9, 0x44, 0xb8, 0x53,
	0xec, 0x2a, 0xbd, 0x92, 0xa1, 0x05, 0x78, 0xfe, 0x9d, 0x3e, 0xeb, 0x97, 0x50, 0x3e, 0x0f, 0x02,
	0x73, 0x89, 0xce, 0x40, 0xb3, 0x3c, 0x97, 0x60, 0x97, 0x84, 0x1d, 0xa5, 0x5b, 0xec, 0xd5, 0x06,
	0xef, 0xfa, 0xeb, 0xb3, 0xe9, 0xdf, 0xb0, 0x94, 0x0d, 0xa9, 0xd7, 0x7f, 0x43, 0xf1, 0xda, 0xf4,
	0xd1, 0x00, 0x4a, 0x8f, 0x78, 0xf9, 0xda, 0x70, 0xa6, 0x45, 0x63, 0x50, 0x59, 0x62, 0x61, 0x67,
	0xe7, 0x55, 0x51, 0x42, 0xad, 0xdf, 0x41, 0x75, 0xea, 0x12, 0x1c, 0xcc, 0x4d, 0x0b, 0x23, 0x04,
	0x25, 0xb2, 0xf4, 0x31, 0xab, 0x49, 0xd5, 0x60, 0x9f, 0xd1, 0x08, 0xca, 0xfc, 0xc6, 0xb4, 0x1c,
	0xdb, 0xb9, 0x5c, 0xac, 0x7f, 0x06, 0xf5, 0x96, 0x04, 0x91, 0x45, 0xd0, 0x27, 0x50, 0xe7, 0x36,
	0x76, 0x1e, 0x56, 0xd7, 0x39, 0xdc, 0x04, 0xb8, 0xa2, 0x2a, 0x43, 0x88, 0xf5, 0x6f, 0x50, 0x66,
	0x07, 0x34, 0x27, 0xd7, 0x7c, 0x92, 0x39, 0xd1, 0xcf, 0xff, 0x98, 0xd3, 0x31, 0x54, 0xee, 0x6c,
	0x97, 0x7c, 0x1c, 0x87, 0xd4, 0x7e, 0x51, 0x2d, 0x9a, 0xd4, 0xae, 0xac, 0x86, 0x90, 0x0c, 0x07,
	0x69, 0x49, 0x25, 0x2d, 0x19, 0x8f, 0xd2, 0x12, 0x55, 0x4a, 0x74, 0xd0, 0xa8, 0xc4, 0x27, 0xc1,
	0x66, 0xcd, 0x11, 0x94, 0xa7, 0x2e, 0x39, 0x4d, 0x0a, 0x94, 0x5e, 0x5d, 0x0a, 0xba, 0xa0, 0x4e,
	0xd7, 0x25, 0x5b, 0x4e, 0x29, 0xb2, 0xb9, 0x36, 0x52, 0x8a, 0x6c, 0xaa, 0xcd, 0x78, 0x1a, 0x17,
	0x9e, 0xe7, 0xa4, 0x05, 0x5a, 0xfc, 0x2e, 0x57, 0x8e, 0x67, 0xae, 0x81, 0x28, 0x19, 0x4d, 0x36,
	0x95, 0x1d, 0xa9, 0xf9, 0x53, 0x03, 0x95, 0xdb, 0x81, 0x8e, 0x00, 0x66, 0x9e, 0xe7, 0x88, 0x41,
	0xa2, 0xb7, 0xd6, 0x26, 0x05, 0xa3, 0x4a, 0xcf, 0xd8, 0x2c, 0xa1, 0xf7, 0x50, 0x0f, 0x49, 0x60,
	0xbb, 0x8b, 0xfb, 0x17, 0x97, 0xeb, 0x93, 0x82, 0x51, 0xe3, 0xa7, 0x5c, 0x74, 0x0c, 0x35, 0x66,
	0x43, 0x6c, 0x1e, 0x8b, 0x93, 0x82, 0x01, 0xec, 0x50, 0x72, 0xa2, 0xb8, 0xa6, 0x44, 0x67, 0x96,
	0x72, 0xa2, 0xa4, 0xe8, 0xc1, 0x8b, 0x66, 0x0e, 0x16, 0xa2, 0x72, 0x57, 0xe9, 0x29, 0x54, 0xc4,
	0x4f, 0xb9, 0xe8, 0x30, 0x3e, 0xfa, 0xaa, 0xc0, 0xc8, 0xe1, 0x47, 0x5f, 0xa0, 0x16, 0xd2, 0xb5,
	0x22, 0x04, 0x15, 0xd6, 0x95, 0x1b, 0x1b, 0x9d, 0x6d, 0x20, 0x9a, 0x2a, 0x8b, 0x91, 0x04, 0x93,
	0xae, 0x0f, 0x41, 0xd0, 0xf2, 0x09, 0x6c, 0xd3, 0x50, 0x02, 0x8b, 0xe1, 0x84, 0xaf, 0xd0, 0xb0,
	0x57, 0x83, 0x2c, 0x28, 0x55, 0x46, 0x39, 0xde, 0x44, 0x91, 0x73, 0x3f, 0x29, 0x18, 0x7b, 0x32,
	0x96, 0xd3, 0x2e, 0x99, 0x05, 0x91, 0x45, 0x04, 0x0a, 0xf2, 0x07, 0x8d, 0xcf, 0xba, 0xb0, 0x28,
	0xb2, 0x08, 0x87, 0x9c, 0x41, 0xf5, 0xc9, 0xf4, 0x05, 0xa1, 0xc6, 0x08, 0x07, 0x9b, 0x08, 0xd7,
	0xa6, 0x4f, 0x4b, 0xfa, 0x64, 0xfa, 0x3c, 0xf6, 0x03, 0x34, 0x67, 0x4b, 0x82, 0xef, 0xe3, 0x55,
	0xa9, 0x8b, 0x3e, 0xd8, 0xa3, 0x6f, 0xce, 0x5f, 0xae, 0x7e, 0x03, 0x28, 0x62, 0x83, 0x9d, 0x50,
	0xef, 0xb2, 0x2f, 0x3c, 0xda, 0xf4, 0x85, 0x62, 0x15, 0x4c, 0x0a, 0x46, 0x93, 0x07, 0x67, 0x81,
	0xc3, 0x41, 0x02, 0xb8, 0xb7, 0x1d, 0x38, 0x1c, 0x48, 0xe0, 0x70, 0x90, 0x05, 0x8e, 0x47, 0x09,
	0x60, 0x63, 0x3b, 0x70, 0x3c, 0x92, 0xc0, 0xf1, 0x28, 0x06, 0x34, 0x60, 0x3f, 0xe2, 0x2b, 0x26,
	0x41, 0x6c, 0x32, 0x62, 0x37, 0x8f, 0x48, 0xb7, 0xd2, 0xa4, 0x60, 0xb4, 0x44, 0x78, 0x8c, 0x39,
	0x85, 0xa6, 0xed, 0x92, 0xd3, 0x04, 0xb0, 0x95, 0xdf, 0x88, 0x6c, 0x85, 0x89, 0xf6, 0x39, 0x3d,
	0x8f, 0x37, 0x63, 0x2b, 0x6b, 0x08, 0xca, 0xef, 0xa1, 0xe9, 0xca, 0x8f, 0x46, 0xda, 0x0e, 0x4e,
	0x4b, 0xb9, 0xb1, 0xbf, 0x95, 0xc6, 0xcd, 0x68, 0xa4, 0xbd, 0xe0, 0xb4, 0x94, 0x15, 0xff, 0x6d,
	0xa5, 0x71, 0x27, 0x1a, 0x69, 0x23, 0xa6, 0xd0, 0x64, 0xcb, 0x2c, 0x0e, 0xfb, 0x3f, 0xbf, 0x68,
	0x6c, 0xe1, 0xb2, 0x36, 0xf6, 0x3c, 0x27, 0xe9, 0xe9, 0x9c, 0xaf, 0xda, 0x04, 0xad, 0x9d, 0xef,
	0xe9, 0x6a, 0x3b, 0x53, 0x4f, 0x45, 0xf8, 0x1a, 0x66, 0xaa, 0x78, 0x6f, 0x5e, 0xc1, 0xe4, 0xe5,
	0x6b, 0x89, 0xf0, 0x17, 0xe6, 0x45, 0x45, 0xfc, 0xf7, 0x9d, 0xa9, 0xec, 0xc7, 0xd6, 0xf0, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x69, 0xc9, 0x45, 0x86, 0x09, 0x00, 0x00,
}

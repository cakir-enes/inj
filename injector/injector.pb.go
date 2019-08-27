// Code generated by protoc-gen-go. DO NOT EDIT.
// source: injector.proto

package injector

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type Parameter_Type int32

const (
	Parameter_UNKNOWN Parameter_Type = 0
	Parameter_STR     Parameter_Type = 1
	Parameter_INT     Parameter_Type = 2
	Parameter_FLOAT   Parameter_Type = 3
	Parameter_BOOL    Parameter_Type = 4
	Parameter_ARR     Parameter_Type = 5
	Parameter_ENUM    Parameter_Type = 6
)

var Parameter_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "STR",
	2: "INT",
	3: "FLOAT",
	4: "BOOL",
	5: "ARR",
	6: "ENUM",
}

var Parameter_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"STR":     1,
	"INT":     2,
	"FLOAT":   3,
	"BOOL":    4,
	"ARR":     5,
	"ENUM":    6,
}

func (x Parameter_Type) String() string {
	return proto.EnumName(Parameter_Type_name, int32(x))
}

func (Parameter_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{6, 0}
}

type Status_Type int32

const (
	Status_Unknown      Status_Type = 0
	Status_Success      Status_Type = 1
	Status_InvalidPath  Status_Type = 2
	Status_InvalidValue Status_Type = 3
)

var Status_Type_name = map[int32]string{
	0: "Unknown",
	1: "Success",
	2: "InvalidPath",
	3: "InvalidValue",
}

var Status_Type_value = map[string]int32{
	"Unknown":      0,
	"Success":      1,
	"InvalidPath":  2,
	"InvalidValue": 3,
}

func (x Status_Type) String() string {
	return proto.EnumName(Status_Type_name, int32(x))
}

func (Status_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{7, 0}
}

type MultiSetReq struct {
	PathToVal            map[string]string `protobuf:"bytes,1,rep,name=pathToVal,proto3" json:"pathToVal,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MultiSetReq) Reset()         { *m = MultiSetReq{} }
func (m *MultiSetReq) String() string { return proto.CompactTextString(m) }
func (*MultiSetReq) ProtoMessage()    {}
func (*MultiSetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{0}
}

func (m *MultiSetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiSetReq.Unmarshal(m, b)
}
func (m *MultiSetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiSetReq.Marshal(b, m, deterministic)
}
func (m *MultiSetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSetReq.Merge(m, src)
}
func (m *MultiSetReq) XXX_Size() int {
	return xxx_messageInfo_MultiSetReq.Size(m)
}
func (m *MultiSetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSetReq.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSetReq proto.InternalMessageInfo

func (m *MultiSetReq) GetPathToVal() map[string]string {
	if m != nil {
		return m.PathToVal
	}
	return nil
}

type MultiSetResp struct {
	Statuses             []*Status `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MultiSetResp) Reset()         { *m = MultiSetResp{} }
func (m *MultiSetResp) String() string { return proto.CompactTextString(m) }
func (*MultiSetResp) ProtoMessage()    {}
func (*MultiSetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{1}
}

func (m *MultiSetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiSetResp.Unmarshal(m, b)
}
func (m *MultiSetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiSetResp.Marshal(b, m, deterministic)
}
func (m *MultiSetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSetResp.Merge(m, src)
}
func (m *MultiSetResp) XXX_Size() int {
	return xxx_messageInfo_MultiSetResp.Size(m)
}
func (m *MultiSetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSetResp.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSetResp proto.InternalMessageInfo

func (m *MultiSetResp) GetStatuses() []*Status {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type MultiGetReq struct {
	ParamPaths           []string `protobuf:"bytes,1,rep,name=paramPaths,proto3" json:"paramPaths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiGetReq) Reset()         { *m = MultiGetReq{} }
func (m *MultiGetReq) String() string { return proto.CompactTextString(m) }
func (*MultiGetReq) ProtoMessage()    {}
func (*MultiGetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{2}
}

func (m *MultiGetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiGetReq.Unmarshal(m, b)
}
func (m *MultiGetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiGetReq.Marshal(b, m, deterministic)
}
func (m *MultiGetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiGetReq.Merge(m, src)
}
func (m *MultiGetReq) XXX_Size() int {
	return xxx_messageInfo_MultiGetReq.Size(m)
}
func (m *MultiGetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiGetReq.DiscardUnknown(m)
}

var xxx_messageInfo_MultiGetReq proto.InternalMessageInfo

func (m *MultiGetReq) GetParamPaths() []string {
	if m != nil {
		return m.ParamPaths
	}
	return nil
}

type MultiGetResp struct {
	Pairs                []*PathAndValue `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MultiGetResp) Reset()         { *m = MultiGetResp{} }
func (m *MultiGetResp) String() string { return proto.CompactTextString(m) }
func (*MultiGetResp) ProtoMessage()    {}
func (*MultiGetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{3}
}

func (m *MultiGetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiGetResp.Unmarshal(m, b)
}
func (m *MultiGetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiGetResp.Marshal(b, m, deterministic)
}
func (m *MultiGetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiGetResp.Merge(m, src)
}
func (m *MultiGetResp) XXX_Size() int {
	return xxx_messageInfo_MultiGetResp.Size(m)
}
func (m *MultiGetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiGetResp.DiscardUnknown(m)
}

var xxx_messageInfo_MultiGetResp proto.InternalMessageInfo

func (m *MultiGetResp) GetPairs() []*PathAndValue {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type ParameterInfoResp struct {
	ParamInfos           []*Parameter `protobuf:"bytes,1,rep,name=paramInfos,proto3" json:"paramInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ParameterInfoResp) Reset()         { *m = ParameterInfoResp{} }
func (m *ParameterInfoResp) String() string { return proto.CompactTextString(m) }
func (*ParameterInfoResp) ProtoMessage()    {}
func (*ParameterInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{4}
}

func (m *ParameterInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterInfoResp.Unmarshal(m, b)
}
func (m *ParameterInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterInfoResp.Marshal(b, m, deterministic)
}
func (m *ParameterInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterInfoResp.Merge(m, src)
}
func (m *ParameterInfoResp) XXX_Size() int {
	return xxx_messageInfo_ParameterInfoResp.Size(m)
}
func (m *ParameterInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterInfoResp proto.InternalMessageInfo

func (m *ParameterInfoResp) GetParamInfos() []*Parameter {
	if m != nil {
		return m.ParamInfos
	}
	return nil
}

type EnumValuesResp struct {
	Values               map[string]*EnumValuesResp_Values `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *EnumValuesResp) Reset()         { *m = EnumValuesResp{} }
func (m *EnumValuesResp) String() string { return proto.CompactTextString(m) }
func (*EnumValuesResp) ProtoMessage()    {}
func (*EnumValuesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{5}
}

func (m *EnumValuesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValuesResp.Unmarshal(m, b)
}
func (m *EnumValuesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValuesResp.Marshal(b, m, deterministic)
}
func (m *EnumValuesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValuesResp.Merge(m, src)
}
func (m *EnumValuesResp) XXX_Size() int {
	return xxx_messageInfo_EnumValuesResp.Size(m)
}
func (m *EnumValuesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValuesResp.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValuesResp proto.InternalMessageInfo

func (m *EnumValuesResp) GetValues() map[string]*EnumValuesResp_Values {
	if m != nil {
		return m.Values
	}
	return nil
}

type EnumValuesResp_Values struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumValuesResp_Values) Reset()         { *m = EnumValuesResp_Values{} }
func (m *EnumValuesResp_Values) String() string { return proto.CompactTextString(m) }
func (*EnumValuesResp_Values) ProtoMessage()    {}
func (*EnumValuesResp_Values) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{5, 0}
}

func (m *EnumValuesResp_Values) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValuesResp_Values.Unmarshal(m, b)
}
func (m *EnumValuesResp_Values) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValuesResp_Values.Marshal(b, m, deterministic)
}
func (m *EnumValuesResp_Values) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValuesResp_Values.Merge(m, src)
}
func (m *EnumValuesResp_Values) XXX_Size() int {
	return xxx_messageInfo_EnumValuesResp_Values.Size(m)
}
func (m *EnumValuesResp_Values) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValuesResp_Values.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValuesResp_Values proto.InternalMessageInfo

func (m *EnumValuesResp_Values) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Parameter struct {
	Path                 string         `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Type                 Parameter_Type `protobuf:"varint,2,opt,name=type,proto3,enum=Parameter_Type" json:"type,omitempty"`
	Value                string         `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{6}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Parameter) GetType() Parameter_Type {
	if m != nil {
		return m.Type
	}
	return Parameter_UNKNOWN
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Status struct {
	Path                 string      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Detail               string      `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Type                 Status_Type `protobuf:"varint,3,opt,name=type,proto3,enum=Status_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{7}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Status) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Status) GetType() Status_Type {
	if m != nil {
		return m.Type
	}
	return Status_Unknown
}

type PathAndValue struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathAndValue) Reset()         { *m = PathAndValue{} }
func (m *PathAndValue) String() string { return proto.CompactTextString(m) }
func (*PathAndValue) ProtoMessage()    {}
func (*PathAndValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde8fb0ee3975dcf, []int{8}
}

func (m *PathAndValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathAndValue.Unmarshal(m, b)
}
func (m *PathAndValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathAndValue.Marshal(b, m, deterministic)
}
func (m *PathAndValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathAndValue.Merge(m, src)
}
func (m *PathAndValue) XXX_Size() int {
	return xxx_messageInfo_PathAndValue.Size(m)
}
func (m *PathAndValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PathAndValue.DiscardUnknown(m)
}

var xxx_messageInfo_PathAndValue proto.InternalMessageInfo

func (m *PathAndValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathAndValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("Parameter_Type", Parameter_Type_name, Parameter_Type_value)
	proto.RegisterEnum("Status_Type", Status_Type_name, Status_Type_value)
	proto.RegisterType((*MultiSetReq)(nil), "MultiSetReq")
	proto.RegisterMapType((map[string]string)(nil), "MultiSetReq.PathToValEntry")
	proto.RegisterType((*MultiSetResp)(nil), "MultiSetResp")
	proto.RegisterType((*MultiGetReq)(nil), "MultiGetReq")
	proto.RegisterType((*MultiGetResp)(nil), "MultiGetResp")
	proto.RegisterType((*ParameterInfoResp)(nil), "ParameterInfoResp")
	proto.RegisterType((*EnumValuesResp)(nil), "EnumValuesResp")
	proto.RegisterMapType((map[string]*EnumValuesResp_Values)(nil), "EnumValuesResp.ValuesEntry")
	proto.RegisterType((*EnumValuesResp_Values)(nil), "EnumValuesResp.Values")
	proto.RegisterType((*Parameter)(nil), "Parameter")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*PathAndValue)(nil), "PathAndValue")
}

func init() { proto.RegisterFile("injector.proto", fileDescriptor_dde8fb0ee3975dcf) }

var fileDescriptor_dde8fb0ee3975dcf = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0x66, 0x58, 0xa0, 0x70, 0xa0, 0x74, 0x9c, 0x34, 0xa4, 0xa1, 0x89, 0x21, 0xd3, 0x9b, 0xfa,
	0x37, 0x35, 0xf4, 0xa6, 0x1a, 0x8d, 0xa1, 0x8a, 0x84, 0xd8, 0x02, 0x2e, 0xb4, 0x5e, 0x6f, 0xe9,
	0xb4, 0x62, 0x97, 0xdd, 0x75, 0x77, 0xa8, 0xe1, 0x05, 0x7c, 0x09, 0x9f, 0xc0, 0x57, 0xf0, 0x85,
	0x7c, 0x0d, 0x33, 0x3f, 0xbb, 0x0c, 0x8a, 0x77, 0x67, 0xce, 0xf9, 0xce, 0x39, 0xdf, 0xf9, 0x1b,
	0xa8, 0xcf, 0x82, 0x2f, 0x7c, 0x2a, 0xc2, 0x98, 0x45, 0x71, 0x28, 0xc2, 0xe6, 0xfe, 0x6d, 0x18,
	0xde, 0xfa, 0xfc, 0x48, 0xbd, 0xae, 0x16, 0x37, 0x47, 0x7c, 0x1e, 0x89, 0xa5, 0x36, 0xd2, 0xef,
	0x08, 0xaa, 0xe7, 0x0b, 0x5f, 0xcc, 0xc6, 0x5c, 0xb8, 0xfc, 0x2b, 0x79, 0x01, 0x95, 0xc8, 0x13,
	0x9f, 0x27, 0xe1, 0xa5, 0xe7, 0xef, 0xa1, 0x96, 0x73, 0x58, 0x6d, 0xef, 0x33, 0x0b, 0xc0, 0x46,
	0xa9, 0xb5, 0x1b, 0x88, 0x78, 0xe9, 0xae, 0xd0, 0xcd, 0x57, 0x50, 0x5f, 0x37, 0x12, 0x0c, 0xce,
	0x1d, 0x5f, 0xee, 0xa1, 0x16, 0x3a, 0xac, 0xb8, 0x52, 0x24, 0xbb, 0x50, 0xbc, 0xf7, 0xfc, 0x05,
	0xdf, 0xcb, 0x2b, 0x9d, 0x7e, 0xbc, 0xcc, 0x9f, 0x20, 0x7a, 0x0c, 0xb5, 0x55, 0x9a, 0x24, 0x22,
	0x07, 0x50, 0x4e, 0x84, 0x27, 0x16, 0x09, 0x4f, 0x0c, 0x8f, 0x2d, 0x36, 0x56, 0x0a, 0x37, 0x33,
	0xd0, 0x67, 0x86, 0x7c, 0x4f, 0x93, 0x7f, 0x08, 0x10, 0x79, 0xb1, 0x37, 0x97, 0x34, 0xb4, 0x57,
	0xc5, 0xb5, 0x34, 0x59, 0x8e, 0x5e, 0x96, 0xa3, 0x18, 0x79, 0xb3, 0x38, 0x4d, 0xb0, 0xad, 0x8a,
	0xeb, 0x04, 0xd7, 0x97, 0x92, 0x95, 0xab, 0x6d, 0xf4, 0x0d, 0x3c, 0x18, 0xc9, 0x10, 0x5c, 0xf0,
	0xb8, 0x1f, 0xdc, 0x84, 0xca, 0xf3, 0xb1, 0xc9, 0x24, 0x15, 0xa9, 0x3b, 0xb0, 0x0c, 0xe7, 0x5a,
	0x56, 0xfa, 0x0b, 0x41, 0xbd, 0x1b, 0x2c, 0xe6, 0x2a, 0x6a, 0xa2, 0xdc, 0x8f, 0xa1, 0xa4, 0x2a,
	0x4f, 0xb2, 0x16, 0xaf, 0x03, 0x98, 0x16, 0x75, 0x8b, 0x0d, 0xb4, 0xd9, 0x82, 0x92, 0x56, 0x93,
	0xc6, 0x9a, 0x7b, 0x25, 0x43, 0x7c, 0x84, 0xaa, 0xe5, 0xb8, 0xa1, 0xfd, 0x4f, 0xed, 0xf6, 0x57,
	0xdb, 0x8d, 0xcd, 0x69, 0xed, 0xb1, 0xfc, 0x44, 0x50, 0xc9, 0xca, 0x22, 0x04, 0x0a, 0x72, 0xde,
	0x26, 0xa4, 0x92, 0xc9, 0x01, 0x14, 0xc4, 0x32, 0xd2, 0x21, 0xeb, 0xed, 0x9d, 0x55, 0x13, 0xd8,
	0x64, 0x19, 0x71, 0x57, 0x19, 0x57, 0x73, 0x77, 0xac, 0xb9, 0xd3, 0x73, 0x28, 0x48, 0x0c, 0xa9,
	0xc2, 0xd6, 0xc5, 0xe0, 0xc3, 0x60, 0xf8, 0x69, 0x80, 0x73, 0x64, 0x0b, 0x9c, 0xf1, 0xc4, 0xc5,
	0x48, 0x0a, 0xfd, 0xc1, 0x04, 0xe7, 0x49, 0x05, 0x8a, 0xef, 0xcf, 0x86, 0x9d, 0x09, 0x76, 0x48,
	0x19, 0x0a, 0xa7, 0xc3, 0xe1, 0x19, 0x2e, 0x48, 0x6b, 0xc7, 0x75, 0x71, 0x51, 0xaa, 0xba, 0x83,
	0x8b, 0x73, 0x5c, 0xa2, 0x3f, 0x10, 0x94, 0xf4, 0x8a, 0x6c, 0x24, 0xda, 0x80, 0xd2, 0x35, 0x17,
	0xde, 0xcc, 0x37, 0xcb, 0x67, 0x5e, 0xa4, 0x65, 0x0a, 0x70, 0x54, 0x01, 0x35, 0xb3, 0x65, 0x16,
	0x7b, 0xfa, 0xd6, 0xe2, 0x19, 0xdc, 0x05, 0xe1, 0xb7, 0x00, 0xe7, 0xe4, 0x63, 0xbc, 0x98, 0x4e,
	0x79, 0x92, 0x60, 0x44, 0x76, 0xa0, 0xda, 0x0f, 0xee, 0x3d, 0x7f, 0x76, 0x2d, 0x57, 0x08, 0xe7,
	0x09, 0x86, 0x9a, 0x51, 0xa8, 0x9e, 0x62, 0x87, 0x9e, 0x40, 0xcd, 0x5e, 0xaf, 0x8d, 0x14, 0x37,
	0x9e, 0x47, 0xfb, 0x37, 0x82, 0x72, 0xdf, 0xdc, 0x34, 0x79, 0x04, 0xe5, 0xf4, 0x4e, 0x48, 0xcd,
	0xbe, 0xcc, 0xe6, 0x36, 0xb3, 0x0f, 0x88, 0xe6, 0xc8, 0x13, 0x03, 0xed, 0xad, 0xa0, 0xbd, 0x35,
	0x68, 0x2f, 0x85, 0x3e, 0x47, 0xe4, 0x1d, 0xec, 0xf6, 0xb8, 0xe8, 0xf8, 0xfe, 0xda, 0xb2, 0x27,
	0xa4, 0xc1, 0xf4, 0xf7, 0xc1, 0xd2, 0xef, 0x83, 0x75, 0xe5, 0xf7, 0xd1, 0x24, 0xec, 0x9f, 0xab,
	0xa0, 0x39, 0xf2, 0x1a, 0xb0, 0x8e, 0xb2, 0x5a, 0xac, 0xff, 0x46, 0xd8, 0xf9, 0x6b, 0xfb, 0x68,
	0xee, 0xf4, 0x00, 0xf0, 0x2c, 0x64, 0xb7, 0x71, 0x34, 0x65, 0x69, 0xc1, 0xa7, 0xdb, 0xa9, 0x34,
	0x92, 0xee, 0x23, 0x74, 0x55, 0x52, 0x71, 0x8e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x7a,
	0xae, 0x7e, 0xe8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InjectorClient is the client API for Injector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InjectorClient interface {
	MultiSet(ctx context.Context, in *MultiSetReq, opts ...grpc.CallOption) (*MultiSetResp, error)
	MultiGet(ctx context.Context, in *MultiGetReq, opts ...grpc.CallOption) (Injector_MultiGetClient, error)
	GetAllParameterInfos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParameterInfoResp, error)
	GetAllEnumValues(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*EnumValuesResp, error)
}

type injectorClient struct {
	cc *grpc.ClientConn
}

func NewInjectorClient(cc *grpc.ClientConn) InjectorClient {
	return &injectorClient{cc}
}

func (c *injectorClient) MultiSet(ctx context.Context, in *MultiSetReq, opts ...grpc.CallOption) (*MultiSetResp, error) {
	out := new(MultiSetResp)
	err := c.cc.Invoke(ctx, "/Injector/MultiSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectorClient) MultiGet(ctx context.Context, in *MultiGetReq, opts ...grpc.CallOption) (Injector_MultiGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Injector_serviceDesc.Streams[0], "/Injector/MultiGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectorMultiGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Injector_MultiGetClient interface {
	Recv() (*MultiGetResp, error)
	grpc.ClientStream
}

type injectorMultiGetClient struct {
	grpc.ClientStream
}

func (x *injectorMultiGetClient) Recv() (*MultiGetResp, error) {
	m := new(MultiGetResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectorClient) GetAllParameterInfos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParameterInfoResp, error) {
	out := new(ParameterInfoResp)
	err := c.cc.Invoke(ctx, "/Injector/GetAllParameterInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectorClient) GetAllEnumValues(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*EnumValuesResp, error) {
	out := new(EnumValuesResp)
	err := c.cc.Invoke(ctx, "/Injector/GetAllEnumValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectorServer is the server API for Injector service.
type InjectorServer interface {
	MultiSet(context.Context, *MultiSetReq) (*MultiSetResp, error)
	MultiGet(*MultiGetReq, Injector_MultiGetServer) error
	GetAllParameterInfos(context.Context, *empty.Empty) (*ParameterInfoResp, error)
	GetAllEnumValues(context.Context, *empty.Empty) (*EnumValuesResp, error)
}

func RegisterInjectorServer(s *grpc.Server, srv InjectorServer) {
	s.RegisterService(&_Injector_serviceDesc, srv)
}

func _Injector_MultiSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiSetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectorServer).MultiSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Injector/MultiSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectorServer).MultiSet(ctx, req.(*MultiSetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injector_MultiGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MultiGetReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectorServer).MultiGet(m, &injectorMultiGetServer{stream})
}

type Injector_MultiGetServer interface {
	Send(*MultiGetResp) error
	grpc.ServerStream
}

type injectorMultiGetServer struct {
	grpc.ServerStream
}

func (x *injectorMultiGetServer) Send(m *MultiGetResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Injector_GetAllParameterInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectorServer).GetAllParameterInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Injector/GetAllParameterInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectorServer).GetAllParameterInfos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injector_GetAllEnumValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectorServer).GetAllEnumValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Injector/GetAllEnumValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectorServer).GetAllEnumValues(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Injector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Injector",
	HandlerType: (*InjectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MultiSet",
			Handler:    _Injector_MultiSet_Handler,
		},
		{
			MethodName: "GetAllParameterInfos",
			Handler:    _Injector_GetAllParameterInfos_Handler,
		},
		{
			MethodName: "GetAllEnumValues",
			Handler:    _Injector_GetAllEnumValues_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MultiGet",
			Handler:       _Injector_MultiGet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "injector.proto",
}
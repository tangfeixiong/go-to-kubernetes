// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/message.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CrdRecipient_ResourceScope int32

const (
	CrdRecipient_Cluster    CrdRecipient_ResourceScope = 0
	CrdRecipient_Namespaced CrdRecipient_ResourceScope = 1
)

var CrdRecipient_ResourceScope_name = map[int32]string{
	0: "Cluster",
	1: "Namespaced",
}
var CrdRecipient_ResourceScope_value = map[string]int32{
	"Cluster":    0,
	"Namespaced": 1,
}

func (x CrdRecipient_ResourceScope) String() string {
	return proto.EnumName(CrdRecipient_ResourceScope_name, int32(x))
}
func (CrdRecipient_ResourceScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorMessage, []int{0, 0}
}

type CrdRecipient struct {
	Name     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group    string                     `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Version  string                     `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Scope    CrdRecipient_ResourceScope `protobuf:"varint,4,opt,name=scope,proto3,enum=pb.CrdRecipient_ResourceScope" json:"scope,omitempty"`
	Plural   string                     `protobuf:"bytes,5,opt,name=plural,proto3" json:"plural,omitempty"`
	Singular string                     `protobuf:"bytes,6,opt,name=singular,proto3" json:"singular,omitempty"`
	Kind     string                     `protobuf:"bytes,7,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (m *CrdRecipient) Reset()                    { *m = CrdRecipient{} }
func (m *CrdRecipient) String() string            { return proto.CompactTextString(m) }
func (*CrdRecipient) ProtoMessage()               {}
func (*CrdRecipient) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

func (m *CrdRecipient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CrdRecipient) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *CrdRecipient) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CrdRecipient) GetScope() CrdRecipient_ResourceScope {
	if m != nil {
		return m.Scope
	}
	return CrdRecipient_Cluster
}

func (m *CrdRecipient) GetPlural() string {
	if m != nil {
		return m.Plural
	}
	return ""
}

func (m *CrdRecipient) GetSingular() string {
	if m != nil {
		return m.Singular
	}
	return ""
}

func (m *CrdRecipient) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*CrdRecipient)(nil), "pb.CrdRecipient")
	proto.RegisterEnum("pb.CrdRecipient_ResourceScope", CrdRecipient_ResourceScope_name, CrdRecipient_ResourceScope_value)
}

func init() { proto.RegisterFile("pb/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xcd, 0xba, 0x6d, 0x75, 0xd4, 0xa5, 0x0c, 0x22, 0xc1, 0x83, 0x2c, 0x7b, 0xda, 0x83,
	0x54, 0x50, 0xff, 0x41, 0xef, 0x1e, 0xea, 0x2f, 0x48, 0xdb, 0xa1, 0x04, 0xdb, 0x64, 0xc8, 0x34,
	0xfe, 0x7c, 0x91, 0xa6, 0x2a, 0xee, 0xed, 0x7d, 0x2f, 0x8f, 0x97, 0xc7, 0x40, 0xc9, 0xed, 0xd3,
	0x44, 0x22, 0x66, 0xa0, 0x8a, 0x83, 0x9f, 0x3d, 0x6e, 0xb8, 0x3d, 0x7c, 0x29, 0xb8, 0xae, 0x43,
	0xdf, 0x50, 0x67, 0xd9, 0x92, 0x9b, 0x11, 0x61, 0xeb, 0xcc, 0x44, 0x5a, 0xed, 0xd5, 0xf1, 0xb2,
	0x49, 0x1a, 0x6f, 0x21, 0x1b, 0x82, 0x8f, 0xac, 0x37, 0xc9, 0x5c, 0x01, 0x35, 0x14, 0x9f, 0x14,
	0xc4, 0x7a, 0xa7, 0xcf, 0x93, 0xff, 0x8b, 0xf8, 0x0a, 0x99, 0x74, 0x9e, 0x49, 0x6f, 0xf7, 0xea,
	0xb8, 0x7b, 0x7e, 0xa8, 0xb8, 0xad, 0xfe, 0x7f, 0x52, 0x35, 0x24, 0x3e, 0x86, 0x8e, 0xde, 0x97,
	0x54, 0xb3, 0x86, 0xf1, 0x0e, 0x72, 0x1e, 0x63, 0x30, 0xa3, 0xce, 0x52, 0xdd, 0x0f, 0xe1, 0x3d,
	0x5c, 0x88, 0x75, 0x43, 0x1c, 0x4d, 0xd0, 0x79, 0x7a, 0xf9, 0xe3, 0x65, 0xed, 0x87, 0x75, 0xbd,
	0x2e, 0xd6, 0xb5, 0x8b, 0x3e, 0x3c, 0xc2, 0xcd, 0x49, 0x3f, 0x5e, 0x41, 0x51, 0x8f, 0x51, 0x66,
	0x0a, 0xe5, 0x19, 0xee, 0x00, 0xde, 0xcc, 0x44, 0xc2, 0xa6, 0xa3, 0xbe, 0x54, 0x6d, 0x9e, 0x6e,
	0xf1, 0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xc0, 0x05, 0xf0, 0x1f, 0x01, 0x00, 0x00,
}

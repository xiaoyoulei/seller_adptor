// Code generated by protoc-gen-go.
// source: iptable.proto
// DO NOT EDIT!

/*
Package jesgoo_protocol is a generated protocol buffer package.

It is generated from these files:
	iptable.proto

It has these top-level messages:
	IPTable
*/
package jesgoo_protocol

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type IPTable struct {
	Classa           []*IPTable_ClassA `protobuf:"group,1,rep,name=ClassA" json:"classa,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *IPTable) Reset()         { *m = IPTable{} }
func (m *IPTable) String() string { return proto.CompactTextString(m) }
func (*IPTable) ProtoMessage()    {}

func (m *IPTable) GetClassa() []*IPTable_ClassA {
	if m != nil {
		return m.Classa
	}
	return nil
}

type IPTable_IPSection struct {
	Start            *uint32 `protobuf:"varint,1,req,name=start" json:"start,omitempty"`
	End              *uint32 `protobuf:"varint,2,req,name=end" json:"end,omitempty"`
	Country          *uint32 `protobuf:"varint,3,req,name=country" json:"country,omitempty"`
	Province         *uint32 `protobuf:"varint,4,opt,name=province,def=0" json:"province,omitempty"`
	City             *uint32 `protobuf:"varint,5,opt,name=city,def=0" json:"city,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IPTable_IPSection) Reset()         { *m = IPTable_IPSection{} }
func (m *IPTable_IPSection) String() string { return proto.CompactTextString(m) }
func (*IPTable_IPSection) ProtoMessage()    {}

const Default_IPTable_IPSection_Province uint32 = 0
const Default_IPTable_IPSection_City uint32 = 0

func (m *IPTable_IPSection) GetStart() uint32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *IPTable_IPSection) GetEnd() uint32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

func (m *IPTable_IPSection) GetCountry() uint32 {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return 0
}

func (m *IPTable_IPSection) GetProvince() uint32 {
	if m != nil && m.Province != nil {
		return *m.Province
	}
	return Default_IPTable_IPSection_Province
}

func (m *IPTable_IPSection) GetCity() uint32 {
	if m != nil && m.City != nil {
		return *m.City
	}
	return Default_IPTable_IPSection_City
}

type IPTable_ClassA struct {
	IpSections       []*IPTable_IPSection `protobuf:"bytes,2,rep,name=ip_sections" json:"ip_sections,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *IPTable_ClassA) Reset()         { *m = IPTable_ClassA{} }
func (m *IPTable_ClassA) String() string { return proto.CompactTextString(m) }
func (*IPTable_ClassA) ProtoMessage()    {}

func (m *IPTable_ClassA) GetIpSections() []*IPTable_IPSection {
	if m != nil {
		return m.IpSections
	}
	return nil
}

func init() {
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ubkoin/ubkoin.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Protocol int32

const (
	Protocol_PL_UNKNOWN Protocol = 0
	Protocol_PL_EMAIL   Protocol = 1
	Protocol_PL_GENERIC Protocol = 100
)

var Protocol_name = map[int32]string{
	0:   "PL_UNKNOWN",
	1:   "PL_EMAIL",
	100: "PL_GENERIC",
}

var Protocol_value = map[string]int32{
	"PL_UNKNOWN": 0,
	"PL_EMAIL":   1,
	"PL_GENERIC": 100,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{0}
}

type Key struct {
	RegistrationTimestamp int64 `protobuf:"varint,1,opt,name=registration_timestamp,json=registrationTimestamp,proto3" json:"registration_timestamp,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetRegistrationTimestamp() int64 {
	if m != nil {
		return m.RegistrationTimestamp
	}
	return 0
}

type RegisterKey struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *RegisterKey) Reset()         { *m = RegisterKey{} }
func (m *RegisterKey) String() string { return proto.CompactTextString(m) }
func (*RegisterKey) ProtoMessage()    {}
func (*RegisterKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{1}
}
func (m *RegisterKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterKey.Merge(m, src)
}
func (m *RegisterKey) XXX_Size() int {
	return m.Size()
}
func (m *RegisterKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterKey.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterKey proto.InternalMessageInfo

func (m *RegisterKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type EncryptionKey struct {
	Key        []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValidUntil int64  `protobuf:"varint,2,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
}

func (m *EncryptionKey) Reset()         { *m = EncryptionKey{} }
func (m *EncryptionKey) String() string { return proto.CompactTextString(m) }
func (*EncryptionKey) ProtoMessage()    {}
func (*EncryptionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{2}
}
func (m *EncryptionKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncryptionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncryptionKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncryptionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionKey.Merge(m, src)
}
func (m *EncryptionKey) XXX_Size() int {
	return m.Size()
}
func (m *EncryptionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionKey.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionKey proto.InternalMessageInfo

func (m *EncryptionKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EncryptionKey) GetValidUntil() int64 {
	if m != nil {
		return m.ValidUntil
	}
	return 0
}

type Endpoint struct {
	Protocol Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=regnull.ubkoin.ubkoin.Protocol" json:"protocol,omitempty"`
	Address  string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{3}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return m.Size()
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_PL_UNKNOWN
}

func (m *Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type NameRecord struct {
	Name          string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnerKey      []byte           `protobuf:"bytes,2,opt,name=owner_key,json=ownerKey,proto3" json:"owner_key,omitempty"`
	Endpoint      []*Endpoint      `protobuf:"bytes,3,rep,name=endpoint,proto3" json:"endpoint,omitempty"`
	EncryptionKey []*EncryptionKey `protobuf:"bytes,4,rep,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
}

func (m *NameRecord) Reset()         { *m = NameRecord{} }
func (m *NameRecord) String() string { return proto.CompactTextString(m) }
func (*NameRecord) ProtoMessage()    {}
func (*NameRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{4}
}
func (m *NameRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRecord.Merge(m, src)
}
func (m *NameRecord) XXX_Size() int {
	return m.Size()
}
func (m *NameRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRecord.DiscardUnknown(m)
}

var xxx_messageInfo_NameRecord proto.InternalMessageInfo

func (m *NameRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameRecord) GetOwnerKey() []byte {
	if m != nil {
		return m.OwnerKey
	}
	return nil
}

func (m *NameRecord) GetEndpoint() []*Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *NameRecord) GetEncryptionKey() []*EncryptionKey {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

type ReserveName struct {
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OwnerAddress string `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *ReserveName) Reset()         { *m = ReserveName{} }
func (m *ReserveName) String() string { return proto.CompactTextString(m) }
func (*ReserveName) ProtoMessage()    {}
func (*ReserveName) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb7cebd4a6143a1, []int{5}
}
func (m *ReserveName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReserveName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReserveName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReserveName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveName.Merge(m, src)
}
func (m *ReserveName) XXX_Size() int {
	return m.Size()
}
func (m *ReserveName) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveName.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveName proto.InternalMessageInfo

func (m *ReserveName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReserveName) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("regnull.ubkoin.ubkoin.Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*Key)(nil), "regnull.ubkoin.ubkoin.Key")
	proto.RegisterType((*RegisterKey)(nil), "regnull.ubkoin.ubkoin.RegisterKey")
	proto.RegisterType((*EncryptionKey)(nil), "regnull.ubkoin.ubkoin.EncryptionKey")
	proto.RegisterType((*Endpoint)(nil), "regnull.ubkoin.ubkoin.Endpoint")
	proto.RegisterType((*NameRecord)(nil), "regnull.ubkoin.ubkoin.NameRecord")
	proto.RegisterType((*ReserveName)(nil), "regnull.ubkoin.ubkoin.ReserveName")
}

func init() { proto.RegisterFile("ubkoin/ubkoin.proto", fileDescriptor_0cb7cebd4a6143a1) }

var fileDescriptor_0cb7cebd4a6143a1 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x6e, 0x36, 0x8b, 0xa6, 0xaf, 0x3f, 0x28, 0xa3, 0x2b, 0x45, 0x21, 0x5d, 0xa2, 0x87, 0xe2,
	0x21, 0x05, 0x45, 0x10, 0xd6, 0xcb, 0x56, 0xa2, 0x2c, 0xad, 0xb1, 0x0c, 0x2e, 0x82, 0x97, 0x90,
	0x36, 0x8f, 0x18, 0x36, 0x99, 0x09, 0x93, 0xe9, 0x6a, 0xfe, 0x0b, 0xff, 0x2b, 0x3d, 0xee, 0xd1,
	0xa3, 0xb4, 0xff, 0x88, 0xcc, 0x24, 0x29, 0x5d, 0x68, 0x4f, 0xf3, 0xe6, 0x7b, 0xef, 0x7d, 0xef,
	0xfb, 0xde, 0x0c, 0x3c, 0x5a, 0x2f, 0x6f, 0x78, 0xc2, 0x26, 0xd5, 0xe1, 0xe6, 0x82, 0x4b, 0x4e,
	0xce, 0x04, 0xc6, 0x6c, 0x9d, 0xa6, 0x6e, 0x8d, 0x56, 0xc7, 0xd3, 0xc7, 0x31, 0x8f, 0xb9, 0xae,
	0x98, 0xa8, 0xa8, 0x2a, 0x76, 0xde, 0x81, 0x39, 0xc3, 0x92, 0xbc, 0x81, 0x27, 0x02, 0xe3, 0xa4,
	0x90, 0x22, 0x94, 0x09, 0x67, 0x81, 0x4c, 0x32, 0x2c, 0x64, 0x98, 0xe5, 0x43, 0xe3, 0xdc, 0x18,
	0x9b, 0xf4, 0x6c, 0x3f, 0xfb, 0xa5, 0x49, 0x3a, 0x23, 0xe8, 0x50, 0x9d, 0x40, 0xa1, 0x58, 0x06,
	0x60, 0xde, 0x60, 0xa9, 0x5b, 0xba, 0x54, 0x85, 0xce, 0x14, 0x7a, 0x1e, 0x5b, 0x89, 0x32, 0x57,
	0x7d, 0x07, 0x4b, 0xc8, 0x08, 0x3a, 0xb7, 0x61, 0x9a, 0x44, 0xc1, 0x9a, 0xc9, 0x24, 0x1d, 0x9e,
	0xe8, 0x79, 0xa0, 0xa1, 0x6b, 0x85, 0x38, 0x21, 0x58, 0x1e, 0x8b, 0x72, 0x9e, 0x30, 0x49, 0x2e,
	0xc0, 0xd2, 0xba, 0x57, 0x3c, 0xd5, 0x1c, 0xfd, 0x57, 0x23, 0xf7, 0xa0, 0x5d, 0x77, 0x51, 0x97,
	0xd1, 0x5d, 0x03, 0x19, 0xc2, 0xc3, 0x30, 0x8a, 0x04, 0x16, 0x85, 0x9e, 0xd2, 0xa6, 0xcd, 0xd5,
	0xf9, 0x6d, 0x00, 0xf8, 0x61, 0x86, 0x14, 0x57, 0x5c, 0x44, 0x84, 0xc0, 0x29, 0x0b, 0x33, 0xd4,
	0x13, 0xda, 0x54, 0xc7, 0xe4, 0x19, 0xb4, 0xf9, 0x0f, 0x86, 0x22, 0x50, 0xf2, 0x4f, 0xb4, 0x7c,
	0x4b, 0x03, 0xca, 0xd5, 0x05, 0x58, 0x58, 0x4b, 0x1c, 0x9a, 0xe7, 0xe6, 0xb8, 0x73, 0x54, 0x56,
	0xe3, 0x84, 0xee, 0x1a, 0xc8, 0x0c, 0xfa, 0xb8, 0xdb, 0x91, 0xa6, 0x3f, 0xd5, 0x14, 0x2f, 0x8e,
	0x52, 0xec, 0x2d, 0x94, 0xf6, 0x70, 0xff, 0xea, 0x7c, 0x50, 0x2f, 0x52, 0xa0, 0xb8, 0x45, 0xe5,
	0xe7, 0xa0, 0x93, 0xe7, 0xd0, 0xab, 0x9c, 0xdc, 0x5f, 0x46, 0x57, 0x83, 0x97, 0x15, 0xf6, 0xf2,
	0x2d, 0x58, 0xcd, 0x06, 0x49, 0x1f, 0x60, 0x31, 0x0f, 0xae, 0xfd, 0x99, 0xff, 0xf9, 0xab, 0x3f,
	0x68, 0x91, 0x2e, 0x58, 0x8b, 0x79, 0xe0, 0x7d, 0xba, 0xbc, 0x9a, 0x0f, 0x8c, 0x3a, 0xfb, 0xd1,
	0xf3, 0x3d, 0x7a, 0xf5, 0x7e, 0x10, 0x4d, 0xa7, 0x7f, 0x36, 0xb6, 0x71, 0xb7, 0xb1, 0x8d, 0x7f,
	0x1b, 0xdb, 0xf8, 0xb5, 0xb5, 0x5b, 0x77, 0x5b, 0xbb, 0xf5, 0x77, 0x6b, 0xb7, 0xbe, 0x8d, 0xe3,
	0x44, 0x7e, 0x5f, 0x2f, 0xdd, 0x15, 0xcf, 0x26, 0xb5, 0xb5, 0xfa, 0xe7, 0x4e, 0x7e, 0x36, 0x81,
	0x2c, 0x73, 0x2c, 0x96, 0x0f, 0xf4, 0x9b, 0xbd, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xf8,
	0x6b, 0xc3, 0xe0, 0x02, 0x00, 0x00,
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RegistrationTimestamp != 0 {
		i = encodeVarintUbkoin(dAtA, i, uint64(m.RegistrationTimestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RegisterKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EncryptionKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncryptionKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncryptionKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValidUntil != 0 {
		i = encodeVarintUbkoin(dAtA, i, uint64(m.ValidUntil))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Endpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Protocol != 0 {
		i = encodeVarintUbkoin(dAtA, i, uint64(m.Protocol))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NameRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptionKey) > 0 {
		for iNdEx := len(m.EncryptionKey) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EncryptionKey[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUbkoin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Endpoint) > 0 {
		for iNdEx := len(m.Endpoint) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoint[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUbkoin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OwnerKey) > 0 {
		i -= len(m.OwnerKey)
		copy(dAtA[i:], m.OwnerKey)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.OwnerKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReserveName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReserveName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReserveName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUbkoin(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUbkoin(dAtA []byte, offset int, v uint64) int {
	offset -= sovUbkoin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegistrationTimestamp != 0 {
		n += 1 + sovUbkoin(uint64(m.RegistrationTimestamp))
	}
	return n
}

func (m *RegisterKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	return n
}

func (m *EncryptionKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	if m.ValidUntil != 0 {
		n += 1 + sovUbkoin(uint64(m.ValidUntil))
	}
	return n
}

func (m *Endpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Protocol != 0 {
		n += 1 + sovUbkoin(uint64(m.Protocol))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	return n
}

func (m *NameRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	l = len(m.OwnerKey)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	if len(m.Endpoint) > 0 {
		for _, e := range m.Endpoint {
			l = e.Size()
			n += 1 + l + sovUbkoin(uint64(l))
		}
	}
	if len(m.EncryptionKey) > 0 {
		for _, e := range m.EncryptionKey {
			l = e.Size()
			n += 1 + l + sovUbkoin(uint64(l))
		}
	}
	return n
}

func (m *ReserveName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovUbkoin(uint64(l))
	}
	return n
}

func sovUbkoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUbkoin(x uint64) (n int) {
	return sovUbkoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistrationTimestamp", wireType)
			}
			m.RegistrationTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistrationTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUbkoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbkoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUbkoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbkoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EncryptionKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EncryptionKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncryptionKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			m.ValidUntil = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUntil |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUbkoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbkoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Protocol |= Protocol(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUbkoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbkoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NameRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NameRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerKey = append(m.OwnerKey[:0], dAtA[iNdEx:postIndex]...)
			if m.OwnerKey == nil {
				m.OwnerKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = append(m.Endpoint, &Endpoint{})
			if err := m.Endpoint[len(m.Endpoint)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionKey = append(m.EncryptionKey, &EncryptionKey{})
			if err := m.EncryptionKey[len(m.EncryptionKey)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUbkoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbkoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReserveName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReserveName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReserveName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUbkoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbkoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUbkoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbkoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUbkoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUbkoin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUbkoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUbkoin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUbkoin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUbkoin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUbkoin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUbkoin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUbkoin = fmt.Errorf("proto: unexpected end of group")
)

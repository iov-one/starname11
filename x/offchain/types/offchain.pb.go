// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/offchain/v1alpha1/offchain.proto

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

// MsgSignData defines an arbitrary, general-purpose, off-chain message
type MsgSignData struct {
	// signer is the bech32 representation of the signer's account address
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// data represents the raw bytes of the content that is signed (text, json,
	// etc)
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgSignData) Reset()         { *m = MsgSignData{} }
func (m *MsgSignData) String() string { return proto.CompactTextString(m) }
func (*MsgSignData) ProtoMessage()    {}
func (*MsgSignData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8921ced3d220e79e, []int{0}
}
func (m *MsgSignData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignData.Merge(m, src)
}
func (m *MsgSignData) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignData proto.InternalMessageInfo

func (m *MsgSignData) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgSignData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// ListOfMsgSignData defines a list of MsgSignData, used to marshal and
// unmarshal them in a clean way
type ListOfMsgSignData struct {
	// msgs is a list of messages
	Msgs []*MsgSignData `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *ListOfMsgSignData) Reset()         { *m = ListOfMsgSignData{} }
func (m *ListOfMsgSignData) String() string { return proto.CompactTextString(m) }
func (*ListOfMsgSignData) ProtoMessage()    {}
func (*ListOfMsgSignData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8921ced3d220e79e, []int{1}
}
func (m *ListOfMsgSignData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListOfMsgSignData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListOfMsgSignData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListOfMsgSignData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOfMsgSignData.Merge(m, src)
}
func (m *ListOfMsgSignData) XXX_Size() int {
	return m.Size()
}
func (m *ListOfMsgSignData) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOfMsgSignData.DiscardUnknown(m)
}

var xxx_messageInfo_ListOfMsgSignData proto.InternalMessageInfo

func (m *ListOfMsgSignData) GetMsgs() []*MsgSignData {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSignData)(nil), "cosmos.offchain.v1alpha1.MsgSignData")
	proto.RegisterType((*ListOfMsgSignData)(nil), "cosmos.offchain.v1alpha1.ListOfMsgSignData")
}

func init() {
	proto.RegisterFile("iov/offchain/v1alpha1/offchain.proto", fileDescriptor_8921ced3d220e79e)
}

var fileDescriptor_8921ced3d220e79e = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0xcc, 0x2f, 0xd3,
	0xcf, 0x4f, 0x4b, 0x4b, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48,
	0x34, 0x84, 0x8b, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x24, 0xe7, 0x17, 0xe7, 0xe6,
	0x17, 0xeb, 0xc1, 0x85, 0x61, 0x0a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x8a, 0xf4, 0x41,
	0x2c, 0x88, 0x7a, 0x25, 0x4b, 0x2e, 0x6e, 0xdf, 0xe2, 0xf4, 0xe0, 0xcc, 0xf4, 0x3c, 0x97, 0xc4,
	0x92, 0x44, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x49, 0x81,
	0x51, 0x83, 0x27, 0x08, 0xcc, 0x56, 0xf2, 0xe3, 0x12, 0xf4, 0xc9, 0x2c, 0x2e, 0xf1, 0x4f, 0x43,
	0x36, 0xc0, 0x92, 0x8b, 0x25, 0xb7, 0x38, 0xbd, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48,
	0x55, 0x0f, 0x97, 0x73, 0xf4, 0x90, 0x34, 0x05, 0x81, 0xb5, 0x38, 0xb9, 0x9f, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x66, 0x7e, 0x99, 0x6e, 0x7e, 0x5e, 0xaa, 0x7e, 0x71, 0x49, 0x62, 0x51, 0x5e,
	0x62, 0x6e, 0x6a, 0x8a, 0x7e, 0x05, 0x22, 0x54, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0,
	0x5e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xb6, 0xb7, 0xbd, 0x32, 0x01, 0x00, 0x00,
}

func (m *MsgSignData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintOffchain(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintOffchain(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListOfMsgSignData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListOfMsgSignData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListOfMsgSignData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOffchain(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintOffchain(dAtA []byte, offset int, v uint64) int {
	offset -= sovOffchain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSignData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovOffchain(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovOffchain(uint64(l))
	}
	return n
}

func (m *ListOfMsgSignData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovOffchain(uint64(l))
		}
	}
	return n
}

func sovOffchain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOffchain(x uint64) (n int) {
	return sovOffchain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSignData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffchain
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
			return fmt.Errorf("proto: MsgSignData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchain
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
				return ErrInvalidLengthOffchain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchain
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
				return ErrInvalidLengthOffchain
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOffchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffchain
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
func (m *ListOfMsgSignData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffchain
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
			return fmt.Errorf("proto: ListOfMsgSignData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListOfMsgSignData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffchain
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
				return ErrInvalidLengthOffchain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &MsgSignData{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffchain
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
func skipOffchain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOffchain
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
					return 0, ErrIntOverflowOffchain
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
					return 0, ErrIntOverflowOffchain
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
				return 0, ErrInvalidLengthOffchain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOffchain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOffchain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOffchain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOffchain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOffchain = fmt.Errorf("proto: unexpected end of group")
)

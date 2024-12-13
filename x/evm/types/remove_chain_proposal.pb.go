// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/remove_chain_proposal.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RemoveChainProposal struct {
	Title            string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description      string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ChainReferenceID string `protobuf:"bytes,3,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
}

func (m *RemoveChainProposal) Reset()         { *m = RemoveChainProposal{} }
func (m *RemoveChainProposal) String() string { return proto.CompactTextString(m) }
func (*RemoveChainProposal) ProtoMessage()    {}
func (*RemoveChainProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b2845f61092eb8, []int{0}
}

func (m *RemoveChainProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RemoveChainProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveChainProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RemoveChainProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveChainProposal.Merge(m, src)
}

func (m *RemoveChainProposal) XXX_Size() int {
	return m.Size()
}

func (m *RemoveChainProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveChainProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveChainProposal proto.InternalMessageInfo

func (m *RemoveChainProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RemoveChainProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoveChainProposal) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func init() {
	proto.RegisterType((*RemoveChainProposal)(nil), "palomachain.paloma.evm.RemoveChainProposal")
}

func init() {
	proto.RegisterFile("palomachain/paloma/evm/remove_chain_proposal.proto", fileDescriptor_a0b2845f61092eb8)
}

var fileDescriptor_a0b2845f61092eb8 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2a, 0x48, 0xcc, 0xc9,
	0xcf, 0x4d, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0xb0, 0xf5, 0x53, 0xcb, 0x72, 0xf5, 0x8b,
	0x52, 0x73, 0xf3, 0xcb, 0x52, 0xe3, 0xc1, 0xe2, 0xf1, 0x05, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0x89,
	0x39, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x62, 0x48, 0x7a, 0xf4, 0x20, 0x6c, 0xbd, 0xd4,
	0xb2, 0x5c, 0xa5, 0x4a, 0x2e, 0xe1, 0x20, 0xb0, 0x36, 0x67, 0x90, 0x4c, 0x00, 0x54, 0x93, 0x90,
	0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84,
	0x23, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99, 0x9f, 0x27,
	0xc1, 0x04, 0x96, 0x43, 0x16, 0x12, 0xd2, 0xe2, 0x12, 0x00, 0x5b, 0x11, 0x94, 0x9a, 0x96, 0x5a,
	0x94, 0x9a, 0x97, 0x9c, 0xea, 0xe9, 0x22, 0xc1, 0x0c, 0x56, 0x86, 0x21, 0xee, 0xe4, 0x76, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70,
	0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x3a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x58, 0xfc, 0x5a, 0x66, 0xa4, 0x5f, 0x01, 0xf6, 0x70, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x87, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xb9,
	0x70, 0x6c, 0x17, 0x01, 0x00, 0x00,
}

func (m *RemoveChainProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveChainProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveChainProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintRemoveChainProposal(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintRemoveChainProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintRemoveChainProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRemoveChainProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovRemoveChainProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *RemoveChainProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovRemoveChainProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovRemoveChainProposal(uint64(l))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovRemoveChainProposal(uint64(l))
	}
	return n
}

func sovRemoveChainProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozRemoveChainProposal(x uint64) (n int) {
	return sovRemoveChainProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *RemoveChainProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoveChainProposal
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
			return fmt.Errorf("proto: RemoveChainProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveChainProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoveChainProposal
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
				return ErrInvalidLengthRemoveChainProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRemoveChainProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoveChainProposal
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
				return ErrInvalidLengthRemoveChainProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRemoveChainProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoveChainProposal
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
				return ErrInvalidLengthRemoveChainProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRemoveChainProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemoveChainProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRemoveChainProposal
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

func skipRemoveChainProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRemoveChainProposal
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
					return 0, ErrIntOverflowRemoveChainProposal
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
					return 0, ErrIntOverflowRemoveChainProposal
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
				return 0, ErrInvalidLengthRemoveChainProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRemoveChainProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRemoveChainProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRemoveChainProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRemoveChainProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRemoveChainProposal = fmt.Errorf("proto: unexpected end of group")
)

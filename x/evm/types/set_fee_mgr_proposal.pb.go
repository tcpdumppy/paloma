// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/set_fee_mgr_proposal.proto

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

type SetFeeManagerAddressProposal struct {
	Title             string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Summary           string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	ChainReferenceID  string `protobuf:"bytes,3,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	FeeManagerAddress string `protobuf:"bytes,4,opt,name=feeManagerAddress,proto3" json:"feeManagerAddress,omitempty"`
}

func (m *SetFeeManagerAddressProposal) Reset()         { *m = SetFeeManagerAddressProposal{} }
func (m *SetFeeManagerAddressProposal) String() string { return proto.CompactTextString(m) }
func (*SetFeeManagerAddressProposal) ProtoMessage()    {}
func (*SetFeeManagerAddressProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b189b7cf6349646e, []int{0}
}

func (m *SetFeeManagerAddressProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SetFeeManagerAddressProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetFeeManagerAddressProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SetFeeManagerAddressProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFeeManagerAddressProposal.Merge(m, src)
}

func (m *SetFeeManagerAddressProposal) XXX_Size() int {
	return m.Size()
}

func (m *SetFeeManagerAddressProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFeeManagerAddressProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetFeeManagerAddressProposal proto.InternalMessageInfo

func (m *SetFeeManagerAddressProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetFeeManagerAddressProposal) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *SetFeeManagerAddressProposal) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *SetFeeManagerAddressProposal) GetFeeManagerAddress() string {
	if m != nil {
		return m.FeeManagerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*SetFeeManagerAddressProposal)(nil), "palomachain.paloma.evm.SetFeeManagerAddressProposal")
}

func init() {
	proto.RegisterFile("palomachain/paloma/evm/set_fee_mgr_proposal.proto", fileDescriptor_b189b7cf6349646e)
}

var fileDescriptor_b189b7cf6349646e = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x48, 0xcc, 0xc9,
	0xcf, 0x4d, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0xb0, 0xf5, 0x53, 0xcb, 0x72, 0xf5, 0x8b,
	0x53, 0x4b, 0xe2, 0xd3, 0x52, 0x53, 0xe3, 0x73, 0xd3, 0x8b, 0xe2, 0x0b, 0x8a, 0xf2, 0x0b, 0xf2,
	0x8b, 0x13, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x90, 0xb4, 0xe8, 0x41, 0xd8,
	0x7a, 0xa9, 0x65, 0xb9, 0x4a, 0x2b, 0x18, 0xb9, 0x64, 0x82, 0x53, 0x4b, 0xdc, 0x52, 0x53, 0x7d,
	0x13, 0xf3, 0x12, 0xd3, 0x53, 0x8b, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x03, 0xa0, 0xda,
	0x85, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x20, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0xe2, 0xd2, 0xdc, 0xdc, 0xc4, 0xa2, 0x4a, 0x09, 0x26, 0xb0,
	0x38, 0x8c, 0x2b, 0xa4, 0xc5, 0x25, 0x00, 0xb6, 0x24, 0x28, 0x35, 0x2d, 0xb5, 0x28, 0x35, 0x2f,
	0x39, 0xd5, 0xd3, 0x45, 0x82, 0x19, 0xac, 0x04, 0x43, 0x5c, 0x48, 0x87, 0x4b, 0x30, 0x0d, 0xdd,
	0x62, 0x09, 0x16, 0xb0, 0x62, 0x4c, 0x09, 0x27, 0xb7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xc7,
	0x12, 0x34, 0x65, 0x46, 0xfa, 0x15, 0xe0, 0xf0, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03,
	0x87, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x59, 0x57, 0x5e, 0x64, 0x46, 0x01, 0x00, 0x00,
}

func (m *SetFeeManagerAddressProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetFeeManagerAddressProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetFeeManagerAddressProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeManagerAddress) > 0 {
		i -= len(m.FeeManagerAddress)
		copy(dAtA[i:], m.FeeManagerAddress)
		i = encodeVarintSetFeeMgrProposal(dAtA, i, uint64(len(m.FeeManagerAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintSetFeeMgrProposal(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Summary) > 0 {
		i -= len(m.Summary)
		copy(dAtA[i:], m.Summary)
		i = encodeVarintSetFeeMgrProposal(dAtA, i, uint64(len(m.Summary)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSetFeeMgrProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSetFeeMgrProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovSetFeeMgrProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *SetFeeManagerAddressProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSetFeeMgrProposal(uint64(l))
	}
	l = len(m.Summary)
	if l > 0 {
		n += 1 + l + sovSetFeeMgrProposal(uint64(l))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovSetFeeMgrProposal(uint64(l))
	}
	l = len(m.FeeManagerAddress)
	if l > 0 {
		n += 1 + l + sovSetFeeMgrProposal(uint64(l))
	}
	return n
}

func sovSetFeeMgrProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozSetFeeMgrProposal(x uint64) (n int) {
	return sovSetFeeMgrProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *SetFeeManagerAddressProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSetFeeMgrProposal
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
			return fmt.Errorf("proto: SetFeeManagerAddressProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetFeeManagerAddressProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetFeeMgrProposal
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
				return ErrInvalidLengthSetFeeMgrProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSetFeeMgrProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Summary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetFeeMgrProposal
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
				return ErrInvalidLengthSetFeeMgrProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSetFeeMgrProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Summary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetFeeMgrProposal
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
				return ErrInvalidLengthSetFeeMgrProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSetFeeMgrProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeManagerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetFeeMgrProposal
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
				return ErrInvalidLengthSetFeeMgrProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSetFeeMgrProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeManagerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSetFeeMgrProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSetFeeMgrProposal
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

func skipSetFeeMgrProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSetFeeMgrProposal
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
					return 0, ErrIntOverflowSetFeeMgrProposal
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
					return 0, ErrIntOverflowSetFeeMgrProposal
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
				return 0, ErrInvalidLengthSetFeeMgrProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSetFeeMgrProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSetFeeMgrProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSetFeeMgrProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSetFeeMgrProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSetFeeMgrProposal = fmt.Errorf("proto: unexpected end of group")
)

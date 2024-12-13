// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/paloma/light_node_client_funders_proposal.proto

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

type SetLightNodeClientFundersProposal struct {
	Title          string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	FunderAccounts []string `protobuf:"bytes,3,rep,name=funder_accounts,json=funderAccounts,proto3" json:"funder_accounts,omitempty"`
}

func (m *SetLightNodeClientFundersProposal) Reset()         { *m = SetLightNodeClientFundersProposal{} }
func (m *SetLightNodeClientFundersProposal) String() string { return proto.CompactTextString(m) }
func (*SetLightNodeClientFundersProposal) ProtoMessage()    {}
func (*SetLightNodeClientFundersProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ae998f0acf2d7e9, []int{0}
}

func (m *SetLightNodeClientFundersProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SetLightNodeClientFundersProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetLightNodeClientFundersProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SetLightNodeClientFundersProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetLightNodeClientFundersProposal.Merge(m, src)
}

func (m *SetLightNodeClientFundersProposal) XXX_Size() int {
	return m.Size()
}

func (m *SetLightNodeClientFundersProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetLightNodeClientFundersProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetLightNodeClientFundersProposal proto.InternalMessageInfo

func (m *SetLightNodeClientFundersProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetLightNodeClientFundersProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SetLightNodeClientFundersProposal) GetFunderAccounts() []string {
	if m != nil {
		return m.FunderAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*SetLightNodeClientFundersProposal)(nil), "palomachain.paloma.paloma.SetLightNodeClientFundersProposal")
}

func init() {
	proto.RegisterFile("palomachain/paloma/paloma/light_node_client_funders_proposal.proto", fileDescriptor_3ae998f0acf2d7e9)
}

var fileDescriptor_3ae998f0acf2d7e9 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2a, 0x48, 0xcc, 0xc9,
	0xcf, 0x4d, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0xb0, 0x61, 0x54, 0x4e, 0x66, 0x7a, 0x46,
	0x49, 0x7c, 0x5e, 0x7e, 0x4a, 0x6a, 0x7c, 0x72, 0x4e, 0x66, 0x6a, 0x5e, 0x49, 0x7c, 0x5a, 0x69,
	0x5e, 0x4a, 0x6a, 0x51, 0x71, 0x7c, 0x41, 0x51, 0x7e, 0x41, 0x7e, 0x71, 0x62, 0x8e, 0x5e, 0x41,
	0x51, 0x7e, 0x49, 0xbe, 0x90, 0x24, 0x92, 0x19, 0x7a, 0x10, 0x36, 0x94, 0x52, 0x6a, 0x61, 0xe4,
	0x52, 0x0c, 0x4e, 0x2d, 0xf1, 0x01, 0x19, 0xe5, 0x97, 0x9f, 0x92, 0xea, 0x0c, 0x36, 0xc8, 0x0d,
	0x62, 0x4e, 0x00, 0xd4, 0x18, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92, 0x9c, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38, 0xb9, 0x28,
	0xb3, 0xa0, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x09, 0x2c, 0x87, 0x2c, 0x24, 0xa4, 0xce, 0xc5, 0x0f,
	0x71, 0x52, 0x7c, 0x62, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x49, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06,
	0x67, 0x10, 0x1f, 0x44, 0xd8, 0x11, 0x2a, 0xea, 0xe4, 0x79, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0xfa, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x58, 0x82, 0xa2, 0xcc, 0x48, 0xbf, 0x02, 0xc6, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03,
	0xfb, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x16, 0x64, 0x6c, 0x76, 0x39, 0x01, 0x00, 0x00,
}

func (m *SetLightNodeClientFundersProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetLightNodeClientFundersProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetLightNodeClientFundersProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FunderAccounts) > 0 {
		for iNdEx := len(m.FunderAccounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FunderAccounts[iNdEx])
			copy(dAtA[i:], m.FunderAccounts[iNdEx])
			i = encodeVarintLightNodeClientFundersProposal(dAtA, i, uint64(len(m.FunderAccounts[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintLightNodeClientFundersProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintLightNodeClientFundersProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLightNodeClientFundersProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovLightNodeClientFundersProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *SetLightNodeClientFundersProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovLightNodeClientFundersProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovLightNodeClientFundersProposal(uint64(l))
	}
	if len(m.FunderAccounts) > 0 {
		for _, s := range m.FunderAccounts {
			l = len(s)
			n += 1 + l + sovLightNodeClientFundersProposal(uint64(l))
		}
	}
	return n
}

func sovLightNodeClientFundersProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozLightNodeClientFundersProposal(x uint64) (n int) {
	return sovLightNodeClientFundersProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *SetLightNodeClientFundersProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLightNodeClientFundersProposal
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
			return fmt.Errorf("proto: SetLightNodeClientFundersProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetLightNodeClientFundersProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightNodeClientFundersProposal
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
				return ErrInvalidLengthLightNodeClientFundersProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLightNodeClientFundersProposal
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
					return ErrIntOverflowLightNodeClientFundersProposal
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
				return ErrInvalidLengthLightNodeClientFundersProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLightNodeClientFundersProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunderAccounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightNodeClientFundersProposal
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
				return ErrInvalidLengthLightNodeClientFundersProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLightNodeClientFundersProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunderAccounts = append(m.FunderAccounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLightNodeClientFundersProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLightNodeClientFundersProposal
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

func skipLightNodeClientFundersProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLightNodeClientFundersProposal
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
					return 0, ErrIntOverflowLightNodeClientFundersProposal
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
					return 0, ErrIntOverflowLightNodeClientFundersProposal
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
				return 0, ErrInvalidLengthLightNodeClientFundersProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLightNodeClientFundersProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLightNodeClientFundersProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLightNodeClientFundersProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLightNodeClientFundersProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLightNodeClientFundersProposal = fmt.Errorf("proto: unexpected end of group")
)

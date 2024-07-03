// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/skyway/ethereum_signer.proto

package types

import (
	fmt "fmt"
	math "math"

	_ "github.com/cosmos/gogoproto/gogoproto"
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

// SignType defines messages that have been signed by an orchestrator
type SignType int32

const (
	// An unspecified type
	SIGN_TYPE_UNSPECIFIED SignType = 0
	// A type for multi-sig updates
	SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE SignType = 1
	// A type for batches
	SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH SignType = 2
)

var SignType_name = map[int32]string{
	0: "SIGN_TYPE_UNSPECIFIED",
	1: "SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE",
	2: "SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH",
}

var SignType_value = map[string]int32{
	"SIGN_TYPE_UNSPECIFIED":                          0,
	"SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE": 1,
	"SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH":   2,
}

func (SignType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7a59dec35d73c73, []int{0}
}

func init() {
	proto.RegisterEnum("palomachain.paloma.skyway.SignType", SignType_name, SignType_value)
}

func init() {
	proto.RegisterFile("palomachain/paloma/skyway/ethereum_signer.proto", fileDescriptor_b7a59dec35d73c73)
}

var fileDescriptor_b7a59dec35d73c73 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x48, 0xcc, 0xc9,
	0xcf, 0x4d, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0x83, 0xb2, 0xf5, 0x8b, 0xb3, 0x2b, 0xcb, 0x13, 0x2b,
	0xf5, 0x53, 0x4b, 0x32, 0x52, 0x8b, 0x52, 0x4b, 0x73, 0xe3, 0x8b, 0x33, 0xd3, 0xf3, 0x52, 0x8b,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0x91, 0x34, 0xe8, 0x41, 0xd8, 0x7a, 0x10, 0x0d,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x55, 0xfa, 0x20, 0x16, 0x44, 0x83, 0xd6, 0x54, 0x46,
	0x2e, 0x8e, 0xe0, 0xcc, 0xf4, 0xbc, 0x90, 0xca, 0x82, 0x54, 0x21, 0x49, 0x2e, 0xd1, 0x60, 0x4f,
	0x77, 0xbf, 0xf8, 0x90, 0xc8, 0x00, 0xd7, 0xf8, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37,
	0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x23, 0x2e, 0x3d, 0x84, 0x94, 0x7f, 0x90, 0xb3, 0x87, 0x6b,
	0x70, 0x48, 0x90, 0x63, 0x88, 0x7f, 0x50, 0x3c, 0x48, 0xd8, 0xd5, 0x25, 0xde, 0x37, 0xd4, 0x27,
	0xc4, 0x13, 0xc4, 0x89, 0x0f, 0x0d, 0x70, 0x71, 0x0c, 0x71, 0x15, 0x60, 0x14, 0x32, 0xe0, 0xd2,
	0xc1, 0xaf, 0x27, 0xdc, 0x33, 0xc4, 0xc3, 0x25, 0xc8, 0x31, 0x3c, 0xde, 0xc9, 0x31, 0xc4, 0xd9,
	0x43, 0x80, 0x49, 0x8a, 0xa3, 0x63, 0xb1, 0x1c, 0xc3, 0x8a, 0x25, 0x72, 0x0c, 0x4e, 0x6e, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x93, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0x8b, 0x2d, 0x78, 0x2a, 0x60, 0x01, 0x54, 0x52, 0x59, 0x90, 0x5a, 0x9c,
	0xc4, 0x06, 0xf6, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x31, 0x0a, 0xa4, 0x02, 0x4a, 0x01,
	0x00, 0x00,
}

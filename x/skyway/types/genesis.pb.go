// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/skyway/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState struct, containing all persistant data required by the Skyway
// module
type GenesisState struct {
	Params                 *Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	SkywayNonces           []SkywayNonces           `protobuf:"bytes,2,rep,name=skyway_nonces,json=skywayNonces,proto3" json:"skyway_nonces"`
	Batches                []OutgoingTxBatch        `protobuf:"bytes,3,rep,name=batches,proto3" json:"batches"`
	BatchConfirms          []MsgConfirmBatch        `protobuf:"bytes,4,rep,name=batch_confirms,json=batchConfirms,proto3" json:"batch_confirms"`
	BatchGasEstimates      []MsgEstimateBatchGas    `protobuf:"bytes,5,rep,name=batch_gas_estimates,json=batchGasEstimates,proto3" json:"batch_gas_estimates"`
	Attestations           []Attestation            `protobuf:"bytes,7,rep,name=attestations,proto3" json:"attestations"`
	Erc20ToDenoms          []ERC20ToDenom           `protobuf:"bytes,9,rep,name=erc20_to_denoms,json=erc20ToDenoms,proto3" json:"erc20_to_denoms"`
	UnbatchedTransfers     []OutgoingTransferTx     `protobuf:"bytes,10,rep,name=unbatched_transfers,json=unbatchedTransfers,proto3" json:"unbatched_transfers"`
	BridgeTransferLimits   []*BridgeTransferLimit   `protobuf:"bytes,12,rep,name=bridge_transfer_limits,json=bridgeTransferLimits,proto3" json:"bridge_transfer_limits,omitempty"`
	BridgeTaxes            []*BridgeTax             `protobuf:"bytes,13,rep,name=bridge_taxes,json=bridgeTaxes,proto3" json:"bridge_taxes,omitempty"`
	LightNodeSaleContracts []*LightNodeSaleContract `protobuf:"bytes,14,rep,name=light_node_sale_contracts,json=lightNodeSaleContracts,proto3" json:"light_node_sale_contracts,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20fa6374f01ce45, []int{0}
}

func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}

func (m *GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetSkywayNonces() []SkywayNonces {
	if m != nil {
		return m.SkywayNonces
	}
	return nil
}

func (m *GenesisState) GetBatches() []OutgoingTxBatch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *GenesisState) GetBatchConfirms() []MsgConfirmBatch {
	if m != nil {
		return m.BatchConfirms
	}
	return nil
}

func (m *GenesisState) GetBatchGasEstimates() []MsgEstimateBatchGas {
	if m != nil {
		return m.BatchGasEstimates
	}
	return nil
}

func (m *GenesisState) GetAttestations() []Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

func (m *GenesisState) GetErc20ToDenoms() []ERC20ToDenom {
	if m != nil {
		return m.Erc20ToDenoms
	}
	return nil
}

func (m *GenesisState) GetUnbatchedTransfers() []OutgoingTransferTx {
	if m != nil {
		return m.UnbatchedTransfers
	}
	return nil
}

func (m *GenesisState) GetBridgeTransferLimits() []*BridgeTransferLimit {
	if m != nil {
		return m.BridgeTransferLimits
	}
	return nil
}

func (m *GenesisState) GetBridgeTaxes() []*BridgeTax {
	if m != nil {
		return m.BridgeTaxes
	}
	return nil
}

func (m *GenesisState) GetLightNodeSaleContracts() []*LightNodeSaleContract {
	if m != nil {
		return m.LightNodeSaleContracts
	}
	return nil
}

// SkywayCounters contains the many noces and counters required to maintain the
// bridge state in the genesis
type SkywayNonces struct {
	// the last observed Skyway.sol contract event nonce
	LastObservedNonce uint64 `protobuf:"varint,1,opt,name=last_observed_nonce,json=lastObservedNonce,proto3" json:"last_observed_nonce,omitempty"`
	// the last batch Cosmos chain block that batch slashing has completed for
	// there is an individual batch nonce for each token type so this removes
	// the need to store them all
	LastSlashedBatchBlock uint64 `protobuf:"varint,2,opt,name=last_slashed_batch_block,json=lastSlashedBatchBlock,proto3" json:"last_slashed_batch_block,omitempty"`
	// the last transaction id from the Skyway TX pool, this prevents ID
	// duplication during chain upgrades
	LastTxPoolId uint64 `protobuf:"varint,3,opt,name=last_tx_pool_id,json=lastTxPoolId,proto3" json:"last_tx_pool_id,omitempty"`
	// the last batch id from the Skyway batch pool, this prevents ID duplication
	// during chain upgrades
	LastBatchId uint64 `protobuf:"varint,4,opt,name=last_batch_id,json=lastBatchId,proto3" json:"last_batch_id,omitempty"`
	// the reference id of the remote chain this data applies to.
	ChainReferenceId string `protobuf:"bytes,5,opt,name=chain_reference_id,json=chainReferenceId,proto3" json:"chain_reference_id,omitempty"`
}

func (m *SkywayNonces) Reset()         { *m = SkywayNonces{} }
func (m *SkywayNonces) String() string { return proto.CompactTextString(m) }
func (*SkywayNonces) ProtoMessage()    {}
func (*SkywayNonces) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20fa6374f01ce45, []int{1}
}

func (m *SkywayNonces) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SkywayNonces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SkywayNonces.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SkywayNonces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkywayNonces.Merge(m, src)
}

func (m *SkywayNonces) XXX_Size() int {
	return m.Size()
}

func (m *SkywayNonces) XXX_DiscardUnknown() {
	xxx_messageInfo_SkywayNonces.DiscardUnknown(m)
}

var xxx_messageInfo_SkywayNonces proto.InternalMessageInfo

func (m *SkywayNonces) GetLastObservedNonce() uint64 {
	if m != nil {
		return m.LastObservedNonce
	}
	return 0
}

func (m *SkywayNonces) GetLastSlashedBatchBlock() uint64 {
	if m != nil {
		return m.LastSlashedBatchBlock
	}
	return 0
}

func (m *SkywayNonces) GetLastTxPoolId() uint64 {
	if m != nil {
		return m.LastTxPoolId
	}
	return 0
}

func (m *SkywayNonces) GetLastBatchId() uint64 {
	if m != nil {
		return m.LastBatchId
	}
	return 0
}

func (m *SkywayNonces) GetChainReferenceId() string {
	if m != nil {
		return m.ChainReferenceId
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "palomachain.paloma.skyway.GenesisState")
	proto.RegisterType((*SkywayNonces)(nil), "palomachain.paloma.skyway.SkywayNonces")
}

func init() {
	proto.RegisterFile("palomachain/paloma/skyway/genesis.proto", fileDescriptor_e20fa6374f01ce45)
}

var fileDescriptor_e20fa6374f01ce45 = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x13, 0x08, 0x70, 0x99, 0x38, 0xfc, 0x19, 0xb8, 0x68, 0x60, 0x91, 0xcb, 0x45, 0xfc,
	0x13, 0xf7, 0xd6, 0x86, 0x54, 0x55, 0xdb, 0x65, 0x43, 0x11, 0x0a, 0xa2, 0x80, 0x9c, 0x54, 0x95,
	0xba, 0xb1, 0xc6, 0xf6, 0xe0, 0x58, 0xd8, 0x9e, 0xc8, 0x67, 0xa0, 0xe1, 0x2d, 0xfa, 0x58, 0x2c,
	0x59, 0x76, 0x55, 0x55, 0xa0, 0xee, 0xfb, 0x08, 0xd5, 0x8c, 0xc7, 0x90, 0xaa, 0x38, 0xe9, 0x8a,
	0xe1, 0x7c, 0xdf, 0xf7, 0x9b, 0xe8, 0x1c, 0x9f, 0x41, 0x5b, 0x3d, 0x1a, 0xf1, 0x98, 0x7a, 0x5d,
	0x1a, 0x26, 0x56, 0x76, 0xb6, 0xe0, 0xe2, 0xfa, 0x13, 0xbd, 0xb6, 0x02, 0x96, 0x30, 0x08, 0xc1,
	0xec, 0xa5, 0x5c, 0x70, 0xbc, 0x3c, 0x60, 0x34, 0xb3, 0xb3, 0x99, 0x19, 0x57, 0xea, 0x1e, 0x87,
	0x98, 0x83, 0xe5, 0x52, 0x60, 0xd6, 0xd5, 0x9e, 0xcb, 0x04, 0xdd, 0xb3, 0x3c, 0x2e, 0x7d, 0x32,
	0xba, 0xb2, 0x18, 0xf0, 0x80, 0xab, 0xa3, 0x25, 0x4f, 0xba, 0xfa, 0x5f, 0xf1, 0xcd, 0x54, 0x08,
	0x06, 0x82, 0x8a, 0x90, 0xe7, 0x88, 0x8d, 0x62, 0xb3, 0x4b, 0x85, 0xd7, 0xd5, 0xb6, 0x9d, 0x21,
	0xb6, 0x34, 0xf4, 0x03, 0xe6, 0x08, 0xda, 0xd7, 0xde, 0x17, 0xa3, 0xbd, 0x29, 0x4d, 0xe0, 0x9c,
	0xa5, 0x4e, 0x14, 0xc6, 0xa1, 0xd0, 0xb1, 0x57, 0xc5, 0xb1, 0x28, 0x0c, 0xba, 0xc2, 0x49, 0xb8,
	0xcf, 0x1c, 0xa0, 0x11, 0x73, 0x3c, 0x9e, 0x88, 0x94, 0x7a, 0x79, 0x72, 0xbd, 0x38, 0x19, 0x43,
	0xa0, 0xfb, 0xbc, 0xb2, 0x59, 0xec, 0xea, 0xd1, 0x94, 0xc6, 0x30, 0xba, 0x23, 0xe2, 0xba, 0xc7,
	0xb4, 0x6d, 0xed, 0xfb, 0x14, 0x32, 0x0e, 0xb3, 0x41, 0xb6, 0x05, 0x15, 0x0c, 0xbf, 0x46, 0x93,
	0x19, 0x87, 0x94, 0x57, 0xcb, 0xdb, 0xd5, 0xc6, 0xbf, 0x66, 0xe1, 0x60, 0xcd, 0x33, 0x65, 0xb4,
	0x75, 0x00, 0xdb, 0xa8, 0x96, 0x09, 0x4e, 0xc2, 0x13, 0x8f, 0x01, 0x19, 0x5b, 0x1d, 0xdf, 0xae,
	0x36, 0xb6, 0x86, 0x10, 0xda, 0xea, 0xcf, 0x89, 0xb2, 0x37, 0x2b, 0x37, 0x5f, 0xff, 0x29, 0xd9,
	0x06, 0x0c, 0xd4, 0xf0, 0x11, 0x9a, 0x52, 0x03, 0x64, 0x40, 0xc6, 0x15, 0x6d, 0x67, 0x08, 0xed,
	0xf4, 0x52, 0x04, 0x3c, 0x4c, 0x82, 0x4e, 0xbf, 0x29, 0x33, 0x1a, 0x98, 0x03, 0xf0, 0x07, 0x34,
	0xa3, 0x8e, 0xb2, 0xf1, 0xe7, 0x61, 0x1a, 0x03, 0xa9, 0x8c, 0x44, 0xbe, 0x83, 0x60, 0x3f, 0x73,
	0x0f, 0x22, 0x6b, 0x8a, 0xa3, 0x05, 0xc0, 0x3e, 0x5a, 0xc8, 0xc0, 0x01, 0x05, 0x87, 0x81, 0x08,
	0x63, 0x2a, 0x18, 0x90, 0x09, 0x45, 0x37, 0x87, 0xd3, 0x0f, 0xb4, 0x5d, 0xe1, 0x0f, 0x69, 0xde,
	0x85, 0x79, 0x57, 0xff, 0x9f, 0xeb, 0x80, 0xcf, 0x90, 0x31, 0xf0, 0xe1, 0x03, 0x99, 0x52, 0xf8,
	0xcd, 0x21, 0xf8, 0x37, 0x8f, 0xf6, 0xbc, 0xb9, 0x83, 0x04, 0xfc, 0x1e, 0xcd, 0xb2, 0xd4, 0x6b,
	0xec, 0x3a, 0x82, 0x3b, 0x3e, 0x4b, 0x78, 0x0c, 0x64, 0x7a, 0xe4, 0xc8, 0x0e, 0xec, 0xfd, 0xc6,
	0x6e, 0x87, 0xbf, 0x95, 0xfe, 0xbc, 0x1d, 0x8a, 0xa2, 0x6b, 0xaa, 0x1d, 0x97, 0x49, 0xd6, 0x74,
	0xff, 0x61, 0x49, 0x80, 0x20, 0x85, 0x7e, 0xf6, 0x27, 0xf3, 0xd3, 0x99, 0x4e, 0x5f, 0x5f, 0x80,
	0x1f, 0x78, 0xb9, 0x24, 0x6f, 0x59, 0x7a, 0x72, 0x0f, 0x81, 0x18, 0x23, 0xfb, 0xde, 0x54, 0xc1,
	0x9c, 0x75, 0x2c, 0x63, 0xf6, 0xa2, 0xfb, 0x7b, 0x11, 0xf0, 0x21, 0x32, 0x1e, 0x5f, 0x06, 0x06,
	0xa4, 0xa6, 0xd8, 0xeb, 0xa3, 0xd9, 0xb4, 0x6f, 0x57, 0xdd, 0xfc, 0xc8, 0x00, 0x5f, 0xa0, 0xe5,
	0xa2, 0xfd, 0x07, 0x32, 0xa3, 0xa8, 0xbb, 0x43, 0xa8, 0xc7, 0x32, 0x7b, 0xc2, 0x7d, 0xd6, 0xa6,
	0x11, 0xdb, 0xd7, 0x41, 0x7b, 0x29, 0x7a, 0xaa, 0x0c, 0x47, 0x95, 0xbf, 0xaa, 0x73, 0x86, 0x8d,
	0x1e, 0x7f, 0xf9, 0xda, 0x8f, 0x32, 0x32, 0x06, 0x97, 0x0d, 0x9b, 0x68, 0x21, 0xa2, 0x20, 0x1c,
	0xee, 0x02, 0x4b, 0xaf, 0x98, 0x9f, 0xed, 0xac, 0x5a, 0xfa, 0x8a, 0x3d, 0x2f, 0xa5, 0x53, 0xad,
	0xa8, 0x00, 0x7e, 0x89, 0x88, 0xf2, 0x43, 0x44, 0x41, 0xce, 0x35, 0xfb, 0xe0, 0xdd, 0x88, 0x7b,
	0x17, 0x64, 0x4c, 0x85, 0xfe, 0x96, 0x7a, 0x3b, 0x93, 0xb3, 0x65, 0x91, 0x22, 0xde, 0x40, 0xb3,
	0x2a, 0x28, 0xfa, 0x4e, 0x8f, 0xf3, 0xc8, 0x09, 0x7d, 0x32, 0xae, 0xfc, 0x86, 0x2c, 0x77, 0xfa,
	0x67, 0x9c, 0x47, 0x2d, 0x1f, 0xaf, 0xa1, 0x9a, 0xb2, 0x65, 0xdc, 0xd0, 0x27, 0x15, 0x65, 0xaa,
	0xca, 0xa2, 0xa2, 0xb5, 0x7c, 0xfc, 0x3f, 0xc2, 0xaa, 0x37, 0x4e, 0xca, 0xce, 0x59, 0xca, 0x12,
	0x8f, 0x49, 0xe3, 0xc4, 0x6a, 0x79, 0x7b, 0xda, 0x9e, 0x53, 0x8a, 0x9d, 0x0b, 0x2d, 0xbf, 0xd9,
	0xba, 0xb9, 0xab, 0x97, 0x6f, 0xef, 0xea, 0xe5, 0x6f, 0x77, 0xf5, 0xf2, 0xe7, 0xfb, 0x7a, 0xe9,
	0xf6, 0xbe, 0x5e, 0xfa, 0x72, 0x5f, 0x2f, 0x7d, 0xb4, 0x82, 0x50, 0x74, 0x2f, 0x5d, 0xd3, 0xe3,
	0xb1, 0xf5, 0xc4, 0x33, 0x79, 0xd5, 0xb0, 0xfa, 0xbf, 0xbc, 0x95, 0xee, 0xa4, 0x7a, 0x2c, 0x9f,
	0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x22, 0x12, 0x86, 0x73, 0x0e, 0x07, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LightNodeSaleContracts) > 0 {
		for iNdEx := len(m.LightNodeSaleContracts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LightNodeSaleContracts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.BridgeTaxes) > 0 {
		for iNdEx := len(m.BridgeTaxes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BridgeTaxes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.BridgeTransferLimits) > 0 {
		for iNdEx := len(m.BridgeTransferLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BridgeTransferLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.UnbatchedTransfers) > 0 {
		for iNdEx := len(m.UnbatchedTransfers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbatchedTransfers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Erc20ToDenoms) > 0 {
		for iNdEx := len(m.Erc20ToDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Erc20ToDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Attestations) > 0 {
		for iNdEx := len(m.Attestations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attestations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BatchGasEstimates) > 0 {
		for iNdEx := len(m.BatchGasEstimates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchGasEstimates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BatchConfirms) > 0 {
		for iNdEx := len(m.BatchConfirms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchConfirms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SkywayNonces) > 0 {
		for iNdEx := len(m.SkywayNonces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SkywayNonces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SkywayNonces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SkywayNonces) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SkywayNonces) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainReferenceId) > 0 {
		i -= len(m.ChainReferenceId)
		copy(dAtA[i:], m.ChainReferenceId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainReferenceId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.LastBatchId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastBatchId))
		i--
		dAtA[i] = 0x20
	}
	if m.LastTxPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastTxPoolId))
		i--
		dAtA[i] = 0x18
	}
	if m.LastSlashedBatchBlock != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSlashedBatchBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.LastObservedNonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastObservedNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.SkywayNonces) > 0 {
		for _, e := range m.SkywayNonces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BatchConfirms) > 0 {
		for _, e := range m.BatchConfirms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BatchGasEstimates) > 0 {
		for _, e := range m.BatchGasEstimates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Attestations) > 0 {
		for _, e := range m.Attestations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Erc20ToDenoms) > 0 {
		for _, e := range m.Erc20ToDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbatchedTransfers) > 0 {
		for _, e := range m.UnbatchedTransfers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BridgeTransferLimits) > 0 {
		for _, e := range m.BridgeTransferLimits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BridgeTaxes) > 0 {
		for _, e := range m.BridgeTaxes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LightNodeSaleContracts) > 0 {
		for _, e := range m.LightNodeSaleContracts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *SkywayNonces) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastObservedNonce != 0 {
		n += 1 + sovGenesis(uint64(m.LastObservedNonce))
	}
	if m.LastSlashedBatchBlock != 0 {
		n += 1 + sovGenesis(uint64(m.LastSlashedBatchBlock))
	}
	if m.LastTxPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.LastTxPoolId))
	}
	if m.LastBatchId != 0 {
		n += 1 + sovGenesis(uint64(m.LastBatchId))
	}
	l = len(m.ChainReferenceId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkywayNonces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SkywayNonces = append(m.SkywayNonces, SkywayNonces{})
			if err := m.SkywayNonces[len(m.SkywayNonces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Batches = append(m.Batches, OutgoingTxBatch{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchConfirms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchConfirms = append(m.BatchConfirms, MsgConfirmBatch{})
			if err := m.BatchConfirms[len(m.BatchConfirms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchGasEstimates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchGasEstimates = append(m.BatchGasEstimates, MsgEstimateBatchGas{})
			if err := m.BatchGasEstimates[len(m.BatchGasEstimates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attestations = append(m.Attestations, Attestation{})
			if err := m.Attestations[len(m.Attestations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20ToDenoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20ToDenoms = append(m.Erc20ToDenoms, ERC20ToDenom{})
			if err := m.Erc20ToDenoms[len(m.Erc20ToDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbatchedTransfers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbatchedTransfers = append(m.UnbatchedTransfers, OutgoingTransferTx{})
			if err := m.UnbatchedTransfers[len(m.UnbatchedTransfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeTransferLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeTransferLimits = append(m.BridgeTransferLimits, &BridgeTransferLimit{})
			if err := m.BridgeTransferLimits[len(m.BridgeTransferLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeTaxes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeTaxes = append(m.BridgeTaxes, &BridgeTax{})
			if err := m.BridgeTaxes[len(m.BridgeTaxes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LightNodeSaleContracts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LightNodeSaleContracts = append(m.LightNodeSaleContracts, &LightNodeSaleContract{})
			if err := m.LightNodeSaleContracts[len(m.LightNodeSaleContracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *SkywayNonces) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: SkywayNonces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SkywayNonces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedNonce", wireType)
			}
			m.LastObservedNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastObservedNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSlashedBatchBlock", wireType)
			}
			m.LastSlashedBatchBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSlashedBatchBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTxPoolId", wireType)
			}
			m.LastTxPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastTxPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBatchId", wireType)
			}
			m.LastBatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBatchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

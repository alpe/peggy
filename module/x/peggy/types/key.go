package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "peggy"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey is the module name router key
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)

var (
	EthAddressKey    = []byte{0x1}
	ValsetRequestKey = []byte{0x2}
	ValsetConfirmKey = []byte{0x3}

	OutgoingTXPoolKey           = []byte{0x5}
	SequenceKeyPrefix           = []byte{0x6}
	DenomiatorPrefix            = []byte{0x7}
	SecondIndexOutgoingTXFeeKey = []byte{0x8}

	// sequence keys
	KeyLastTXPoolID        = append(SequenceKeyPrefix, []byte("lastTxPoolId")...)
	KeyLastOutgoingBatchID = append(SequenceKeyPrefix, []byte("lastBatchId")...)
)

func GetEthAddressKey(validator sdk.AccAddress) []byte {
	return append(EthAddressKey, []byte(validator)...)
}

func GetValsetRequestKey(nonce int64) []byte {
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, uint64(nonce))

	return append(ValsetRequestKey, nonceBytes...)
}

func GetValsetConfirmKey(nonce int64, validator sdk.AccAddress) []byte {
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, uint64(nonce))

	return append(ValsetConfirmKey, append(nonceBytes, []byte(validator)...)...)
}

func GetOutgoingTxPoolKey(id uint64) []byte {
	return append(OutgoingTXPoolKey, sdk.Uint64ToBigEndian(id)...)
}

func GetFeeSecondIndexKey(fee sdk.Coin) []byte {
	assertPeggyVoucher(fee)

	r := make([]byte, 1+VoucherDenomLen+8)
	copy(r[0:], SecondIndexOutgoingTXFeeKey)
	voucherDenom, _ := AsVoucherDenom(fee.Denom)
	copy(r[len(SecondIndexOutgoingTXFeeKey):], voucherDenom.Unprefixed())
	copy(r[len(SecondIndexOutgoingTXFeeKey)+len(voucherDenom.Unprefixed()):], sdk.Uint64ToBigEndian(fee.Amount.Uint64()))
	return r
}

func GetDenominatorKey(voucherDenominator string) []byte {
	return append(DenomiatorPrefix, []byte(voucherDenominator)...)
}

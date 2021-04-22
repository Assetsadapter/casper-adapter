package casper

import (
	"encoding/hex"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openwallet"
	"strings"
)

var (
	alphabet   = addressEncoder.BTCAlphabet
	ssPrefix   = []byte{0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45}
	encodeType = "base58"
)

var (
	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
	wm *WalletManager
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2(wm *WalletManager) *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	decoder.wm = wm
	return &decoder
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	return nil, nil
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(pub []byte, opts ...interface{}) (string, error) {
	if len(pub) != 32 {
		pub, _ = owcrypt.CURVE25519_convert_Ed_to_X(pub)
	}
	//split,_:=hex.DecodeString("00")
	//prefix := append([]byte("ed25519"),split...)
	//pubKeyBytesAll := append(prefix, pub...)
	//pkHash := owcrypt.Hash(pubKeyBytesAll,32 , owcrypt.HASH_ALG_BLAKE2B)
	return "01" + hex.EncodeToString(pub), nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	if len(address) == 66 && strings.HasPrefix(address, "01") {
		_, err := hex.DecodeString(address)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

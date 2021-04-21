package txsigner

import (
	"errors"
	"github.com/blocktree/go-owcrypt"
)

var Default = &TransactionSigner{}

type TransactionSigner struct {
}

// SignTransactionHash 交易哈希签名算法
// required
func (singer *TransactionSigner) SignTransactionHash(msg []byte, privateKey []byte, eccType uint32) ([]byte, error) {
	signature, _, retCode := owcrypt.Signature(privateKey, nil, msg, owcrypt.ECC_CURVE_ED25519)
	if retCode != owcrypt.SUCCESS {
		return nil, errors.New("sign failed")
	}
	return signature, nil
}

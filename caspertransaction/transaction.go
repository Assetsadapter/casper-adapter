package caspertransaction

import (
	"encoding/hex"
	"errors"
	"github.com/blocktree/go-owcrypt"
)

//https://docs.casperlabs.io/en/latest/implementation/serialization-standard.html
type Deploy struct {
	Approvals []Approvals
	Header    DeployHeader
	Hash      []byte
	Payment   Payment
	Session   Transfer
}

type DeployHeader struct {
	Account      string
	Timestamp    uint64
	Ttl          uint64
	GasPrice     uint64
	BodyHash     []byte
	Dependencies [][]byte
	ChainName    string
}

type Approvals struct {
	Signer    string
	signature string
}
type Payment struct {
	Amount uint64
}
type Transfer struct {
	Amount     uint64
	To         string //public key hex
	SourceUref string
	TransferId uint64
}

func NewDeploy(payAmount, transAmount, timeStamp, gasPrice uint64, fromAccount, toAccount, chainName string) (*Deploy, error) {
	payment := Payment{Amount: payAmount}
	paymentBytes, err := payment.toBytes()
	if err != nil {
		return nil, err
	}
	trans := Transfer{To: toAccount, Amount: transAmount}
	transBytes, err := trans.toBytes()
	if err != nil {
		return nil, err
	}
	var deployBodyBytes []byte
	deployBodyBytes = append(deployBodyBytes, paymentBytes...)
	deployBodyBytes = append(deployBodyBytes, transBytes...)
	deployBodyHash := owcrypt.Hash(deployBodyBytes, 32, owcrypt.HASH_ALG_BLAKE2B)

	deployHeader := DeployHeader{Account: fromAccount, Timestamp: timeStamp, Ttl: 0, GasPrice: gasPrice, BodyHash: deployBodyHash, ChainName: chainName}
	deployHeaderBytes, err := deployHeader.toBytes()
	if err != nil {
		return nil, err
	}
	deployHeaderHash := owcrypt.Hash(deployHeaderBytes, 32, owcrypt.HASH_ALG_BLAKE2B)
	deploy := &Deploy{Header: deployHeader, Session: trans, Payment: payment, Hash: deployHeaderHash}
	return deploy, nil
}

//deployHeader 序列化
func (deployHeader *DeployHeader) toBytes() ([]byte, error) {
	var bytesData []byte
	//tag is 1
	bytesData = append(bytesData, byte(2))
	//public key bytes
	acountPublicKeyBytes, err := hex.DecodeString(deployHeader.Account)
	if err != nil {
		return nil, err
	}
	bytesData = append(bytesData, acountPublicKeyBytes...)

	//timestamp
	bytesData = append(bytesData, uint64ToLittleEndianBytes(deployHeader.Timestamp)...)

	//gasPrice
	bytesData = append(bytesData, uint64ToLittleEndianBytes(deployHeader.GasPrice)...)

	//ttl
	bytesData = append(bytesData, uint64ToLittleEndianBytes(deployHeader.Ttl)...)

	//body hash
	bytesData = append(bytesData, deployHeader.BodyHash...)

	//dependencies
	bytesData = append(bytesData, []byte{0, 0, 0, 0}...)

	//length of chainName String
	bytesData = append(bytesData, uint32ToLittleEndianBytes(uint32(len(deployHeader.ChainName)))...)
	//Amount string
	bytesData = append(bytesData, []byte(deployHeader.ChainName)...)

	return nil, nil
}

// payment 序列化
func (payment *Payment) toBytes() ([]byte, error) {
	var bytesData []byte
	//tag
	bytesData = append(bytesData, byte(0))

	//modoule bytes
	bytesData = append(bytesData, []byte{0, 0, 0, 0}...)
	//length of args 只有1个参数可用
	bytesData = append(bytesData, uint32ToLittleEndianBytes(1)...)
	//Amount
	//length of "Amount" String
	bytesData = append(bytesData, uint32ToLittleEndianBytes(6)...)
	//Amount string
	bytesData = append(bytesData, []byte("amount")...)
	//Amount number 512 bit little endian Byte
	amountBytes := uintToShortByte(payment.Amount)
	bytesData = append(bytesData, uint32ToLittleEndianBytes(uint32(len(amountBytes)))...)
	bytesData = append(bytesData, amountBytes...)
	//Amount u512 tag = 8
	bytesData = append(bytesData, byte(8))

	return bytesData, nil
}

// transfer 序列化
func (transfer *Transfer) toBytes() ([]byte, error) {
	var bytesData []byte

	//tag
	bytesData = append(bytesData, byte(5))

	//length of args 只有3个参数可用
	bytesData = append(bytesData, uint32ToLittleEndianBytes(3)...)

	//length of "Amount" String
	bytesData = append(bytesData, uint32ToLittleEndianBytes(6)...)
	//Amount string
	bytesData = append(bytesData, []byte("amount")...)
	//Amount number 512 bit little endian Byte
	amountBytes := uintToShortByte(transfer.Amount)
	bytesData = append(bytesData, uint32ToLittleEndianBytes(uint32(len(amountBytes)))...)
	bytesData = append(bytesData, amountBytes...)
	//Amount u512 tag = 8
	bytesData = append(bytesData, byte(8))

	//target
	toPublicBytes, err := hex.DecodeString(transfer.To)
	if err != nil {
		return nil, err
	}
	if len(toPublicBytes) != 32 {
		return nil, errors.New("transfer serialize error, wrong public key len")
	}

	//length of "target" String
	bytesData = append(bytesData, uint32ToLittleEndianBytes(6)...)
	//target string
	bytesData = append(bytesData, []byte("target")...)
	//public key len
	bytesData = append(bytesData, uint32ToLittleEndianBytes(32)...)
	//public key
	bytesData = append(bytesData, toPublicBytes...)
	//public key tag  =  15
	bytesData = append(bytesData, byte(15))
	//public key size 32
	bytesData = append(bytesData, uint32ToLittleEndianBytes(32)...)

	//length of "id" String
	bytesData = append(bytesData, uint32ToLittleEndianBytes(2)...)
	//id string
	bytesData = append(bytesData, []byte("id")...)
	// left bytes fixed
	bytesData = append(bytesData, []byte{1, 0, 0, 0, 0, 13, 5}...)

	return bytesData, nil

}

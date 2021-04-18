package caspertransaction

import (
	"encoding/hex"
	"errors"
)

//https://docs.casperlabs.io/en/latest/implementation/serialization-standard.html
type Deploy struct {
	Approvals []Approvals
	Header    DeployHeader
	Hash      string
	Payment   string
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

type Transfer struct {
	Amount     uint64
	To         string //public key hex
	SourceUref string
	TransferId uint64
}

//序列化
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

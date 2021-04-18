package caspertransaction

import (
	"bytes"
	"encoding/binary"
	"github.com/prometheus/common/log"
	"testing"
)

//{
//"0": 3,
//"1": 199,
//"2": 149,
//"3": 3
//}
func Test_uint32(t *testing.T) {
	//n:= uintToShortByte(33892232)
	//log.Info(n)
	log.Info(uint32ToLittleEndianBytes(2))
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil
	}

	return buff.Bytes()
}

func Test_transferSer(t *testing.T) {
	trans := Transfer{}
	trans.To = "d9bf2148748a85c89da5aad8ee0b0fc2d105fd39d41a4c796536354f0ae2900c"
	trans.Amount = 33892232
	b, _ := trans.toBytes()
	log.Info("len=", len(b), "data=", b)
}

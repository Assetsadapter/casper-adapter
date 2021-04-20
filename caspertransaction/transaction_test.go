package caspertransaction

import (
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
	log.Info(uint64ToLittleEndianBytes(1618833271075))
}

func Test_transferSer(t *testing.T) {
	trans := Transfer{}
	trans.To = "d9bf2148748a85c89da5aad8ee0b0fc2d105fd39d41a4c796536354f0ae2900c"
	trans.Amount = 33892232
	b, _ := trans.toBytes()
	log.Info("transfer len=", len(b), "data=", b)
}

func Test_payment(t *testing.T) {
	payment := Payment{}
	payment.Amount = 10000000000000
	b, _ := payment.toBytes()
	log.Info("payment len=", len(b), "data=", b)
}

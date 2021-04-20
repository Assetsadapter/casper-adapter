package caspertransaction

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"testing"
	"time"
)

//{
//"0": 3,
//"1": 199,
//"2": 149,
//"3": 3
//}
func Test_uint32(t *testing.T) {
	//log.Info(uintToShortByte(10000))
	data := time.Unix(1618931725700/1000, 0)
	fmt.Println(data.UTC().Format(time.RFC3339))

}

func Test_transferSer(t *testing.T) {
	trans := Transfer{}
	trans.To = "d9bf2148748a85c89da5aad8ee0b0fc2d105fd39d41a4c796536354f0ae2900c"
	trans.Amount = 33892232
	b, _ := trans.toBytes()
	log.Info("transfer len=", len(b), "data=", b)
	transJson, _ := trans.toJson()
	log.Info("transfer json=", transJson)
}

func Test_payment(t *testing.T) {
	payment := Payment{}
	payment.Amount = 10000000000000
	paymentJson, _ := payment.toJson()
	j, _ := json.Marshal(paymentJson)
	log.Info("payment json=", string(j))
}

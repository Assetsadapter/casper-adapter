package caspertransaction

import (
	"encoding/json"
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
	i := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info(i)
	log.Info(time.Now().Unix())
	//a:=uint64(1619030752862)

}

func Test_transferSer(t *testing.T) {
	trans := Transfer{}
	trans.To = "01322ef12cbb08749b2160743ec11f7ff34b96feadeecfe356c75b364a6b514cba"
	trans.Amount = 2500000000
	b, _ := trans.toBytes()
	log.Info("transfer len=", len(b), "data=", b)
	transJson, _ := trans.toJson()
	log.Info("transfer json=", transJson)
}

func Test_payment(t *testing.T) {
	payment := Payment{}
	payment.Amount = 10000000000000
	payBytes, _ := payment.toBytes()
	log.Info("payBytes =", payBytes)
	paymentJson, _ := payment.ToJson()
	j, _ := json.Marshal(paymentJson)
	log.Info("payment json=", string(j))
}

func Test_newDeploy(t *testing.T) {
	//timeStamp := uint64(time.Now().Unix())
	timeStamp := uint64(1619024183064)
	ttl := uint64(1800000)
	from := "013697f07afdb5e28d774ed0166adc429ea815eb7cfde38aec95119ff1c7b356d3"
	to := "01c85fa6c3c9a0bb23f19c5e2b3f4c76e5fb23793f14ccc815f7752ba1d3f45aaa"
	chainName := "casper-test"
	deploy, err := NewDeploy(10000000000000, 2500000000, timeStamp, 1, ttl, from, to, chainName)
	if err != nil {
		t.Fatal(err)
	}
	log.Error(deploy.Hash)
	djson, _ := deploy.ToJson()
	j, _ := json.Marshal(djson)
	log.Info("transfer json=", string(j))
}

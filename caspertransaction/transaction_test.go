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
	paymentJson, _ := payment.ToJson()
	j, _ := json.Marshal(paymentJson)
	log.Info("payment json=", string(j))
}

func Test_newDeploy(t *testing.T) {
	timeStamp := uint64(time.Now().Unix())
	ttl := uint64(30 * 60 * 1000)
	from := "01664adcf74db3887accb10af5dccb8e3c2a6b6d33f900ffa69cb42b356aa2ca52"
	to := "01322ef12cbb08749b2160743ec11f7ff34b96feadeecfe356c75b364a6b514cba"
	chainName := "casper-testnet"
	deploy, err := NewDeploy(10000000000000, 33892232, timeStamp, 1, ttl, from, to, chainName)
	if err != nil {
		t.Fatal(err)
	}
	djson, _ := deploy.ToJson()
	j, _ := json.Marshal(djson)
	log.Info("transfer json=", string(j))
}

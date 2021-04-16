package casper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blocktree/openwallet/v2/log"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	testNodeAPIHttp = "http://3.208.91.63:8888"
	testNodeAPIRpc  = "http://3.208.91.63:7777/rpc"
)

func PrintJsonLog(t *testing.T, logCont string) {
	if strings.HasPrefix(logCont, "{") {
		var str bytes.Buffer
		_ = json.Indent(&str, []byte(logCont), "", "    ")
		t.Logf("Get Call Result return: \n\t%+v\n", str.String())
	} else {
		t.Logf("Get Call Result return: \n\t%+v\n", logCont)
	}
}

func TestGetCall(t *testing.T) {
	tw := NewClient(testNodeAPIRpc, true)

	if r, err := tw.GetCall("/metadata/"); err != nil {
		t.Errorf("Get Call Result failed: %v\n", err)
	} else {
		PrintJsonLog(t, r.String())
	}
}

func Test_getBlockHeight(t *testing.T) {

	c := NewClient(testNodeAPIHttp, true)

	r, err := c.getLastBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("height:", r)
	}

}

func Test_getBlockByHeight(t *testing.T) {
	c := NewClient(testNodeAPIRpc, true)
	r, err := c.getBlockByHeight(1830393)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
func Test_clientRpc(t *testing.T) {
	c := NewClient(testNodeAPIRpc, true)
	data := make([]interface{}, 0)
	dataResp, err := c.RpcCall("info_get_status", data)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Info(dataResp.String())

}

func Test_getBlockRpc(t *testing.T) {
	c := NewClient(testNodeAPIRpc, true)
	dataResp, err := c.getBlockByHeight(5210)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Info(dataResp)
}

func Test_getBalance(t *testing.T) {

	c := NewClient(testNodeAPIRpc, true)
	address := "CyVGZAGjfD9bbQcN7Ja7cFrazM4yAzawRhKcWYN1ZmMrpwf"

	r, err := c.getBalance(address, true, 20000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	address = "J5cTdWxoMRZyQHyvvDMxB6dp7YitNpzEkj3ZrJFsGmARcC2"

	r, err = c.getBalance(address, true, 20000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_sendTransaction(t *testing.T) {
	c := NewClient(testNodeAPIRpc, true)
	r, err := c.sendTransaction("0x39028453538d40098561df3a2fe577c2995d7f6a5bd45f0f5708e9b9e11cc4896e1db800ba6d5" +
		"4c43a5c07e6d1b89ff51d356ac686b0bcd592d00a5140c4842a054c567830aad78e0bd03c86792e794756a516507b1911f824267d51b75" +
		"b3676d0a9b40295030000050009d0dbc83629dcd90f7e1cc989cd2cd713205adff4a3b0f2699e31b7e35981340700dc5c2402")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
func Test_rpc(t *testing.T) {
	// test (POST http://3.208.91.63:7777/rpc)
	//	Content-Length: 67
	//Accept: application/json
	//	Content-Type: application/json
	//	Accept-Encoding: gzip

	//{"id":1,"jsonrpc":"2.0","method":"info_get_status","params":[null]}
	json := []byte(`{"jsonrpc": "2.0","id": 123123,"method": "info_get_status","params": []}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "http://3.208.91.63:7777/rpc", body)

	// Headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

}

/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package casper

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/tidwall/gjson"
	"testing"

	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/pborman/uuid"
)

// func TestONTBlockScanner_GetCurrentBlockHeight(t *testing.T) {
// 	bs := NewONTBlockScanner(tw)
// 	header, _ := bs.GetCurrentBlockHeader()
// 	t.Logf("GetCurrentBlockHeight height = %d \n", header.Height)
// 	t.Logf("GetCurrentBlockHeight hash = %v \n", header.Hash)
// }

func TestGetCurrentBlockHeight(t *testing.T) {
	header, _ := tw.Blockscanner.GetCurrentBlockHeader()
	fmt.Println(header)
}

func TestGetBlockHeight(t *testing.T) {
	height, _ := tw.GetBlockHeight()
	t.Logf("GetBlockHeight height = %d \n", height)
}

func TestGetTransaction(t *testing.T) {
	// raw, err := tw.GetTransaction("7792d54a7eb9f467de7a1292c0d317b2d5462e7ec35d229a383d948c18d9c873")
	// if err != nil {
	// 	t.Errorf("GetTransaction failed unexpected error: %v\n", err)
	// 	return
	// }

	// t.Logf("BlockHash = %v \n", raw.BlockHash)
	// t.Logf("BlockHeight = %v \n", raw.BlockHeight)
	// t.Logf("Blocktime = %v \n", raw.Blocktime)
	// t.Logf("Fees = %v \n", raw.Fees)

	// t.Logf("========= vins ========= \n")

	// for i, vin := range raw.Vins {
	// 	t.Logf("TxID[%d] = %v \n", i, vin.TxID)
	// 	t.Logf("Vout[%d] = %v \n", i, vin.Vout)
	// 	t.Logf("Addr[%d] = %v \n", i, vin.Addr)
	// 	t.Logf("Value[%d] = %v \n", i, vin.Value)
	// }

	// t.Logf("========= vouts ========= \n")

	// for i, out := range raw.Vouts {
	// 	t.Logf("ScriptPubKey[%d] = %v \n", i, out.ScriptPubKey)
	// 	t.Logf("Addr[%d] = %v \n", i, out.Addr)
	// 	t.Logf("Value[%d] = %v \n", i, out.Value)
	// }
}

func TestGetTxIDsInMemPool(t *testing.T) {
	txids, err := tw.GetTxIDsInMemPool()
	if err != nil {
		t.Errorf("GetTxIDsInMemPool failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetTxIDsInMemPool = %v \n", txids)
}

// func TestONTBlockScanner_scanning(t *testing.T) {

// 	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
// 	//address := "miphUAzHHeM1VXGSFnw6owopsQW3jAQZAk"

// 	//wallet, err := tw.GetWalletInfo(accountID)
// 	//if err != nil {
// 	//	t.Errorf("ONTBlockScanner_scanning failed unexpected error: %v\n", err)
// 	//	return
// 	//}

// 	bs := NewONTBlockScanner(tw)

// 	//bs.DropRechargeRecords(accountID)

// 	bs.SetRescanBlockHeight(1384586)
// 	//tw.SaveLocalNewBlock(1355030, "00000000000000125b86abb80b1f94af13a5d9b07340076092eda92dade27686")

// 	//bs.AddAddress(address, accountID)

// 	bs.ScanBlockTask()
// }

// func TestONTBlockScanner_Run(t *testing.T) {

// 	var (
// 		endRunning = make(chan bool, 1)
// 	)

// 	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
// 	//address := "msnYsBdBXQZqYYqNNJZsjShzwCx9fJVSin"

// 	//wallet, err := tw.GetWalletInfo(accountID)
// 	//if err != nil {
// 	//	t.Errorf("ONTBlockScanner_Run failed unexpected error: %v\n", err)
// 	//	return
// 	//}

// 	bs := NewONTBlockScanner(tw)

// 	//bs.DropRechargeRecords(accountID)

// 	//bs.SetRescanBlockHeight(1384586)

// 	//bs.AddAddress(address, accountID)

// 	bs.Run()

// 	<-endRunning

// }

func TestONTBlockScanner_ScanBlock(t *testing.T) {

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "msnYsBdBXQZqYYqNNJZsjShzwCx9fJVSin"

	bs := tw.Blockscanner
	//bs.AddAddress(address, accountID)
	bs.ScanBlock(1739242)
}

func TestONTBlockScanner_ExtractTransaction(t *testing.T) {

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "msHemmfSZ3au6h9S1annGcTGrTVryRbSFV"

	//bs := tw.Blockscanner
	////bs.AddAddress(address, accountID)
	//bs.ExtractTransaction(
	//	1435497,
	//	"00000000e271b8234ed2271cb80f1cf2701469a4e02b0536fdce4f4306ff7852",
	//	"c550ae3ffafdda46c13217797dd0aa8ee870727d3e8cab1551d6a3f5e3f7ace0", bs.GetSourceKeyByAddress)

}

func TestWallet_GetRecharges(t *testing.T) {
	accountID := "WFvvr5q83WxWp1neUMiTaNuH7ZbaxJFpWu"
	wallet, err := tw.GetWalletInfo(accountID)
	if err != nil {
		t.Errorf("GetRecharges failed unexpected error: %v\n", err)
		return
	}

	recharges, err := wallet.GetRecharges(false)
	if err != nil {
		t.Errorf("GetRecharges failed unexpected error: %v\n", err)
		return
	}

	t.Logf("recharges.count = %v", len(recharges))
	//for _, r := range recharges {
	//	t.Logf("rechanges.count = %v", len(r))
	//}
}

//func TestONTBlockScanner_DropRechargeRecords(t *testing.T) {
//	accountID := "W4ruoAyS5HdBMrEeeHQTBxo4XtaAixheXQ"
//	bs := NewONTBlockScanner(tw)
//	bs.DropRechargeRecords(accountID)
//}

//func TestGetUnscanRecords(t *testing.T) {
//	list, err := tw.GetUnscanRecords()
//	if err != nil {
//		t.Errorf("GetUnscanRecords failed unexpected error: %v\n", err)
//		return
//	}
//
//	for _, r := range list {
//		t.Logf("GetUnscanRecords unscan: %v", r)
//	}
//}

// func TestONTBlockScanner_RescanFailedRecord(t *testing.T) {
// 	bs := NewONTBlockScanner(tw)
// 	bs.RescanFailedRecord()
// }

func TestFullAddress(t *testing.T) {

	dic := make(map[string]string)
	for i := 0; i < 20000000; i++ {
		dic[uuid.NewUUID().String()] = uuid.NewUUID().String()
	}
}

func TestONTBlockScanner_GetTransactionsByAddress(t *testing.T) {
	coin := openwallet.Coin{
		Symbol:     "BTC",
		IsContract: false,
	}
	txExtractDatas, err := tw.Blockscanner.GetTransactionsByAddress(0, 50, coin, "2N7Mh6PLX39japSF76r2MAf7wT7WKU5TdpK")
	if err != nil {
		t.Errorf("GetTransactionsByAddress failed unexpected error: %v\n", err)
		return
	}

	for _, ted := range txExtractDatas {
		t.Logf("tx = %v", ted.Transaction)
	}

}
func TestNewBlock(t *testing.T) {
	blockStr := "{\"number\":\"4560516\",\"hash\":\"0x2cc6e5165c9f65860e328abeb74fd6630c0fc931ddae7117baca8ce319e0a7e8\",\"parentHash\":\"0x6acf34c7fb05f00599cacf92dedd87dec41a0e2eb68ab76056babae4dfd5be9a\",\"stateRoot\":\"0x73c66118430deedad33deb0c9ee8991cb65e2ff1b5dcc1f52a3d9e4e7808d7be\",\"extrinsicsRoot\":\"0x97800ea27600528e3466594d4b773626401f7c512e6e2c03972b0ede8eafb92e\",\"authorId\":\"1zugcaiwmKdWsfuubmCMBgKKMLSef2TEC3Gfvv5GxLGTKMN\",\"logs\":[{\"type\":\"PreRuntime\",\"index\":\"6\",\"value\":[\"BABE\",\"0x039800000091c212100000000060320ae5c2362334f0021e7296b37a89455eddf4ef97bf90966653d7a8793a625990ef902b1cc4af0f2709ad6f3e597eb77fa8b496cc759e5697fbb85f14480823d31afac9a4a80194b86a9c6d7af7072c297ee957d08cd28142f527d1c93107\"]},{\"type\":\"Seal\",\"index\":\"5\",\"value\":[\"BABE\",\"0xf29bb2a0aae823d2499697529e5f6966ae8bd6210cd99ba0f6323e34b24d16776d2cb6b2f2159a98d83bf2d9452f92a3f0b31f7704e8b0332400c2b1a0724887\"]}],\"onInitialize\":{\"events\":[]},\"extrinsics\":[{\"method\":{\"pallet\":\"timestamp\",\"method\":\"set\"},\"signature\":null,\"nonce\":null,\"args\":{\"now\":\"1617989478000\"},\"tip\":null,\"hash\":\"0x92bb00448f7088c8c09f90ea38a56fe4c73e92c8c7f8d7a7e2251f9c41e8c031\",\"info\":{},\"events\":[{\"method\":{\"pallet\":\"system\",\"method\":\"ExtrinsicSuccess\"},\"data\":[{\"weight\":\"186153000\",\"class\":\"Mandatory\",\"paysFee\":\"Yes\"}]}],\"success\":true,\"paysFee\":false},{\"method\":{\"pallet\":\"utility\",\"method\":\"batch\"},\"signature\":{\"signature\":\"0xd1c63e8129865e033662e91b5d37d726720b4076ef7661ef04c4adc5e5e0fdbdb03530956d177dbdef4f99c1330b8bb8be117a764c2039b8ddf082c32540530c\",\"signer\":{\"id\":\"13oRgXk8cHVxpi6FHxycniEz4ggRzt7K3rr8P2D79UC8abE2\"}},\"nonce\":\"18392\",\"args\":{\"calls\":[{\"method\":{\"pallet\":\"balances\",\"method\":\"transferKeepAlive\"},\"args\":{\"dest\":{\"id\":\"12qAanZRNsLK8KJE7L17aVZs5K8MmubtFR3qpYEhUdm981XU\"},\"value\":\"281866000000\"}},{\"method\":{\"pallet\":\"balances\",\"method\":\"transferKeepAlive\"},\"args\":{\"dest\":{\"id\":\"1298KLZYMpF34d73Txg8X35Z79ERtvMWaj1xfAN6xR328Hcn\"},\"value\":\"495200000000\"}}]},\"tip\":\"0\",\"hash\":\"0x2ba93fc540e293b17865779c56af2e1c372096410a07f9443a7f1825fcb3344c\",\"info\":{\"weight\":\"385843000\",\"class\":\"Normal\",\"partialFee\":\"203000030\"},\"events\":[{\"method\":{\"pallet\":\"system\",\"method\":\"NewAccount\"},\"data\":[\"12qAanZRNsLK8KJE7L17aVZs5K8MmubtFR3qpYEhUdm981XU\"]},{\"method\":{\"pallet\":\"balances\",\"method\":\"Endowed\"},\"data\":[\"12qAanZRNsLK8KJE7L17aVZs5K8MmubtFR3qpYEhUdm981XU\",\"281866000000\"]},{\"method\":{\"pallet\":\"balances\",\"method\":\"Transfer\"},\"data\":[\"13oRgXk8cHVxpi6FHxycniEz4ggRzt7K3rr8P2D79UC8abE2\",\"12qAanZRNsLK8KJE7L17aVZs5K8MmubtFR3qpYEhUdm981XU\",\"281866000000\"]},{\"method\":{\"pallet\":\"system\",\"method\":\"NewAccount\"},\"data\":[\"1298KLZYMpF34d73Txg8X35Z79ERtvMWaj1xfAN6xR328Hcn\"]},{\"method\":{\"pallet\":\"balances\",\"method\":\"Endowed\"},\"data\":[\"1298KLZYMpF34d73Txg8X35Z79ERtvMWaj1xfAN6xR328Hcn\",\"495200000000\"]},{\"method\":{\"pallet\":\"balances\",\"method\":\"Transfer\"},\"data\":[\"13oRgXk8cHVxpi6FHxycniEz4ggRzt7K3rr8P2D79UC8abE2\",\"1298KLZYMpF34d73Txg8X35Z79ERtvMWaj1xfAN6xR328Hcn\",\"495200000000\"]},{\"method\":{\"pallet\":\"utility\",\"method\":\"BatchCompleted\"},\"data\":[]},{\"method\":{\"pallet\":\"balances\",\"method\":\"Deposit\"},\"data\":[\"1zugcaiwmKdWsfuubmCMBgKKMLSef2TEC3Gfvv5GxLGTKMN\",\"203000030\"]},{\"method\":{\"pallet\":\"system\",\"method\":\"ExtrinsicSuccess\"},\"data\":[{\"weight\":\"385843000\",\"class\":\"Normal\",\"paysFee\":\"Yes\"}]}],\"success\":true,\"paysFee\":true}],\"onFinalize\":{\"events\":[]},\"finalized\":true}"
	result := gjson.ParseBytes([]byte(blockStr))
	block := NewBlock(&result)
	log.Info(block)
}

//func TestGetLocalBlock(t *testing.T) {
//	db, err := storm.Open(filepath.Join(tw.Config.dbPath, tw.Config.BlockchainFile))
//	if err != nil {
//		return
//	}
//	defer db.Close()
//
//	var blocks []*Block
//	err = db.All(&blocks)
//	if err != nil {
//		log.Error("no find")
//		return
//	}
//	log.Info("blocks = ", len(blocks))
//}

// func Test_tmp(t *testing.T) {
// 	addr := "AYmuoVvtCojm1F3ATMf2fNww3wBNvAxbi5"
// 	params := []interface{}{addr}
// 	//c := NewRpcClient("http://1.1.1.1:12345")
// 	c := NewRpcClient("http://192.168.27.124:20336")

// 	txid, err := c.sendRpcRequest("0", "getbalance", params)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(txid))
// }

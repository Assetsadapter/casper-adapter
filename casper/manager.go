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
	"encoding/hex"
	"errors"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openw"
	"path/filepath"
	"strings"

	"github.com/blocktree/openwallet/v2/hdkeystore"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

type WalletManager struct {
	openwallet.AssetsAdapterBase

	Storage         *hdkeystore.HDKeystore //秘钥存取
	ApiClient       *ApiClient
	Config          *WalletConfig                 //钱包管理配置
	WalletsInSum    map[string]*openwallet.Wallet //参与汇总的钱包
	Blockscanner    *CSPRBlockScanner             //区块扫描器
	Decoder         openwallet.AddressDecoderV2   //地址编码器
	TxDecoder       openwallet.TransactionDecoder //交易单编码器
	Log             *log.OWLogger                 //日志工具
	ContractDecoder *ContractDecoder              //智能合约解析器
	wrapper         *openw.WalletWrapper
}

func NewWalletManager(wrapper *openw.WalletWrapper) *WalletManager {
	wm := WalletManager{}
	wm.Config = NewConfig(Symbol, MasterKey)
	storage := hdkeystore.NewHDKeystore(wm.Config.keyDir, hdkeystore.StandardScryptN, hdkeystore.StandardScryptP)
	wm.Storage = storage
	//参与汇总的钱包
	wm.WalletsInSum = make(map[string]*openwallet.Wallet)
	//区块扫描器
	wm.wrapper = wrapper
	wm.Blockscanner = NewCSPRBlockScanner(&wm)
	wm.Decoder = NewAddressDecoderV2(&wm)
	wm.TxDecoder = NewTransactionDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.ContractDecoder = NewContractDecoder(&wm)
	//	wm.RPCClient = NewRpcClient("http://localhost:20336/")
	return &wm
}

//GetWalletInfo 获取钱包列表
func (wm *WalletManager) GetWalletInfo(walletID string) (*openwallet.Wallet, error) {

	wallets, err := wm.GetWallets()
	if err != nil {
		return nil, err
	}

	//获取钱包余额
	for _, w := range wallets {
		if w.WalletID == walletID {
			return w, nil
		}

	}

	return nil, errors.New("The wallet that your given name is not exist!")
}

//GetWallets 获取钱包列表
func (wm *WalletManager) GetWallets() ([]*openwallet.Wallet, error) {

	wallets, err := openwallet.GetWalletsByKeyDir(wm.Config.keyDir)
	if err != nil {
		return nil, err
	}

	for _, w := range wallets {
		w.DBFile = filepath.Join(wm.Config.dbPath, w.FileName()+".db")
	}

	return wallets, nil

}

//SendRawTransaction 广播交易
func (wm *WalletManager) SendRawTransaction(txJson map[string]interface{}) (string, error) {

	txid, err := wm.ApiClient.sendTransaction(txJson)

	if err != nil {
		return "", err
	}
	return txid, nil
}

// UpdateAddressNonce
func (wm *WalletManager) UpdateAddressNonce(wrapper openwallet.WalletDAI, address string, nonce uint64) {
	key := wm.Symbol() + "-nonce"
	wm.Log.Info(address, " set nonce ", nonce)
	err := wrapper.SetAddressExtParam(address, key, nonce)
	if err != nil {
		wm.Log.Errorf("WalletDAI SetAddressExtParam failed, err: %v", err)
	}
}

func (wm *WalletManager) ConvertPublicToAccountHash(pubKeyHex string) (string, error) {

	accountHashBytes, err := ConvertPublicToAccountHashBytes(pubKeyHex)
	if err != nil {
		return "", nil
	}
	return hex.EncodeToString(accountHashBytes), nil
}

func ConvertPublicToAccountHashBytes(pubKeyHex string) ([]byte, error) {
	if len(pubKeyHex) == 66 && strings.HasPrefix(pubKeyHex, "01") {
		pubKeyHex = pubKeyHex[2:]
	}
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, nil
	}
	split, _ := hex.DecodeString("00")
	prefix := append([]byte("ed25519"), split...)
	pubKeyBytesAll := append(prefix, pubKeyBytes...)
	pkHash := owcrypt.Hash(pubKeyBytesAll, 32, owcrypt.HASH_ALG_BLAKE2B)
	return pkHash, nil
}

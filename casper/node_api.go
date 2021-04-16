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
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

type ClientInterface interface {
	Call(path string, request []interface{}) (*gjson.Result, error)
}

// A Client is a Elastos RPC client. It performs RPCs over HTTP using JSON
// request and responses. A Client must be configured with a secret token
// to authenticate with other Cores on the network.
type Client struct {
	BaseURL     string
	AccessToken string
	Debug       bool
	client      *req.Req
}

// NewClient 创建 API 客户端
func NewClient(url string, debug bool) *Client {
	c := Client{
		BaseURL: url,
		Debug:   debug,
	}

	api := req.New()
	c.client = api
	return &c
}

// PostCall 发送 POST 请求
func (c *Client) PostCall(path string, v map[string]interface{}) (*gjson.Result, error) {
	if c.Debug {
		log.Debug("Start Request API...")
	}

	r, err := c.client.Post(c.BaseURL+path, req.BodyJSON(&v))

	if c.Debug {
		log.Std.Info("Request API Completed")
	}

	if c.Debug {
		log.Debugf("%+v\n", r)
	}

	if err != nil {
		return nil, err
	}

	result := gjson.ParseBytes(r.Bytes())

	err = isError(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCall 发送 GET 请求
func (c *Client) GetCall(path string) (*gjson.Result, error) {

	if c.Debug {
		log.Debug("Start Request API...")
	}

	r, err := c.client.Get(c.BaseURL + path)

	if c.Debug {
		log.Std.Info("Request API Completed")
	}

	if c.Debug {
		log.Debugf("%+v\n", r)
	}

	if err != nil {
		return nil, err
	}

	result := gjson.ParseBytes(r.Bytes())

	err = isError(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
func (c *Client) RpcCall(method string, params interface{}) (*gjson.Result, error) {
	authHeader := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	body := make(map[string]interface{}, 0)
	body["jsonrpc"] = "2.0"
	body["id"] = 1
	body["method"] = method
	body["params"] = params

	if c.Debug {
		log.Debugf("url : %+v", c.BaseURL)
	}

	r, err := req.Post(c.BaseURL, req.BodyJSON(&body), authHeader)

	if c.Debug {
		log.Debugf("%+v\n", r)
	}

	if err != nil {
		return nil, err
	}

	resp := gjson.ParseBytes(r.Bytes())
	err = isError(&resp)
	if err != nil {
		log.Info("scan near resp info", resp.String())
		return nil, err
	}

	result := resp.Get("result")

	return &result, nil
}

func (c *Client) RpcCall2(method string, params map[string]interface{}) (*gjson.Result, error) {
	authHeader := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	body := make(map[string]interface{}, 0)
	body["jsonrpc"] = "2.0"
	body["id"] = 1
	body["method"] = method
	body["params"] = params

	if c.Debug {
		log.Debugf("url : %+v", c.BaseURL)
	}

	r, err := req.Post(c.BaseURL, req.BodyJSON(&body), authHeader)

	if c.Debug {
		log.Debugf("%+v\n", r)
	}

	if err != nil {
		return nil, err
	}

	resp := gjson.ParseBytes(r.Bytes())
	err = isError(&resp)
	if err != nil {
		log.Info("scan near resp info", resp.String())
		return nil, err
	}

	result := resp.Get("result")

	return &result, nil
}

// isError 检查请求结果是否异常
func isError(result *gjson.Result) error {
	if result == nil {
		return fmt.Errorf("request failed result is nil")
	}

	if result.Get("message").Exists() {
		return fmt.Errorf("request failed resp message: %s", result.Get("message").String())
	}

	if result.Get("error").Exists() {
		return fmt.Errorf("request failed resp error: %s", result.Get("error").String())
	}

	return nil
}

// getLastBlockHeight 获取当前最高区块
func (c *Client) getLastBlockHeight() (uint64, error) {
	status, err := c.getLastStatus()
	if err != nil {
		return 0, err
	}
	return status.Height, nil
}

// getTxMaterial 获取离线签名所需的参数
func (c *Client) getTxMaterial() (*TxArtifacts, error) {
	resp, err := c.GetCall("/transaction/material")

	if err != nil {
		return nil, err
	}
	return GetTxArtifacts(resp), nil
}

// getLastBlock 获取当前最新状态
func (c *Client) getLastStatus() (*Status, error) {
	resp, err := c.RpcCall("info_get_status", nil)

	if err != nil {
		return nil, err
	}

	return NewStatus(resp)
}

func (c *Client) getBlockByHeight(blockHeight uint64) (*Block, error) {
	method := "chain_get_block"
	param := make(map[string]interface{}, 0)
	blockIdentifier := make(map[string]interface{}, 0)
	param["block_identifier"] = blockIdentifier
	blockIdentifier["Height"] = blockHeight
	resp, err := c.RpcCall(method, param)

	if err != nil {
		return nil, err
	}
	block := NewBlock(resp)
	txArray, err := c.getBlockTransferTxByHeight(blockHeight)
	if err != nil {
		return nil, err
	}
	if len(txArray) > 0 && txArray[0].BlockHash != block.Hash {
		return nil, errors.New("block hash mismatch with txData")
	}
	block.Transactions = txArray
	return block, nil
}

func (c *Client) getBlockTransferTxByHeight(blockHeight uint64) ([]*Transaction, error) {
	method := "chain_get_block_transfers"
	param := make(map[string]interface{}, 0)
	blockIdentifier := make(map[string]interface{}, 0)
	param["block_identifier"] = blockIdentifier
	blockIdentifier["Height"] = blockHeight
	resp, err := c.RpcCall(method, param)

	if err != nil {
		return nil, err
	}
	return GetTransactionInBlock(resp, blockHeight), nil
}

// getBalance 获取地址余额
func (c *Client) getBalance(address string, ignoreReserve bool, reserveAmount int64) (*AddrBalance, error) {
	r, err := c.GetCall(fmt.Sprintf("/accounts/%s/balance-info", address))

	if err != nil {
		return nil, err
	}

	if r.Get("error").String() == "actNotFound" {
		return &AddrBalance{Address: address, Balance: big.NewInt(0), Actived: false, Nonce: uint64(0)}, nil
	}

	free := big.NewInt(r.Get("free").Int())
	if ignoreReserve {
		if free.Cmp(big.NewInt(reserveAmount)) == 1 {
			free.Sub(free, big.NewInt(reserveAmount))
		} else {
			free = big.NewInt(0)
		}
	}
	feeFrozen := big.NewInt(r.Get("feeFrozen").Int())
	nonce := uint64(r.Get("nonce").Uint())
	balance := new(big.Int)
	balance = balance.Sub(free, feeFrozen)
	return &AddrBalance{Address: address, Balance: balance, Freeze: feeFrozen, Free: free, Actived: true, Nonce: nonce}, nil
}

// sendTransaction 发送签名交易
func (c *Client) sendTransaction(rawTx string) (string, error) {
	body := map[string]interface{}{
		"tx": rawTx,
	}

	resp, err := c.PostCall("/transaction", body)
	if err != nil {
		return "", err
	}

	time.Sleep(time.Duration(1) * time.Second)

	log.Debug("sendTransaction result : ", resp)

	if resp.Get("error").String() != "" && resp.Get("cause").String() != "" {
		return "", errors.New("Submit transaction with error: " + resp.Get("error").String() + "," + resp.Get("cause").String())
	}

	return resp.Get("hash").String(), nil
}

func RemoveOxToAddress(addr string) string {
	if strings.Index(addr, "0x") == 0 {
		return addr[2:]
	}
	return addr
}

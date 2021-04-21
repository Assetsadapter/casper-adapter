package casper

type ApiClient struct {
	Client    *Client
	APIChoose string
}

func NewApiClient(wm *WalletManager) error {
	api := ApiClient{}

	if len(wm.Config.APIChoose) == 0 {
		wm.Config.APIChoose = "rpc" //默认采用rpc连接
	}
	api.APIChoose = wm.Config.APIChoose
	if api.APIChoose == "rpc" {
		api.Client = NewClient(wm.Config.NodeAPI, false)
	} else if api.APIChoose == "ws" {
	}

	wm.ApiClient = &api

	return nil
}

// 获取当前最高区块
func (c *ApiClient) getLstBlockHeight() (uint64, error) {
	var (
		currentHeight uint64
		err           error
	)
	if c.APIChoose == "rpc" {
		currentHeight, err = c.Client.getLastBlockHeight()
	}

	return currentHeight, err
}

//获取当前最新高度
func (c *ApiClient) getLastBlock() (*Block, error) {
	var (
		mostHeightBlock *Block
		err             error
	)
	if c.APIChoose == "rpc" {
		//mostHeightBlock, err = c.Client.getLastBlock()
	} else if c.APIChoose == "ws" {
		//mostHeightBlock, err = decoder.wm.WSClient.getBlockHeight()
	}

	return mostHeightBlock, err
}

// 获取地址余额
func (c *ApiClient) getBalance(address, stateRootHash string) (*AddrBalance, error) {
	var (
		balance *AddrBalance
		err     error
	)

	if c.APIChoose == "rpc" {
		balance, err = c.Client.getBalance(address, stateRootHash)
	}

	return balance, err
}

func (c *ApiClient) getBlockByHeight(height uint64) (*Block, error) {
	var (
		block *Block
		err   error
	)
	if c.APIChoose == "rpc" {
		block, err = c.Client.getBlockByHeight(height)
	}

	return block, err
	return nil, nil
}

func (c *ApiClient) sendTransaction(txJson map[string]interface{}) (string, error) {
	var (
		txid string
		err  error
	)
	if c.APIChoose == "rpc" {
		txid, err = c.Client.sendTransaction(txJson)
	}

	return txid, err
}

func (c *ApiClient) getTxMaterial() (*TxArtifacts, error) {
	var (
		txMaterial *TxArtifacts
		err        error
	)
	if c.APIChoose == "rpc" {
		txMaterial, err = c.Client.getTxMaterial()
	}

	return txMaterial, err
}

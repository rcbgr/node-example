package model

type NodeAdvancedApiRequest struct {
	Config *NodeConfig     `json:"config"`
	Body   *JsonRpcRequest `json:"body"`
	Result interface{}     `json:"result"`
}

type JsonRpcRequest struct {
	Id             string      `json:"id"`
	Method         string      `json:"method"`
	JsonRpcVersion string      `json:"jsonrpc"`
	Params         interface{} `json:"params"`
}

type JsonRpcResponse struct {
	Id             string      `json:"id"`
	Method         string      `json:"method"`
	JsonRpcVersion string      `json:"jsonrpc"`
	Result         interface{} `json:"result"`
}

type NodeAdvancedApiResponse struct {
	Body   []byte
	Result interface{}
}

type NodeAdvancedApiGetBalances struct {
	AddressAndContracts []*NodeAdvancedApiAddressAndContract `json:"addressAndContractList"`
	Blockchain          string                               `json:"blockchain"`
	Network             string                               `json:"network"`
}

type NodeAdvancedApiAddressAndContract struct {
	Address  string `json:"address"`
	Contract string `json:"contract"`
}

type NodeAdvancedApiGetSingleBalanceRequest struct {
	Address    string `json:"address"`
	Contract   string `json:"contract"`
	Blockchain string `json:"blockchain"`
	Network    string `json:"network"`
}

type NodeAdvancedApiTokenBalance struct {
	Contract string `json:"contract"`
	Amount   string `json:"amount"`
	Decimals uint   `json:"decimals"`
}

type NodeAdvancedApiGetSingleBalanceResponse struct {
	NativeAmount string                       `json:"nativeAmount"`
	NativeUnit   string                       `json:"nativeUnit"`
	TokenBalance *NodeAdvancedApiTokenBalance `json:"tokenBalance"`
}

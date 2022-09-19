package model

type NodeAdvancedApiRequest struct {
	Config *NodeConfig  `json:"config"`
	Body   *JsonRpcBody `json:"body"`
}

type JsonRpcBody struct {
	Id             string      `json:"id"`
	Method         string      `json:"method"`
	JsonRpcVersion string      `json:"jsonrpc"`
	Params         interface{} `json:"params"`
}

type NodeAdvancedApiResponse struct {
	Body []byte
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

type NodeAdvancedApiGetSingleBalanceResponse struct {

	/*
	   	BlockHeight string `json:"blockHeight"`

	       "blockHeight": "0xed746b",
	       "address": "0x0a59649758aa4d66e25f08dd01271e891fe52199",
	       "nativeAmount": "0x0",
	       "nativeUnit": "Wei",
	       "tokenBalance": {
	         "contract": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	         "amount": "0xc1d2329e6a074",
	         "decimals": 6
	       }
	     }
	*/
}

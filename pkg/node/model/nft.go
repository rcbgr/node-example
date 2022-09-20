package model

// This is a partial list of of fields in the doc
// See: https://docs.cloud.coinbase.com/node/reference/gettokensbycollectionid
type NodeNftApiContractToken struct {
	CurrentOwner    string `json:"currentOwner"`
	ContractAddress string `json:"contractAddress"`
	TokenId         string `json:"tokenId"`
	Name            string `json:"name"`
}

type NodeNftApiContractTokenResponse struct {
	Token *NodeNftApiContractToken `json:"token"`
}

type NodeNftApiContractTokens struct {
	Tokens         []*NodeNftApiContractToken `json:"tokenList"`
	Timestamp      string                     `json:"timeStamp"` // This is a typo in the API
	NextPageCursor string                     `json:"nextPageCursor"`
}

type NodeNftApiRequest struct {
	Config     *NodeConfig
	HttpMethod string
	Url        string
	Body       []byte
}

type NodeNftApiResponse struct {
	Body       []byte
	Status     string
	StatusCode int
	Proto      string
}

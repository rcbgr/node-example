package model

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

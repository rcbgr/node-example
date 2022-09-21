package node

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rcbgr/node-example/pkg/node/config"
	"github.com/rcbgr/node-example/pkg/node/model"
)

func NodeNftApiTokenOwnerByContractAndToken(
	contractAddress string,
	token string,
) (string, error) {

	url := fmt.Sprintf(
		"%s/api/nft/v2/contracts/%s/tokens/%s?networkName=%s",
		config.NodeConfig.Endpoint,
		contractAddress,
		token,
		config.NodeConfig.NftNetworkName,
	)

	fmt.Println(url)

	res, err := NodeNftApiCall(
		&model.NodeNftApiRequest{
			Config:     config.NodeConfig,
			HttpMethod: "GET",
			Url:        url,
		},
	)

	if err != nil {
		return "", fmt.Errorf("NFT API call: %v", err)
	}

	val := &model.NodeNftApiContractTokenResponse{}

	if err := json.Unmarshal(res.Body, val); err != nil {
		return "", fmt.Errorf("NFT API unmarshal: %v", err)
	}

	return val.Token.CurrentOwner, nil
}

func NodeNftApiTokensByContract(contractAddress string) ([]string, error) {

	var tokens []string
	var cursor string
	var url string
	for {

		url = fmt.Sprintf(
			"%s/api/nft/v2/contracts/%s/tokens?networkName=%s&pageSize=25",
			config.NodeConfig.Endpoint,
			contractAddress,
			config.NodeConfig.NftNetworkName,
		)

		if len(cursor) > 0 {
			url = fmt.Sprintf("%s&cursor=%s", url, cursor)
		}

		res, err := NodeNftApiCall(
			&model.NodeNftApiRequest{
				Config:     config.NodeConfig,
				HttpMethod: "GET",
				Url:        url,
			},
		)

		if err != nil {
			return tokens, fmt.Errorf("NFT API call: %v", err)
		}

		vals := &model.NodeNftApiContractTokens{}

		if err := json.Unmarshal(res.Body, vals); err != nil {
			return tokens, fmt.Errorf("NFT API unmarshal: %v", err)
		}

		for _, token := range vals.Tokens {
			tokens = append(tokens, token.TokenId)
		}

		cursor = vals.NextPageCursor

		if len(cursor) == 0 {
			break
		}

	}

	return tokens, nil

}

func NodeNftApiCall(req *model.NodeNftApiRequest) (*model.NodeNftApiResponse, error) {

	method := "POST"
	if req.HttpMethod == "GET" {
		method = "GET"
	}

	r, err := http.NewRequest(method, req.Url, bytes.NewReader(req.Body))
	if err != nil {
		return nil, fmt.Errorf("Unable to create request: %v", err)
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	r.Header.Add(
		"Authorization",
		fmt.Sprintf(
			"Basic %s", base64.StdEncoding.EncodeToString([]byte(
				fmt.Sprintf("%s:%s", req.Config.Username, req.Config.Password),
			)),
		),
	)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Unable to to do request: %v", err)
	}

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to to read response body: %v", err)
	}

	return &model.NodeNftApiResponse{
		Body:       out,
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
	}, nil

}

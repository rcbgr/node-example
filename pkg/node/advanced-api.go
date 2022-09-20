package node

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rcbgr/node-example/pkg/node/config"
	"github.com/rcbgr/node-example/pkg/node/model"
)

const (
	weiDecimals = uint(18)
)

func NodeAdvancedApiBalanceByContract(contract, address string) (
	contractAssets,
	nativeAssets float64,
	err error,
) {

	req := &model.NodeAdvancedApiRequest{
		Config: config.NodeConfig,
		Body: &model.JsonRpcRequest{
			Id:             uuid.New().String(),
			JsonRpcVersion: "2.0",
			Method:         "coinbaseCloud_getSingleBalance",
			Params: &model.NodeAdvancedApiGetSingleBalanceRequest{
				Address:    address,
				Contract:   contract,
				Blockchain: config.NodeConfig.Blockchain,
				Network:    config.NodeConfig.Network,
			},
		},
		Result: &model.JsonRpcResponse{
			Result: &model.NodeAdvancedApiGetSingleBalanceResponse{
				TokenBalance: &model.NodeAdvancedApiTokenBalance{},
			},
		},
	}

	res, e := NodeAdvancedApiCall(req)
	if e != nil {
		err = fmt.Errorf("Failed to call coinbaseCloud_getSingleBalance: %v", err)
		return
	}

	nativeAmount := res.Result.(*model.JsonRpcResponse).Result.(*model.NodeAdvancedApiGetSingleBalanceResponse).NativeAmount
	nativeUnit := res.Result.(*model.JsonRpcResponse).Result.(*model.NodeAdvancedApiGetSingleBalanceResponse).NativeUnit

	if nativeUnit != "Wei" {
		err = fmt.Errorf("Only know how to decode ETH + Wei - received:: %s", nativeUnit)
		return
	}

	nbalance, success := new(big.Float).SetString(nativeAmount)
	if !success {
		err = fmt.Errorf("Unable to convert nativeAmount to BigFloat: %s", nativeAmount)
		return
	}

	z := new(big.Float).Quo(nbalance, big.NewFloat(math.Pow10(18)))
	if c, accuracy := z.Float64(); accuracy != big.Exact {
		nativeAssets = c
	} else {
		err = fmt.Errorf("Unable to convert nativeAmount to BigInt: %s", nativeAmount)
		return
	}

	amount := res.Result.(*model.JsonRpcResponse).Result.(*model.NodeAdvancedApiGetSingleBalanceResponse).TokenBalance.Amount
	decimals := res.Result.(*model.JsonRpcResponse).Result.(*model.NodeAdvancedApiGetSingleBalanceResponse).TokenBalance.Decimals

	v, e := strconv.ParseInt(amount, 0, 64)
	if e != nil {
		err = fmt.Errorf("Unable to decode amount: %v", err)
		return
	}
	contractAssets = float64(v) / math.Pow(10, float64(decimals))

	return
}

func NodeAdvancedApiCall(req *model.NodeAdvancedApiRequest) (*model.NodeAdvancedApiResponse, error) {

	body, err := json.Marshal(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to to serialize api request body: %v", err)
	}

	r, err := http.NewRequest("POST", req.Config.Endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("Unable to to create request: %v", err)
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
		Timeout: 3 * time.Second,
	}

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Failed call to URL: %v", err)
	}

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Unable to read response body: %v", err)
	}

	if err := json.Unmarshal(out, req.Result); err != nil {
		return nil, fmt.Errorf("Advanced API unmarshal: %v", err)
	}

	return &model.NodeAdvancedApiResponse{Body: out, Result: req.Result}, nil
}

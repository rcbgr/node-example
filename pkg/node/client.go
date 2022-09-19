package node

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rcbgr/node-example/pkg/node/model"
)

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

	return &model.NodeAdvancedApiResponse{Body: out}, nil
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
		Timeout: 3 * time.Second,
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

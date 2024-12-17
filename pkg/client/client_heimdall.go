package client

import (
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"net/http"
	"net/url"
	"time"
)

type HeimdallClient struct {
	client  *httpclient.Client
	baseURL string
}

func NewHeimdall(timeout, retries int, baseURL string) *HeimdallClient {
	to := time.Duration(timeout) * time.Second

	backoff := heimdall.NewConstantBackoff(500*time.Millisecond, 2*time.Second)
	retrier := heimdall.NewRetrier(backoff)

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(to),
		httpclient.WithRetryCount(retries),
		httpclient.WithRetrier(retrier),
	)

	return &HeimdallClient{client: client, baseURL: baseURL}
}

func (hc *HeimdallClient) Get(path string, queryParams ...map[string]string) (*http.Response, error) {
	baseURL, err := url.Parse(hc.baseURL + path)
	if err != nil {
		return nil, err
	}

	if len(queryParams) > 0 {
		params := url.Values{}
		for key, value := range queryParams[0] {
			params.Add(key, value)
		}
		baseURL.RawQuery = params.Encode()
	}

	response, err := hc.client.Get(baseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

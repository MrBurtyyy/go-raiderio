package raiderio

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiError struct {
	StatusCode int `json:"statusCode"`
	ErrorCode string `json:"error"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}

type Region int

const (
	httpGet      = "GET"
	apiBaseUrl   = "https://raider.io/api/v1"
)

const (
	_ Region = iota
	US
	EU
	KR
	TW
	CN
)

func (r Region) String() string {
	rr := []string {
		"",
		"us",
		"eu",
		"kr",
		"tw",
		"cn",
	}

	return rr[r]
}

// Client is a client that can communicate with the raider.io API
type Client struct {
	// Region is the region that the client should connect to.  Use
	Region     Region
	httpClient *http.Client
}

// NewClient creates a new Client with the supplied Region
func NewClient(region Region) *Client {
	return &Client{
		Region:     region,
		httpClient: &http.Client{},
	}
}

func (c *Client) doRequest(ctx context.Context, httpMethod string, path string, queryParams url.Values, data interface{}) (interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, httpMethod, path, nil)
	if err != nil {
		return nil, err
	}

	// Add the region to the query params
	queryParams.Add("region", c.Region.String())
	req.URL.RawQuery = queryParams.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, makeErrorResponse(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func makeErrorResponse(res *http.Response) *ApiError {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	var r ApiError
	err = json.Unmarshal(body, &r)
	return &r
}

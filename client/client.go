package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// note: BASE_URL should be ENV
var (
	BASE_URL     = "https://example.com"
	EXAMPLE_PATH = BASE_URL + "/example"
)

type ExampleClient interface {
	GetName(idToken string) (*ExampleResponse, error)
}

type exampleClientImp struct {
	Client *resty.Client
}

func NewExampleClient() ExampleClient {
	return &exampleClientImp{
		Client: func() *resty.Client {
			client := resty.New()
			return client
		}(),
	}
}

type ExampleResponse struct {
	Name string `json:"name"`
}

type ExampleError struct {
	Error string `json:"error"`
}

func (c *exampleClientImp) GetName(id string) (*ExampleResponse, error) {
	url := EXAMPLE_PATH

	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"id": id,
		}).
		SetResult(&ExampleResponse{}).
		SetError(&ExampleError{}).
		Post(url)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("Not success")
	}

	return resp.Result().(*ExampleResponse), nil
}

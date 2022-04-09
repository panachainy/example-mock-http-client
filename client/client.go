package client

import (
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
)

var (
	clientOnce sync.Once
	exampleC   *ExampleClientImp
)

// note: BASE_URL should be ENV
var (
	BASE_URL     = "https://example.com"
	EXAMPLE_PATH = BASE_URL + "/example"
)

type ExampleClient interface {
	GetName(idToken string) (*ExampleResponse, error)
}

type ExampleClientImp struct {
	Client *resty.Client
}

func NewExampleClient() *ExampleClientImp {
	clientOnce.Do(func() {
		client := resty.New()
		exampleC = &ExampleClientImp{Client: client}
	})

	return exampleC
}

type ExampleResponse struct {
	Name string `json:"name"`
}

type ExampleError struct {
	Error string `json:"error"`
}

func (c *ExampleClientImp) GetName(id string) (*ExampleResponse, error) {
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
		return nil, fmt.Errorf("Not success %v", resp.StatusCode())
	}

	return resp.Result().(*ExampleResponse), nil
}

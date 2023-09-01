package moysklad

import (
	"fmt"
	"net/http"
)

type MoyskladClient struct {
	client http.Client
	token  string
}

func NewMoyskladClient(token string) *MoyskladClient {
	var client http.Client

	return &MoyskladClient{
		client: client,
		token:  token,
	}
}

func (c *MoyskladClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	return c.client.Do(req)
}

package hw3

import (
	"context"
	"net/http"
	"net/url"
)

type RemoteMethodInvocationClient interface {
	Invoke(ctx context.Context, url *url.URL, rmiMethod string, data []any, response any) error
}

var _ RemoteMethodInvocationClient = (*RmiClient)(nil)

type RmiClient struct{}

func NewRmiClient(client *http.Client) *RmiClient

func (r *RmiClient) Invoke(
	ctx context.Context,
	address *url.URL,
	rmiMethod string,
	arguments []any,
	response any,
) error

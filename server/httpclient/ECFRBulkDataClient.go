package httpclient

import (
	"context"
	"net/http"
)

type ECFRBulkDataClient struct {
	APIRoot    string
	HttpClient *Client
}

func (s *ECFRBulkDataClient) Get(
	ctx context.Context,
	path string,
) (*http.Response, error) {
	return s.HttpClient.Get(ctx, s.APIRoot+path)
}

func (s *ECFRBulkDataClient) GetXML(
	ctx context.Context,
	url string,
) (*http.Response, error) {
	return s.HttpClient.Get(ctx, url)
}

package httpclient

import (
	"context"
	"fmt"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

func (s *Client) Get(
	ctx context.Context,
	url string,
) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request, %v, %w", url, err)
	}

	resp, err := s.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed, %v, %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request retuned non-200 response: %v, %v", resp.StatusCode, url)
	}

	return resp, nil
}

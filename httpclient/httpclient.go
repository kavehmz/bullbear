package httpclient

import (
	htclient "github.com/tamnd/httpclient"
)

// HTTPClient represent a client that can handle json output or other formats
type HTTPClient struct{}

// JSON can parse http outputs in json format
func (s *HTTPClient) JSON(url string, v interface{}) error {
	return htclient.JSON(url, v)
}

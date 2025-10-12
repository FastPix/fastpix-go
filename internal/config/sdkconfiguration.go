

package config

import (
	"context"
	"github.com/fastpix/fastpix-go/retry"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SDKConfiguration struct {
	Client      HTTPClient
	Security    func(context.Context) (interface{}, error)
	ServerURL   string
	ServerIndex int
	ServerList  []string
	UserAgent   string
	RetryConfig *retry.Config
	Timeout     *time.Duration
}

func (c *SDKConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return c.ServerList[c.ServerIndex], nil
}

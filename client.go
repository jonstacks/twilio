package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// API constants
const (
	APIBase = "https://api.twilio.com"
)

// Client is a Twilio client
type Client struct {
	apiBase    string
	accountSID string
	authToken  string

	httpClient *http.Client
}

// NewClient initializes and returns a new client.
func NewClient(baseURL, accountSID, authToken string) *Client {
	return &Client{
		apiBase:    baseURL,
		accountSID: accountSID,
		authToken:  authToken,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type form interface {
	toForm() *url.Values
}

func (c Client) send(method, path string, f form) (*http.Response, error) {
	body := strings.NewReader(f.toForm().Encode())
	req, err := http.NewRequest(method, c.reqPath(path), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.accountSID, c.authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return c.httpClient.Do(req)
}

func (c Client) reqPath(path string) string {
	return fmt.Sprintf("%s%s", c.apiBase, path)
}

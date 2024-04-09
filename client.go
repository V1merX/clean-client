package clean_http

import (
	"golang.org/x/time/rate"
	"net/http"
)

type Client struct {
	HttpClient  *http.Client
	Headers     map[string]string
	Body        any
	RateLimiter *rate.Limiter
	MaxRetries  int
}

func NewClient() *Client {
	return &Client{
		HttpClient: &http.Client{},
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Head(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Connect(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodConnect, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Trace(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodTrace, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Post(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Put(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Patch(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Delete(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Options(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.Headers != nil {
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}
	}

	if c.MaxRetries > 1 {
		for count := 0; count < c.MaxRetries; count++ {
			res, err := c.HttpClient.Do(req)
			if err == nil {
				continue
			}
			return res, err
		}
	}

	return c.HttpClient.Do(req)
}

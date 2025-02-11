package fhirclient

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type CacheKey[V any] struct{}

type Client struct {
	HttpClient httpClient

	cache         *lru.Cache[CacheKey[any], any]
	authenticator Authenticator
}

type Option func(c *Client) error

func NewClient(host string, options ...Option) (*Client, error) {
	c := Client{
		HttpClient: httpClient{
			Host: host,
		},
	}

	for _, opt := range options {
		if err := opt(&c); err != nil {
			return nil, err
		}
	}

	c.HttpClient.client.Transport = &httpTransport{
		authenticator: c.authenticator,
	}

	return &c, nil
}

func WithCache(size int) Option {
	return func(c *Client) (err error) {
		c.cache, err = lru.New[CacheKey[any], any](size)
		return err
	}
}

type Queryable interface {
	GetType()
}

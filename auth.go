package fhirclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

var (
	ErrFailedToRetrieveToken = errors.New("failed to retrieve oauth2 token")
	ErrCfgInvalid            = errors.New("invalid oauth2 config")
)

type Authenticator interface {
	Handler(r http.Request) error
}

type BearerAuthenticator struct {
	Token string
}

func (b *BearerAuthenticator) Handler(r http.Request) error {
	strippedToken := strings.TrimPrefix(b.Token, "Bearer ")

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", strippedToken))

	return nil
}

func WithAuthBearer(token string) Option {
	return func(c *Client) (err error) {
		c.authenticator = &BearerAuthenticator{Token: token}

		return nil
	}
}

type BasicAuthenticator struct {
	username string
	password string
}

func WithBasicAuthenticator(username, password string) Option {
	return func(c *Client) (err error) {
		c.authenticator = &BasicAuthenticator{username: username, password: password}

		return nil
	}
}

func (b *BasicAuthenticator) Handler(r http.Request) error {
	r.SetBasicAuth(b.username, b.password)

	return nil
}

type OAuthAuthenticator struct {
	tokenSource oauth2.TokenSource
	cfg         oauth2.Config
	mu          sync.Mutex
}

func WithAuthOAuth2(cfg oauth2.Config, initialToken *oauth2.Token) Option {
	return func(c *Client) (err error) {
		tokenSource := cfg.TokenSource(context.Background(), initialToken)

		c.authenticator = &OAuthAuthenticator{
			cfg:         cfg,
			tokenSource: tokenSource,
		}

		if _, err := tokenSource.Token(); err != nil {
			return ErrCfgInvalid
		}

		return nil
	}
}

func (o *OAuthAuthenticator) Handler(r http.Request) error {
	o.mu.Lock()
	token, err := o.tokenSource.Token()
	if err != nil {
		return ErrFailedToRetrieveToken
	}
	o.mu.Unlock()

	r.Header.Set("Authorization", "Bearer "+token.AccessToken)

	return nil
}

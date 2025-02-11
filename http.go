package fhirclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	httpHeaderContentType = "Content-Type"
)

type httpClient struct {
	client http.Client
	Host   string
}

type httpTransport struct {
	authenticator Authenticator
}

func (t *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.authenticator != nil {
		if err := t.authenticator.Handler(*req); err != nil {
			return nil, err
		}
	}

	log.Debug().Msgf("[%s] %s", req.Method, req.URL.String())

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	log.Debug().Msgf("[%d] %s", resp.StatusCode, string(bodyBytes))

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return resp, nil
}

func RemoveURLProtocol(url string) string {
	parts := strings.Split(url, "://")[1:]
	return strings.Join(parts, "://")
}

func (h *httpClient) GetWithContext(ctx context.Context, url url.URL, target any) (resp *http.Response, err error) {
	path := fmt.Sprintf("%s%s", h.Host, url.String())

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return resp, err
	}
	req.Header.Set(httpHeaderContentType, "application/json")

	resp, err = h.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return resp, err
	}

	res, err := io.ReadAll(resp.Body)
	fmt.Print(string(res))

	return resp, nil
}

func (h *httpClient) Get(url url.URL, target *any) (*http.Response, error) {
	return h.GetWithContext(context.Background(), url, target)
}

func (h *httpClient) PostWithContext(ctx context.Context, url url.URL, payload any, target any) (resp *http.Response, err error) {
	path := fmt.Sprintf("%s%s", h.Host, url.String())

	log.Debug().Msgf("POST %s", path)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set(httpHeaderContentType, "application/json")

	resp, err = h.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if err := IsUnexpectedError(resp.StatusCode); err != nil {
		return resp, NewError(resp.Body)
	}

	if target != nil {
		if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func (h *httpClient) Post(url url.URL, payload any, target any) (*http.Response, error) {
	return h.PostWithContext(context.Background(), url, payload, target)
}

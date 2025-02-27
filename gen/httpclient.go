package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// ThrottledTransport Rate Limited HTTP Client
type ThrottledTransport struct {
	roundTripperWrap http.RoundTripper
	ratelimiter      *rate.Limiter
}

func (c *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	const maxRetries = 5 // Example retry limit

	for retryCount := 0; retryCount <= maxRetries; retryCount++ {
		err := c.ratelimiter.Wait(r.Context()) // Wait for rate limit on each retry
		if err != nil {
			return nil, err // Return error if rate limiter fails
		}

		resp, err := c.roundTripperWrap.RoundTrip(r) // Perform the actual request
		if err != nil {
			return resp, err // Return error from RoundTrip
		}

		if resp.StatusCode != http.StatusAccepted { // StatusAccepted is 202
			return resp, nil // Return successful response (not 202)
		}

		if retryCount < maxRetries { // Only pause and retry if we haven't reached max retries
			time.Sleep(2 * time.Second) // Pause for 2 seconds before retrying
			continue                    // Retry the request from the beginning of the loop
		} else {
			// Exceeded retry limit, return the 202 response or an error
			// You can decide what to do here when max retries are reached after getting 202s.
			return nil, fmt.Errorf("exceeded maximum retry attempts after receiving 202 status")
			// or, to return the last 202 response:
			// return resp, nil
		}
	}
	return nil, fmt.Errorf("unexpectedly reached end of retry loop without returning") // Should not happen
}

// NewThrottledTransport wraps transportWrap with a rate limitter
// example usage:
// client := http.DefaultClient
// client.Transport = NewThrottledTransport(10*time.Seconds, 60, http.DefaultTransport) allows 60 requests every 10 seconds
func NewThrottledTransport(limitPeriod time.Duration, requestCount int, transportWrap http.RoundTripper) http.RoundTripper {
	return &ThrottledTransport{
		roundTripperWrap: transportWrap,
		ratelimiter:      rate.NewLimiter(rate.Every(limitPeriod), requestCount),
	}
}

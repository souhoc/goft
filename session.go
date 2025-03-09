// Package goft provides Go bindings for the 42 school API.
//
// This package facilitates interaction with the 42 school API by providing
// a session-based approach to manage HTTP requests, including rate limiting
// and automatic handling of paginated responses.

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Copyright (c) 2025 Souria-Saky HOCQUENGHEM
package goft

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

// Session represents a session for interacting with the 42 school API.
// It manages rate limiting and provides methods to perform HTTP requests.
type Session struct {
	client              *http.Client
	secondlyRateLimiter *rate.Limiter
	hourlyRateLimiter   *rate.Limiter
}

// New creates a new Session with the provided HTTP client.
// It initializes rate limiters to control the frequency of API requests.
func New(client *http.Client) *Session {
	return &Session{
		client:              client,
		secondlyRateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
		hourlyRateLimiter:   rate.NewLimiter(rate.Every(time.Hour), 1200),
	}
}

// waitRateLimits waits for the rate limiters to allow a request to proceed.
// It ensures that the API rate limits are respected.
func (s *Session) waitRateLimits(ctx context.Context) error {
	if err := s.secondlyRateLimiter.Wait(ctx); err != nil {
		return err
	}
	if err := s.hourlyRateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

// Do performs an HTTP request using the session's client.
// It waits for the rate limiters and handles the Authorization header.
func (s *Session) Do(req *http.Request) (*http.Response, error) {
	if err := s.waitRateLimits(req.Context()); err != nil {
		return nil, err
	}
	return s.client.Do(req)
}

// Get performs a GET request to the specified URI and decodes the JSON response into v.
// It prepends the API base URL to the URI and waits for the rate limiters.
//
// Example usage:
//
//	var events []model.Event
//	Session.Get("/events", &events)
func (s *Session) Get(uri string, v any) (*http.Response, error) {
	if err := s.waitRateLimits(context.Background()); err != nil {
		return nil, err
	}

	resp, err := s.client.Get(BaseUrlVersioned + uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get %s: %s", uri, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, err
	}

	return resp, nil
}

// fetchOptions holds options for the Fetch method.
type fetchOptions struct {
	total int
}

// FetchOptionFunc is a function that configures fetchOptions.
type FetchOptionFunc func(o *fetchOptions)

// FetchWithTotal sets the total number of items to fetch.
func FetchWithTotal(total int) FetchOptionFunc {
	return func(o *fetchOptions) {
		o.total = total
	}
}

// Fetch performs a GET request to the specified URI and decodes the JSON response into v.
// It handles paginated responses and accumulates the results.
//
// Example usage:
//
//	var events []model.Event
//	Session.Fetch("/events", &events, FetchWithTotal(100))
func (s *Session) Fetch(uri string, v any, opts ...FetchOptionFunc) (*http.Response, error) {
	o := fetchOptions{
		total: -1,
	}
	for _, of := range opts {
		of(&o)
	}
	var lastResponse *http.Response

	rValue := reflect.ValueOf(v)
	if rValue.Kind() != reflect.Ptr || rValue.Elem().Kind() != reflect.Slice {
		return nil, fmt.Errorf("value must be a pointer to a slice")
	}

	accumulated := reflect.MakeSlice(rValue.Elem().Type(), 0, 0)
	for {
		tmp := reflect.New(rValue.Elem().Type()).Interface()
		resp, err := s.Get(uri, &tmp)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return resp, nil
		}
		lastResponse = resp

		link := ParseLink(resp.Header.Get("link"))
		if link.Next != "" {
			uri = link.Next
		}

		accumulated = reflect.AppendSlice(accumulated, reflect.ValueOf(tmp).Elem())

		xTotal, err := strconv.Atoi(resp.Header.Get("x-total"))
		if err != nil {
			return nil, err
		}
		if o.total < 0 {
			o.total = xTotal
		}
		if o.total <= accumulated.Len() || link.Next == "" {
			break
		}
	}

	rValue.Elem().Set(accumulated)

	return lastResponse, nil
}

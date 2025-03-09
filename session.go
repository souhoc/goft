// goft - go bindings for 42 school API
//
// Copyright 2025 by Souria-Saky Hocquenghem - All Rights Reserved
// You may use, distribute and modify this code under the terms of the MIT license
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

type Session struct {
	client              *http.Client
	secondlyRateLimiter *rate.Limiter
	hourlyRateLimiter   *rate.Limiter
}

func New(client *http.Client) *Session {
	return &Session{
		client:              client,
		secondlyRateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
		hourlyRateLimiter:   rate.NewLimiter(rate.Every(time.Hour), 1200),
	}

}

func (s *Session) waitRateLimits(ctx context.Context) error {
	if err := s.secondlyRateLimiter.Wait(ctx); err != nil {
		return err
	}
	if err := s.hourlyRateLimiter.Wait(ctx); err != nil {
		return err
	}

	return nil
}

// Do is only doing an http client.Do. It wait the rate limit and handle the Authorization
func (s *Session) Do(req *http.Request) (*http.Response, error) {
	if err := s.waitRateLimits(req.Context()); err != nil {
		return nil, err
	}
	return s.client.Do(req)
}

// Get do the same as Session.Do(), but decode the json body to v and prepend the api base url
//
//	e.g:
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

type fetchOptions struct {
	total int
}

type FetchOptionFunc func(o *fetchOptions)

func FetchWithTotal(total int) FetchOptionFunc {
	return func(o *fetchOptions) {
		o.total = total
	}
}

// Fetch do as Get. It return the last http.Response
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

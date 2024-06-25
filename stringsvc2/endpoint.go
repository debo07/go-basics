package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
		v, err := svc.UpperCase(req.S)
		if err != nil {
			return uppercaseResponse{ V: v, Err: err.Error()}, err
		}
		return uppercaseResponse{V: v, Err: ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{V: v}, nil
	}
}
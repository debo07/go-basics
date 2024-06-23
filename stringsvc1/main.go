package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Layer: Service
// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// stringService is a concrete implementation of StringService
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	fmt.Println("Step 4: In Service(BL) layer - UpperCase")
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	fmt.Println("Step 4: In Service(BL) layer - Count")
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// Layer: Transport
// For each method, we define request and response structs
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("Step 2: Decode Request - Uppercase")
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("Step 2: Decode Request - Count")
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("Step 5: Encode Response - Uppercase/Count")
	return json.NewEncoder(w).Encode(response)
}

// Layer: Endpoint
// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	fmt.Println("Step 1: Making endpoints - UpperCase")
	return func(_ context.Context, request interface{}) (interface{}, error) {
		fmt.Println("Step 3: Executing Endpoint - Uppercase")
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	fmt.Println("Step 1: Making endpoints - Count")
	return func(_ context.Context, request interface{}) (interface{}, error) {
		fmt.Println("Step 3: Executing Endpoint - Count")
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

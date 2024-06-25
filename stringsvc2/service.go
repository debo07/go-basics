package main

import (
	"errors"
	"strings"
)

type StringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}

type stringService struct {}

func (stringService) UpperCase(s string) (string, error) {
	if len(s) == 0 {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) (int) {
	return len(s)
}

var ErrEmpty = errors.New("empty string")
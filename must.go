package fhirclient

import (
	"fmt"
	"os"
	"testing"
)

func Must[T any](obj T, err error) T {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return obj
}

func MustTesting[T any](t *testing.T, obj T, err error) T {
	if err != nil {
		t.Error(err)
	}
	return obj
}

func Ptr[T any](v T) *T {
	return &v
}

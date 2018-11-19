package fakes_test

import (
	"testing"

	"github.com/ctco-dev/go-testing-frameworks/fakes"
)

type fakeReader struct{}

func (f *fakeReader) Read() string {
	return "foobar"
}

func TestService_Serve(t *testing.T) {
	service := fakes.NewService(&fakeReader{})
	result := service.Serve()
	if result != "foobar_" {
		t.Error("Expected foobar, but received ", result)
	}
}

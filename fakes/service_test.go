package fakes_test

import (
	"ctco-dev/go-testing-frameworks/fakes"
	"testing"
)

type fakeReader struct {}

func (f *fakeReader) Read() string {
	return "foobar"
}

func TestService_Serve(t *testing.T) {
	service := fakes.NewService(&fakeReader{})
	result := service.Serve()
	if result != "foobar" {
		t.Error("Expected foobar, but received ", result)
	}
}
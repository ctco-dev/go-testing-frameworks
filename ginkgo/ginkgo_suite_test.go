package ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// $ ginkgo bootstrap
func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Suite")
}

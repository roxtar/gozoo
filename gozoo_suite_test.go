package gozoo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGozoo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gozoo Suite")
}

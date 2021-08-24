package api

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoadmap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saver Suite")
}

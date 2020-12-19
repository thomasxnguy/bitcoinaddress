package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

// TestCommon is the main function to run the test suite for "common" package using gingko testing framework
func TestCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

package redaction_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedaction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redaction Suite")
}

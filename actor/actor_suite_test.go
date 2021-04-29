package actor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestActor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Actor Suite")
}

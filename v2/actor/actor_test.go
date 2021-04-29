package actor_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/blgm/foo/v2/actor"
)

var _ = Describe("Actor", func() {
	It("does not panic", func() {
		actor.Do()
	})
})

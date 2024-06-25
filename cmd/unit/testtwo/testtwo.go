package testtwo

import (
	"github.com/shantanubansal/gest/cmd/wrapper"
	"testing"
)

func init() {
	wrapper.AddTest("Two", Two)
}
func Two(t *testing.T) {
	t.Fail()
}

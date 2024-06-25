package testone

import (
	"github.com/shantanubansal/gest/cmd/wrapper"
	"testing"
)

func init() {
	wrapper.AddTest("One", One)
}
func One(t *testing.T) {
	t.Fail()
}

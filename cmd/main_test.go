package main_test

import (
	_ "github.com/shantanubansal/gest/cmd/unit/testone"
	_ "github.com/shantanubansal/gest/cmd/unit/testtwo"
	"github.com/shantanubansal/gest/cmd/wrapper"
	"testing"
)

func TestWrapperTests(t *testing.T) {
	tw := wrapper.GetTestWrapper()

	tw.RunAll(t)
}

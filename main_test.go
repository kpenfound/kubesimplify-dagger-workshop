package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGreeting(t *testing.T) {
	g := greeting()
	should := "{\"greeting\":\"Hello Daggernauts!\"}"

	assert.Equal(t, should, g)
}

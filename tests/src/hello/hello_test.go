package tests

import (
	"github.com/stretchr/testify/assert"
	"lib/helper"
	"testing"
)

func TestHello(t *testing.T) {
	val := helper.GetEnv("check", "ok")
	assert.Equal(t, "ok", val, "val is not equal to ok")
}

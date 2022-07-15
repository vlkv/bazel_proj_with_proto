package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	main "github.com/vlkv/bazel_proj_with_proto"
)

func TestOne(t *testing.T) {
	assert.Equal(t, "foo", main.GetFoo())
}

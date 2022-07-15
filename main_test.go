package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	main "gitlab.com/myorg/myproj"
)

func TestOne(t *testing.T) {
	assert.Equal(t, "foo", main.GetFoo())
}

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCookie(t *testing.T) {
	cookie := ParseCookie("a=1; b=2")
	assert.Equal(t, "1", cookie["a"])
	assert.Equal(t, "2", cookie["b"])
}

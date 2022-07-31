package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustBeGreaterThanZeroFloat(t *testing.T) {
	assert.NotNil(t, MustBeGreaterThanZeroFloat("foo", -1.0)) // coverage only
}

func TestMustBeGreaterThanZeroDecimal(t *testing.T) {
	assert.NotNil(t, MustBeGreaterThanZeroDecimal("foo", -1)) // coverage only
}

func TestMustBeGreaterThanOrEqualToZeroFloat(t *testing.T) {
	assert.NotNil(t, MustBeGreaterThanOrEqualToZeroFloat("foo", -1.0)) // coverage only
}

func TestMustBeGreaterThanOrEqualToZeroDecimal(t *testing.T) {
	assert.NotNil(t, MustBeGreaterThanOrEqualToZeroDecimal("foo", -1)) // coverage only
}

func TestMustBeGreaterThanFloat(t *testing.T) {
	assert.NotNil(t, MustBeGreaterThanFloat("foo", 5.0, "bar", 10.0)) // coverage only
}

func TestMustBeGreaterThanDecimal(t *testing.T) {
	assert.NotNil(t, MustBeGreaterThanDecimal("foo", 5, "bar", 10)) // coverage only
}

func TestMustNotBeNil(t *testing.T) {
	assert.NotNil(t, MustNotBeNil("foo")) // coverage only
}

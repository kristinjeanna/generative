package random

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	iterations int = 100000
)

func TestNew(t *testing.T) {
	_, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
}

func TestNew_Error(t *testing.T) {
	_, err := New(nil)
	assert.Error(t, err)
}

func TestInt63(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	r.Int63() // coverage only
}

func TestSeed(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	r.Seed(time.Now().UnixMilli()) // coverage only
}

func TestFloat64(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got := r.Float64()
		assert.True(t, got >= 0 && got <= 1,
			"expected a number between 0 and 1, but got '%f'", got)
	}
}

func TestFloat64InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got, err := r.Float64InRange(50, 200)
		assert.NoError(t, err)
		assert.True(t, got >= 50 && got <= 200,
			"expected a number between 50 and 200, but got '%f'", got)
	}
}

func TestFloat64InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Float64InRange(500, 20)
	assert.Error(t, err)
}

func TestFloat32(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got := r.Float32()
		assert.True(t, got >= 0 && got <= 1,
			"expected a number between 0 and 1, but got '%f'", got)
	}
}

func TestFloat32InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got, err := r.Float32InRange(50, 200)
		assert.NoError(t, err)
		assert.True(t, got >= 50 && got <= 200,
			"expected a number between 50 and 200, but got '%f'", got)
	}
}

func TestFloat32InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Float32InRange(500, 20)
	assert.Error(t, err)
}

func TestUint64(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		r.Uint64() // coverage only
	}
}

func TestUint64InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got, err := r.Uint64InRange(50, 200)
		assert.NoError(t, err)
		assert.True(t, got >= 50 && got <= 200,
			"expected a number between 50 and 200, but got '%d'", got)
	}
}

func TestUint64InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Uint64InRange(500, 20)
	assert.Error(t, err)
}

func TestUint32(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		r.Uint32() // coverage only
	}
}

func TestUint32InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got, err := r.Uint32InRange(50, 200)
		assert.NoError(t, err)
		assert.True(t, got >= 50 && got <= 200,
			"expected a number between 50 and 200, but got '%d'", got)
	}
}

func TestUint32InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Uint32InRange(500, 20)
	assert.Error(t, err)
}

func TestUint16(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		r.Uint16() // coverage only
	}
}

func TestUint16InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got, err := r.Uint16InRange(50, 200)
		assert.NoError(t, err)
		assert.True(t, got >= 50 && got <= 200,
			"expected a number between 50 and 200, but got '%d'", got)
	}
}

func TestUint16InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Uint16InRange(500, 20)
	assert.Error(t, err)
}

func TestUint8(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		r.Uint8() // coverage only
	}
}

func TestUint8InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	for i := 0; i < iterations; i++ {
		got, err := r.Uint8InRange(50, 200)
		assert.NoError(t, err)
		assert.True(t, got >= 50 && got <= 200,
			"expected a number between 50 and 200, but got '%d'", got)
	}
}

func TestUint8InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Uint8InRange(25, 20)
	assert.Error(t, err)
}

func TestBool(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var tResult, fResult int

	for i := 0; i < iterations; i++ {
		got := r.Bool()

		switch got {
		case true:
			tResult++
		default:
			fResult++
		}
	}

	// we should expect a fairly equal distribution of true and false values;
	// a deviation of 5% or more from a roughly equal distribution may cause
	// an occasional failure of this test.
	assert.True(t, float64(tResult) <= float64(iterations)*0.55,
		"got an unexpectedly large number of true values; count=%d", tResult)
}

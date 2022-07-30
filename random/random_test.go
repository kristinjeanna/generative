package random

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	iterations  int = 100000
	minExpected int = 45000
	maxExpected int = 55000
)

var (
	msgTrueValues     string = fmt.Sprintf("the number of true values was either less than %d or greater than %d; count=", minExpected, maxExpected) + "%d"
	msgPositiveValues string = fmt.Sprintf("the number of positive values was either less than %d or greater than %d; count=", minExpected, maxExpected) + "%d"
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
	var tCount int

	for i := 0; i < iterations; i++ {
		got := r.Bool()

		if got {
			tCount++
		}
	}

	result := tCount >= minExpected && tCount <= maxExpected
	assert.True(t, result, msgTrueValues, tCount)
}

func TestInt8(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got := r.Int8()

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt8InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got, err := r.Int8InRange(-100, 100)
		assert.NoError(t, err)

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt8InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Int8InRange(50, 20)
	assert.Error(t, err)
}

func TestInt16(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got := r.Int16()

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt16InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got, err := r.Int16InRange(-15000, 15000)
		assert.NoError(t, err)

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt16InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Int16InRange(500, 20)
	assert.Error(t, err)
}

func TestInt32(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got := r.Int32()

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt32InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got, err := r.Int32InRange(-21474836, 21474836)
		assert.NoError(t, err)

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt32InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Int32InRange(500, 20)
	assert.Error(t, err)
}

func TestInt64(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got := r.Int64()

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt64InRange(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)
	var posCount int

	for i := 0; i < iterations; i++ {
		got, err := r.Int64InRange(-922337203685477, 922337203685477)
		assert.NoError(t, err)

		if got >= 0 {
			posCount++
		}
	}

	result := posCount >= minExpected && posCount <= maxExpected
	assert.True(t, result, msgPositiveValues, posCount)
}

func TestInt64InRange_Error(t *testing.T) {
	r, err := New(NewXoshiro512StarStar())
	assert.NoError(t, err)

	_, err = r.Int64InRange(500, 20)
	assert.Error(t, err)
}

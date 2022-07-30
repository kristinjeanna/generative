package random

import (
	"errors"
	"math"
	mrnd "math/rand"

	generrs "github.com/kristinjeanna/generative/errors"
)

var (
	ErrPRNGNotInitialized = errors.New("PRNG not initialized")
)

// Rand is a math.Source64 that wraps an underlying math.Source64
// and implements additional random number methods.
type Rand struct {
	prng mrnd.Source64
}

func New(src mrnd.Source64) (*Rand, error) {
	if src == nil {
		return nil, generrs.MustNotBeNil("src")
	}
	return &Rand{prng: src}, nil
}

func (r Rand) Int63() int64 {
	return r.prng.Int63() // coverage only
}

func (r Rand) Seed(seed int64) {
	r.prng.Seed(seed) // coverage only
}

// Float64 returns a random float64 from 0 to 1.
func (r Rand) Float64() float64 {
	return float64(r.prng.Uint64()) / float64(math.MaxUint64)
}

// Float64InRange returns a random float32 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Float64InRange(min, max float64) (float64, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanFloat("max", max, "min", min)
	}

	diff := max - min
	return r.Float64()*diff + min, nil
}

// Float32 returns a random float32 from 0 to 1.
func (r Rand) Float32() float32 {
	return float32(r.Float64())
}

// Float32InRange returns a random float32 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Float32InRange(min, max float32) (float32, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanFloat("max", max, "min", min)
	}

	diff := max - min
	return r.Float32()*diff + min, nil
}

// Uint64 returns a random uint64.
//
// An error is returned if the PRNG has not been initialized.
func (r Rand) Uint64() uint64 {
	return r.prng.Uint64()
}

// Uint64InRange returns a random uint64 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Uint64InRange(min, max uint64) (uint64, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return uint64(r.Float64()*float64(diff)) + min, nil
}

// Uint32 returns a random uint32.
func (r Rand) Uint32() uint32 {
	return uint32(r.Uint64())
}

// Uint32InRange returns a random uint32 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Uint32InRange(min, max uint32) (uint32, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return uint32(r.Float64()*float64(diff)) + min, nil
}

// Uint16 returns a random uint16.
func (r Rand) Uint16() uint16 {
	return uint16(r.Uint64())
}

// Uint16InRange returns a random uint16 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Uint16InRange(min, max uint16) (uint16, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return uint16(r.Float64()*float64(diff)) + min, nil
}

// Uint8 returns a random uint8.
func (r Rand) Uint8() uint8 {
	return uint8(r.Uint64())
}

// Uint8InRange returns a random uint8 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Uint8InRange(min, max uint8) (uint8, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return uint8(r.Float64()*float64(diff)) + min, nil
}

// Bool returns a random boolean.
func (r Rand) Bool() bool {
	return r.Float64() >= 0.5
}

// Int8 returns a random int8.
func (r Rand) Int8() int8 {
	return int8(r.Uint8())
}

// Int8InRange returns a random int8 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Int8InRange(min, max int8) (int8, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return int8(r.Float64()*float64(diff)) + min, nil
}

// Int16 returns a random int16.
func (r Rand) Int16() int16 {
	return int16(r.Uint16())
}

// Int16InRange returns a random int16 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Int16InRange(min, max int16) (int16, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return int16(r.Float64()*float64(diff)) + min, nil
}

// Int32 returns a random int32.
func (r Rand) Int32() int32 {
	return int32(r.Uint32())
}

// Int32InRange returns a random int32 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Int32InRange(min, max int32) (int32, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return int32(r.Float64()*float64(diff)) + min, nil
}

// Int64 returns a random int64.
func (r Rand) Int64() int64 {
	return int64(r.Uint64())
}

// Int64InRange returns a random int64 from min to max.
//
// An error is returned if max is less than or equal to min.
func (r Rand) Int64InRange(min, max int64) (int64, error) {
	if max <= min {
		return 0, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	return int64(r.Float64()*float64(diff)) + min, nil
}

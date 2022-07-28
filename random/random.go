package random

import (
	"errors"
	"math"
	mrnd "math/rand"

	generrs "github.com/kristinjeanna/generative/errors"
)

var (
	prng mrnd.Source64

	ErrPRNGNotInitialized = errors.New("PRNG not initialized")
)

func Initialize(src mrnd.Source64) {
	prng = src
}

// Float64 returns a random float64 from 0 to 1.
//
// An error is returned if the PRNG has not been initialized.
func Float64() (*float64, error) {
	if prng == nil {
		return nil, ErrPRNGNotInitialized
	}

	result := float64(prng.Uint64()) / float64(math.MaxUint64)
	return &result, nil
}

// Float64InRange returns a random float32 from min to max.
//
// An error is returned if either (1) the PRNG has not been initialized, or
// (2) min is greater than max.
func Float64InRange(min, max float64) (*float64, error) {
	fraction, err := Float64()
	if err != nil {
		return nil, err
	}

	if max <= min {
		return nil, generrs.MustBeGreaterThanFloat("max", max, "min", min)
	}

	diff := max - min
	result := *fraction*diff + min
	return &result, nil
}

// Float32 returns a random float32 from 0 to 1.
//
// An error is returned if the PRNG has not been initialized.
func Float32() (*float32, error) {
	f, err := Float64()
	if err != nil {
		return nil, err
	}

	result := float32(*f)
	return &result, nil
}

// Float32InRange returns a random float32 from min to max.
//
// An error is returned if either (1) the PRNG has not been initialized, or
// (2) min is greater than max.
func Float32InRange(min, max float32) (*float32, error) {
	fraction, err := Float32()
	if err != nil {
		return nil, err
	}

	if max <= min {
		return nil, generrs.MustBeGreaterThanFloat("max", max, "min", min)
	}

	diff := max - min
	result := *fraction*diff + min
	return &result, nil
}

// Uint64 returns a random uint64.
//
// An error is returned if the PRNG has not been initialized.
func Uint64() (*uint64, error) {
	if prng == nil {
		return nil, ErrPRNGNotInitialized
	}

	result := prng.Uint64()
	return &result, nil
}

// Uint64InRange returns a random uint64 from min to max.
//
// An error is returned if either (1) the PRNG has not been initialized, or
// (2) min is greater than max.
func Uint64InRange(min, max uint64) (*uint64, error) {
	fraction, err := Float64()
	if err != nil {
		return nil, err
	}

	if max <= min {
		return nil, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	result := uint64(*fraction*float64(diff)) + min
	return &result, nil
}

// Uint32 returns a random uint32.
//
// An error is returned if the PRNG has not been initialized.
func Uint32() (*uint32, error) {
	if prng == nil {
		return nil, ErrPRNGNotInitialized
	}

	result := uint32(prng.Uint64())
	return &result, nil
}

// Uint32InRange returns a random uint32 from min to max.
//
// An error is returned if either (1) the PRNG has not been initialized, or
// (2) min is greater than max.
func Uint32InRange(min, max uint32) (*uint32, error) {
	fraction, err := Float64()
	if err != nil {
		return nil, err
	}

	if max <= min {
		return nil, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	result := uint32(*fraction*float64(diff)) + min
	return &result, nil
}

// Uint16 returns a random uint16.
//
// An error is returned if the PRNG has not been initialized.
func Uint16() (*uint16, error) {
	if prng == nil {
		return nil, ErrPRNGNotInitialized
	}

	result := uint16(prng.Uint64())
	return &result, nil
}

// Uint16InRange returns a random uint16 from min to max.
//
// An error is returned if either (1) the PRNG has not been initialized, or
// (2) min is greater than max.
func Uint16InRange(min, max uint16) (*uint16, error) {
	fraction, err := Float64()
	if err != nil {
		return nil, err
	}

	if max <= min {
		return nil, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	result := uint16(*fraction*float64(diff)) + min
	return &result, nil
}

// Uint8 returns a random uint8.
//
// An error is returned if the PRNG has not been initialized.
func Uint8() (*uint8, error) {
	if prng == nil {
		return nil, ErrPRNGNotInitialized
	}

	result := uint8(prng.Uint64())
	return &result, nil
}

// Uint8InRange returns a random uint8 from min to max.
//
// An error is returned if either (1) the PRNG has not been initialized, or
// (2) min is greater than max.
func Uint8InRange(min, max uint8) (*uint8, error) {
	fraction, err := Float64()
	if err != nil {
		return nil, err
	}

	if max <= min {
		return nil, generrs.MustBeGreaterThanDecimal("max", max, "min", min)
	}

	diff := max - min
	result := uint8(*fraction*float64(diff)) + min
	return &result, nil
}

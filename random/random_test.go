package random

import (
	"testing"
)

const (
	iterations int = 100000
)

func TestFloat64(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Float64()
		if err != nil {
			t.Error(err)
		}

		if *got < 0 || *got > 1 {
			t.Errorf("expected a number between 0 and 1, but got '%f'", *got)
		}
	}
}

func TestFloat64InRange(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Float64InRange(50, 200)
		if err != nil {
			t.Error(err)
		}

		if *got < 50 || *got > 200 {
			t.Errorf("expected a number between 50 and 200, but got '%f'", *got)
		}
	}
}

func TestFloat32(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Float32()
		if err != nil {
			t.Error(err)
		}

		if *got < 0 || *got > 1 {
			t.Errorf("expected a number between 0 and 1, but got '%f'", *got)
		}
	}
}

func TestFloat32InRange(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Float32InRange(50, 200)
		if err != nil {
			t.Error(err)
		}

		if *got < 50 || *got > 200 {
			t.Errorf("expected a number between 50 and 200, but got '%f'", *got)
		}
	}
}

func TestUint64(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		_, err := Uint64() // coverage only
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUint64InRange(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Uint64InRange(50, 200)
		if err != nil {
			t.Error(err)
		}

		if *got < 50 || *got > 200 {
			t.Errorf("expected a number between 50 and 200, but got '%d'", *got)
		}
	}
}

func TestUint32(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		_, err := Uint32() // coverage only
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUint32InRange(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Uint32InRange(50, 200)
		if err != nil {
			t.Error(err)
		}

		if *got < 50 || *got > 200 {
			t.Errorf("expected a number between 50 and 200, but got '%d'", *got)
		}
	}
}

func TestUint16(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		_, err := Uint16() // coverage only
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUint16InRange(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Uint16InRange(50, 200)
		if err != nil {
			t.Error(err)
		}

		if *got < 50 || *got > 200 {
			t.Errorf("expected a number between 50 and 200, but got '%d'", *got)
		}
	}
}

func TestUint8(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		_, err := Uint8() // coverage only
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUint8InRange(t *testing.T) {
	Initialize(NewXoshiro512StarStar())

	for i := 0; i < iterations; i++ {
		got, err := Uint8InRange(50, 200)
		if err != nil {
			t.Error(err)
		}

		if *got < 50 || *got > 200 {
			t.Errorf("expected a number between 50 and 200, but got '%d'", *got)
		}
	}
}

func TestBool(t *testing.T) {
	Initialize(NewXoshiro512StarStar())
	var tResult, fResult int

	for i := 0; i < iterations; i++ {
		got, err := Bool()
		if err != nil {
			t.Error(err)
		}

		switch *got {
		case true:
			tResult++
		default:
			fResult++
		}
	}

	// we should expect a fairly equal distribution of true and false values;
	// a deviation of 5% or more from an equal distribution may cause an
	// occasional failure of this test.
	if float64(tResult) > float64(iterations)*0.55 {
		t.Errorf("got an unexpectedly large number of true values; count=%d", tResult)
	}
}

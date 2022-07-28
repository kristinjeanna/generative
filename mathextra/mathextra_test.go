package mathextra

import (
	"fmt"
	"testing"
)

func TestFloor(t *testing.T) {
	type testCase struct {
		input    float64
		expected float64
	}

	cases := []testCase{
		{1.3, 1.0},
		{1.9, 1.0},
		{-1.3, -1.0},
		{-1.9, -1.0},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%f;expected=%f", tc.input, tc.expected), func(t *testing.T) {
			got := Floor(tc.input)
			if tc.expected != got {
				t.Errorf("expected '%v', but got '%v'", tc.expected, got)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	type testCase struct {
		input    float64
		expected float64
	}

	cases := []testCase{
		// positive inputs
		{HalfPi, HalfPi},
		{Pi, Pi},
		{1.5 * Pi, -0.5 * Pi},
		{2 * Pi, 0},
		{2.5 * Pi, 0.5 * Pi},
		{3 * Pi, Pi},
		{4 * Pi, 0},

		// negative inputs
		{-HalfPi, -HalfPi},
		{-Pi, Pi}, // -π is returned as π
		{-1.5 * Pi, 0.5 * Pi},
		{-2 * Pi, 0},
		{-2.5 * Pi, -0.5 * Pi},
		{-3 * Pi, Pi}, // -π is returned as π
		{-4 * Pi, 0},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%f;expected=%f", tc.input, tc.expected), func(t *testing.T) {
			got := Normalize(tc.input)
			if tc.expected != got {
				t.Errorf("expected '%v', but got '%v'", tc.expected, got)
			}
		})
	}
}

func TestToRadians(t *testing.T) {
	type testCase struct {
		inputDegrees float64
		expected     float64
	}

	cases := []testCase{
		// positive inputs
		{90, HalfPi},
		{180, Pi},
		{270, -0.5 * Pi},
		{360, 0},
		{450, 0.5 * Pi},
		{540, Pi},
		{720, 0},

		// negative inputs
		{-90, -HalfPi},
		{-180, Pi}, // -π is returned as π
		{-270, 0.5 * Pi},
		{-360, 0},
		{-450, -0.5 * Pi},
		{-540, Pi}, // -π is returned as π
		{-720, 0},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("inputDegrees=%f;expected=%f", tc.inputDegrees, tc.expected), func(t *testing.T) {
			got := ToRadians(tc.inputDegrees)
			if tc.expected != got {
				t.Errorf("expected '%v', but got '%v'", tc.expected, got)
			}
		})
	}
}

package misc

import (
	"fmt"
	"testing"

	"github.com/kristinjeanna/generative/errors"

	"golang.org/x/exp/slices"
)

func TestLastNIndexes(t *testing.T) {
	type testCase struct {
		index    int
		n        int
		expected []int
	}

	cases := []testCase{
		{0, 5, []int{0}},
		{1, 5, []int{0, 1}},
		{2, 5, []int{0, 1, 2}},
		{3, 5, []int{0, 1, 2, 3}},
		{4, 5, []int{0, 1, 2, 3, 4}},
		{5, 5, []int{1, 2, 3, 4, 5}},
		{5, 4, []int{2, 3, 4, 5}},
		{5, 3, []int{3, 4, 5}},
		{5, 2, []int{4, 5}},
		{5, 1, []int{5}},
		{3, 4, []int{0, 1, 2, 3}},
		{3, 3, []int{1, 2, 3}},
		{3, 2, []int{2, 3}},
		{3, 1, []int{3}},
		{97, 5, []int{93, 94, 95, 96, 97}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("index=%d;n=%d", tc.index, tc.n), func(t *testing.T) {
			got, err := LastNIndexes(tc.index, tc.n)
			if err != nil {
				t.Error(err)
			}
			if !slices.Equal(tc.expected, got) {
				t.Errorf("expected '%v', but got '%v'", tc.expected, got)
			}
		})
	}
}

func TestLastNIndexes_WithError(t *testing.T) {
	type testCase struct {
		index int
		n     int
	}
	cases := []testCase{
		{10, -1}, // n cannot be less than or equal to zero
		{10, 0},  // n cannot be less than or equal to zero
		{-1, 5},  // index cannot be less than zero
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("index=%d;n=%d", tc.index, tc.n), func(t *testing.T) {
			_, err := LastNIndexes(tc.index, tc.n)
			if err == nil {
				t.Error(errors.ExpectedError)
			}
		})
	}
}

func TestLastNIndexesWrap(t *testing.T) {
	type testCase struct {
		index    int
		n        int
		max      int
		expected []int
	}

	cases := []testCase{
		{100, 5, 1000, []int{96, 97, 98, 99, 100}},
		{0, 5, 100, []int{96, 97, 98, 99, 0}},
		{1, 5, 100, []int{97, 98, 99, 0, 1}},
		{2, 5, 100, []int{98, 99, 0, 1, 2}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("index=%d;n=%d;max=%d", tc.index, tc.n, tc.max), func(t *testing.T) {
			got, err := LastNIndexesWrap(tc.index, tc.n, tc.max)
			if err != nil {
				t.Error(err)
			}
			if !slices.Equal(tc.expected, got) {
				t.Errorf("expected '%v', but got '%v'", tc.expected, got)
			}
		})
	}
}

func TestLastNIndexesWrap_WithError(t *testing.T) {
	type testCase struct {
		index int
		n     int
		max   int
	}
	cases := []testCase{
		{10, -1, 100}, // n cannot be less than or equal to zero
		{10, 0, 100},  // n cannot be less than or equal to zero
		{-1, 5, 100},  // index cannot be less than zero
		{10, 5, -1},   // max cannot be less than zero
		{10, 5, 5},    // max cannot be less than or equal to index
		{10, 5, 10},   // max cannot be less than or equal to index
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("index=%d;n=%d;max=%d; ", tc.index, tc.n, tc.max), func(t *testing.T) {
			_, err := LastNIndexesWrap(tc.index, tc.n, tc.max)
			if err == nil {
				t.Error(errors.ExpectedError)
			}
		})
	}
}

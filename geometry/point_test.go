package geometry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		x, y     float64
		expected *Point
	}

	cases := []testCase{
		{5, 10, &Point{5, 10}},
		{-2, -0.3, &Point{-2, -0.3}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("x=%f;y=%f", tc.x, tc.y), func(t *testing.T) {
			got := NewPoint(tc.x, tc.y)
			assert.NotNil(t, got)
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.expected.X(), got.X())
			assert.Equal(t, tc.expected.Y(), got.Y())
		})
	}
}

func TestAdd(t *testing.T) {
	type testCase struct {
		x1, y1, x2, y2 float64
		expected       *Point
	}

	cases := []testCase{
		{5, 10, 2, 7, &Point{7, 17}},
		{-1, 3, 5, -9, &Point{4, -6}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("x1=%f;y1=%f;x2=%f;y2=%f", tc.x1, tc.y1, tc.x2, tc.y2), func(t *testing.T) {
			p1 := NewPoint(tc.x1, tc.y1)
			p2 := NewPoint(tc.x2, tc.y2)
			got := p1.Add(*p2)
			assert.NotNil(t, got)
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.expected.X(), got.X())
			assert.Equal(t, tc.expected.Y(), got.Y())
		})
	}
}

func TestEquals(t *testing.T) {
	type testCase struct {
		arg      interface{}
		expected bool
	}

	cases := []testCase{
		{&Point{5, 10}, true},
		{Point{5, 10}, true},
		{&Point{-3, 2}, false},
		{"sample text", false},
	}

	point := NewPoint(5, 10)

	for _, tc := range cases {
		t.Run(fmt.Sprintf("arg=%v;expected=%v", tc.arg, tc.expected), func(t *testing.T) {
			got := point.Equals(tc.arg)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDistance(t *testing.T) {
	type testCase struct {
		px, py, expected float64
	}

	cases := []testCase{
		{3, 4, 5},
		{3, -4, 5},
	}

	point := NewPoint(0, 0)

	for _, tc := range cases {
		t.Run(fmt.Sprintf("px=%v;py=%v;py=%v", tc.px, tc.py, tc.expected), func(t *testing.T) {
			got := point.Distance(tc.px, tc.py)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDistanceSq(t *testing.T) {
	type testCase struct {
		px, py, expected float64
	}

	cases := []testCase{
		{3, 4, 25},
		{3, -4, 25},
	}

	point := NewPoint(0, 0)

	for _, tc := range cases {
		t.Run(fmt.Sprintf("px=%v;py=%v;py=%v", tc.px, tc.py, tc.expected), func(t *testing.T) {
			got := point.DistanceSq(tc.px, tc.py)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestCompareTo(t *testing.T) {
	type testCase struct {
		in       Point
		expected int
	}

	cases := []testCase{
		{Point{5, 10}, -1},
		{Point{15, 10}, 1},
		{Point{10, 10}, 0},
		{Point{10, 5}, -1},
		{Point{10, 15}, 1},
	}

	point := NewPoint(10, 10)

	for _, tc := range cases {
		t.Run(fmt.Sprintf("in=%v;expected=%d", tc.in, tc.expected), func(t *testing.T) {
			got := point.CompareTo(tc.in)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	type testCase struct {
		in       Point
		expected string
	}

	cases := []testCase{
		{Point{5, 10}, `{"x":5,"y":10}`},
		{Point{-3, 2}, `{"x":-3,"y":2}`},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("in=%v;expected=%v", tc.in, tc.expected), func(t *testing.T) {
			got, err := tc.in.MarshalJSON()
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, string(got))
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	type testCase struct {
		in       string
		expected Point
	}

	cases := []testCase{
		{`{"x":5,"y":10}`, Point{5, 10}},
		{`{"x":-3,"y":2}`, Point{-3, 2}},
		{`{"a":5,"b":-7}`, Point{0, 0}},
	}

	for _, tc := range cases {
		var got Point
		t.Run(fmt.Sprintf("in=%v;expected=%v", tc.in, tc.expected), func(t *testing.T) {
			err := got.UnmarshalJSON([]byte(tc.in))
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestTranslate(t *testing.T) {
	type testCase struct {
		pointIn    *Point
		fIn        PointTranslator
		distanceIn float64
		expected   *Point
	}

	cases := []testCase{
		{
			&Point{0, 0},
			func(p Point, f float64) Point {
				return Point{p.x + f, p.y + f}
			},
			5,
			&Point{5, 5},
		},
		{
			&Point{-10, 2},
			func(p Point, f float64) Point {
				return Point{p.x + 2*f, p.y - 3*f}
			},
			2,
			&Point{-6, -4},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("pointIn=%v;expected=%v", tc.pointIn, tc.expected), func(t *testing.T) {
			got := tc.pointIn.Translate(tc.fIn, tc.distanceIn)
			assert.NotNil(t, got)
			assert.Equal(t, *tc.expected, got)
		})
	}
}

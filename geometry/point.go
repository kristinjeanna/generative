package geometry

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/kristinjeanna/generative/functional"
)

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

type Point struct {
	x float64
	y float64
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) Add(o Point) *Point {
	return &Point{
		x: p.x + o.x,
		y: p.y + o.y,
	}
}

func (p Point) CompareTo(o Point) int {
	compare := func(a, b float64) int {
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	}

	if p.y == o.y {
		return compare(p.x, o.x)
	} else {
		return compare(p.y, o.y)
	}
}

func (p Point) Distance(px, py float64) float64 {
	px -= p.x
	py -= p.y
	return math.Hypot(px, py)
}

func (p Point) DistanceSq(px, py float64) float64 {
	px -= p.x
	py -= p.y
	return px*px + py*py
}

func (p Point) Equals(o interface{}) bool {
	var isSameType bool
	switch o.(type) {
	case Point:
		isSameType = true
	default:
		isSameType = false
	}

	if !isSameType {
		return false
	}

	p2 := o.(Point)
	return p.x == p2.x && p.y == p2.y
}

func (p Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.x, p.y)
}

type PointTranslator functional.BiFunction[Point, float64, Point]

func (p Point) Translate(f PointTranslator, distance float64) Point {
	return f(p, distance)
}

type RadialPointTranslator functional.BiFunction[PointPair, float64, Point]

func (p *Point) RadialTranslate(f RadialPointTranslator, center *Point, distance float64) Point {
	return f(*NewPointPair(center, p), distance)
}

// TODO: implement Vector
// func (p Point) VectorTo(point Point) Vector {
//     px := point.x - p.x
//     py := point.y - p.y
//     return NewVector(NewPoint(px, py))
// }

// internalPoint is used solely for marshalling/unmarshalling purposes so
// as to keep Point as immutable as possible (i.e., by not needing to export
// Point.x and Point.y).
type internalPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(&internalPoint{
		X: p.x,
		Y: p.y,
	})
}

func (p *Point) UnmarshalJSON(data []byte) error {
	p2 := &internalPoint{}
	if err := json.Unmarshal(data, &p2); err != nil {
		return err
	}

	p.x = p2.X
	p.y = p2.Y

	return nil
}

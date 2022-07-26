package geometry

import "fmt"

func NewPointPair(pt1, pt2 *Point) *PointPair {
	pp := &PointPair{
		pt1: pt1,
		pt2: pt2,
	}

	if pp.pt2.x == pp.pt1.x {
		// no slope or y-intercept; vertical line connects points
		pp.slope = nil
		pp.yIntercept = nil
	} else {
		slope := (pp.pt2.y - pp.pt1.y) / (pp.pt2.x - pp.pt1.x)
		pp.slope = &slope

		yIntercept := pp.pt2.y - slope*pp.pt2.x
		pp.yIntercept = &yIntercept
	}

	return pp
}

type PointPair struct {
	pt1        *Point
	pt2        *Point
	slope      *float64
	yIntercept *float64
}

func (pp PointPair) Point1() *Point {
	return pp.pt1
}

func (pp PointPair) Point2() *Point {
	return pp.pt2
}

func (pp PointPair) Slope() float64 {
	return *pp.slope
}

func (pp PointPair) YIntercept() float64 {
	return *pp.yIntercept
}

func (pp PointPair) Distance() float64 {
	return pp.pt1.Distance(pp.pt2.x, pp.pt2.y)
}

func (pp PointPair) Equals(o interface{}) bool {
	var isSameType bool
	switch o.(type) {
	case PointPair:
		isSameType = true
	default:
		isSameType = false
	}

	if !isSameType {
		return false
	}

	pp2 := o.(PointPair)
	return pp.pt1.Equals(pp2.pt1) && pp.pt2.Equals(pp2.pt2)
}

func (pp PointPair) Midpoint() *Point {
	x := (pp.pt1.x + pp.pt2.x) / 2
	y := (pp.pt1.y + pp.pt2.y) / 2
	return NewPoint(x, y)
}

func (pp PointPair) String() string {
	return fmt.Sprintf("[%f, %f]", pp.pt1, pp.pt2)
}

package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (s *Triangle) CalcPerimeter() float64 {
	return s.Side * 3
}

func (s *Triangle) CalcArea() float64 {
	return math.Pow(s.Side, 2) * math.Sqrt(3.0)/4.0
}
package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return fmt.Errorf("shape capacity limit")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < len(b.shapes) {
		return b.shapes[i], nil
	}
	return nil, fmt.Errorf("this index not found")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < len(b.shapes) {
		r := b.shapes[i]

		copy(b.shapes[i:], b.shapes[i+1:])
		b.shapes = b.shapes[:len(b.shapes)-1]
		return r, nil
	}
	return nil, fmt.Errorf("this index not found")

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("this index not found")
	}
	r := b.shapes[i]
	b.shapes[i] = shape
	return r, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var circles []int
	for id, shape := range b.shapes {
		switch shape.(type) {
		case Circle:
			circles = append(circles, id)
		case *Circle:
			circles = append(circles, id)
		}
	}
	if len(circles) == 0 {
		return fmt.Errorf("no circles in shapes")
	}

	for i := len(circles) - 1; i >= 0; i-- {
		if _, err := b.ExtractByIndex(circles[i]); err != nil {
			return err
		}
	}
	return nil

}

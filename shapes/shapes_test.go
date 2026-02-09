package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{4, 5}
	if !almostEqual(r.Area(), 20) {
		t.Errorf("expected 20, got %v", r.Area())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{4, 5}
	if !almostEqual(r.Perimeter(), 18) {
		t.Errorf("expected 18, got %v", r.Perimeter())
	}
}

func TestRectangleScale(t *testing.T) {
	r := Rectangle{2, 3}
	r.Scale(2)
	if r.Width != 4 || r.Height != 6 {
		t.Errorf("rectangle not scaled correctly")
	}
}

func TestCircleArea(t *testing.T) {
	c := Circle{1}
	if !almostEqual(c.Area(), math.Pi) {
		t.Errorf("incorrect circle area")
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{1}
	if !almostEqual(c.Perimeter(), 2*math.Pi) {
		t.Errorf("incorrect circle perimeter")
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{2}
	c.Scale(3)
	if c.Radius != 6 {
		t.Errorf("circle not scaled correctly")
	}
}

func TestTriangleArea(t *testing.T) {
	tr := Triangle{4, 6}
	if !almostEqual(tr.Area(), 12) {
		t.Errorf("incorrect triangle area")
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tr := Triangle{4, 6}
	if tr.Perimeter() != 12 {
		t.Errorf("incorrect triangle perimeter")
	}
}

func TestTriangleScale(t *testing.T) {
	tr := Triangle{3, 5}
	tr.Scale(2)
	if tr.Base != 6 || tr.Height != 10 {
		t.Errorf("triangle not scaled correctly")
	}
}


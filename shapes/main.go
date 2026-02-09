package main

import "fmt"

func main() {
	rect := Rectangle{Width: 4, Height: 5}
	circle := Circle{Radius: 3}
	triangle := Triangle{Base: 6, Height: 4}

	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())

	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())

	fmt.Printf("Triangle Area: %.2f\n", triangle.Area())
	fmt.Printf("Triangle Perimeter: %.2f\n", triangle.Perimeter())

	rect.Scale(2)
	fmt.Printf("Scaled Rectangle Area: %.2f\n", rect.Area())
}


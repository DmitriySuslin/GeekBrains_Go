package main

import (
	"fmt"
	"github.com/DmitriySuslin/GeekBrains_Go/lesson2/geometry"
)

func main() {
	var rectangle_h, rectangle_l float64
	fmt.Print("Введите длину прямоугольника:")
	fmt.Scan(&rectangle_l)

	fmt.Print("Введите ширину прямоугольника:")
	fmt.Scan(&rectangle_h)

	fmt.Printf("Площадь круга: %4f квадратных\n", geometry.Area_rectangle(rectangle_l,rectangle_h))
}

package main

import (
	"fmt"
	"github.com/DmitriySuslin/GeekBrains_Go/lesson2/geometry"
)

func main()  {
	var area float64

	fmt.Print("Ввети площадь круга:")
	fmt.Scan(&area)

	fmt.Println("Диаметр окружности: ", geometry.Diameter_arena(area))
	fmt.Println("Длина окружности: ", geometry.Length_arena(area))
	
}

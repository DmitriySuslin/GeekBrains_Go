package geometry

import "math"

// Area_circle ............
func Area_circle(radius float64) float64 {
	return radius * radius * math.Pi
}

// Diameter_arena ...........
func Diameter_arena(area float64) float64 {
	return math.Sqrt(4 * area / math.Pi)
}

// Length_arena ...........
func Length_arena(area float64) float64 {
	return math.Sqrt(4 * area * math.Pi)
}
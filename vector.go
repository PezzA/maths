package maths

import "math"

// Vector2D represents a 2d vector that holds direction and magnitude
type Vector2D struct {
	X float64
	Y float64
}

//Normalised returns the normalised unit versin of a Vector2D
func (v Vector2D) Normalised() Vector2D {
	return v.DivideBy(v.GetLength())
}

// GetLength gets the length of the vector
func (v Vector2D) GetLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// GetLengthSquared gets the length of the vector, squared.
// This omits the Sqrt call in GetLength and will be faster when comparing lengths by size
func (v Vector2D) GetLengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// MultiplyBy will multiply a vector by a given speed
func (v Vector2D) MultiplyBy(s float64) Vector2D {
	return Vector2D{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// DivideBy will divide a vector by a given speed
func (v Vector2D) DivideBy(s float64) Vector2D {
	return Vector2D{
		X: v.X / s,
		Y: v.Y / s,
	}
}

// Add returns the sum of the receiver with the passed in vector
func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

// Subtract returns the result of the subtraction of the receiver with the passed in vector
func (v Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2D) dotProduct(v2 Vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

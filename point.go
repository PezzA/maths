package maths

// Point2D contains a point on a 2D grid
type Point2D struct {
	X float64
	Y float64
}

// AddVector adds a Vector2D to the receiver Point2D
func (p Point2D) AddVector(v Vector2D) Point2D {
	return Point2D{
		p.X + v.X,
		p.Y + v.Y,
	}
}

// GetVectorToPoint returns the vector to travel from the receiver to the target
func (p Point2D) GetVectorToPoint(p2 Point2D) Vector2D {
	return Vector2D{
		p2.X - p.X,
		p2.Y - p.Y,
	}
}

package utils

type Point struct {
	X, Y float64
}

// left top and right bottom
type Rectangle struct {
	Left, Right Point
}

func IsOverlappingPoint(r1, r2 Rectangle) bool {
	if r1.Right.X < r2.Left.X ||
		r1.Left.X > r2.Right.X ||
		r1.Right.Y < r2.Left.Y ||
		r1.Left.Y > r2.Right.Y {
		return false
	}

	return true
}

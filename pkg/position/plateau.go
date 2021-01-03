package position

type Plateau struct {
	LowerBoundCoordinateX int
	LowerBoundCoordinateY int
	UpperBoundCoordinateX int
	UpperBoundCoordinateY int
}

func NewPlateau(upperBoundCoordinateX, upperBoundCoordinateY int) *Plateau {
	p := &Plateau{}
	p.UpperBoundCoordinateX = upperBoundCoordinateX
	p.UpperBoundCoordinateY = upperBoundCoordinateY
	return p
}

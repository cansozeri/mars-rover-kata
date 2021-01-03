package direction

import (
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/state"
)

type North struct {
	Coordinates *position.Coordinates
}

func (n *North) MoveForward() (s state.State, e error) {
	n.Coordinates.Y++
	return n, nil
}

func (n *North) TurnLeft() (s state.State, e error) {
	return &West{}, nil
}

func (n *North) TurnRight() (s state.State, e error) {
	return &East{}, nil
}

func (n *North) SetPosition(p *position.Coordinates) {
	n.Coordinates = p
}

func (n *North) GetPosition() *position.Coordinates {
	return n.Coordinates
}

func (n *North) String() string {
	return "N"
}

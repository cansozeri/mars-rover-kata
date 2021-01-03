package direction

import (
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/state"
)

type East struct {
	Coordinates *position.Coordinates
}

func (e *East) MoveForward() (s state.State, err error) {
	e.Coordinates.X++
	return e, nil
}

func (e *East) TurnLeft() (s state.State, err error) {
	return &North{}, nil
}

func (e *East) TurnRight() (s state.State, err error) {
	return &South{}, nil
}

func (e *East) SetPosition(p *position.Coordinates) {
	e.Coordinates = p
}

func (e *East) GetPosition() *position.Coordinates {
	return e.Coordinates
}

func (e *East) String() string {
	return "E"
}

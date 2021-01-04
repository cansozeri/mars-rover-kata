package direction

import (
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/state"
)

type West struct {
	Coordinates *position.Coordinates
}

func (w *West) MoveForward() (s state.State, e error) {
	w.Coordinates.X--
	return w, nil
}

func (w *West) TurnLeft90Degrees() (s state.State, e error) {
	return &South{}, nil
}

func (w *West) TurnRight90Degrees() (s state.State, e error) {
	return &North{}, nil
}

func (w *West) SetPosition(p *position.Coordinates) {
	w.Coordinates = p
}

func (w *West) GetPosition() *position.Coordinates {
	return w.Coordinates
}
func (w *West) String() string {
	return "W"
}

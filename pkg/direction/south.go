package direction

import (
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/state"
)

type South struct {
	Coordinates *position.Coordinates
}

func (s *South) MoveForward() (st state.State, e error) {
	s.Coordinates.Y--
	return s, nil
}

func (s *South) TurnLeft90Degrees() (st state.State, e error) {
	return &East{}, nil
}

func (s *South) TurnRight90Degrees() (st state.State, e error) {
	return &West{}, nil
}

func (s *South) SetPosition(p *position.Coordinates) {
	s.Coordinates = p
}

func (s *South) GetPosition() *position.Coordinates {
	return s.Coordinates
}

func (s *South) String() string {
	return "S"
}

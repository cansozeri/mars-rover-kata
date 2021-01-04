package state

import "mars-rover-kata/pkg/position"

type State interface {
	MoveForward() (State, error)
	TurnLeft90Degrees() (State, error)
	TurnRight90Degrees() (State, error)
	SetPosition(*position.Coordinates)
	GetPosition() *position.Coordinates
	String() string
}

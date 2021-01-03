package state

import "mars-rover-kata/pkg/position"

type State interface {
	MoveForward() (State, error)
	TurnLeft() (State, error)
	TurnRight() (State, error)
	SetPosition(*position.Coordinates)
	GetPosition() *position.Coordinates
	String() string
}

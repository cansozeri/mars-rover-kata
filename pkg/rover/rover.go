package rover

import (
	"fmt"
	"log"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/state"
)

const (
	Disabled = "Disabled"
	Active   = "Active"
)

type Rover struct {
	north state.State
	south state.State
	east  state.State
	west  state.State

	currentState state.State
	plateau      *position.Plateau
	Coordinates  *position.Coordinates
	Status       string
}

func NewRover() *Rover {
	r := &Rover{
		Coordinates: &position.Coordinates{
			X: 0,
			Y: 0,
		},
	}

	north := &direction.North{}

	south := &direction.South{}

	east := &direction.East{}

	west := &direction.West{}

	r.north = north
	r.south = south
	r.east = east
	r.west = west
	r.Status = Disabled

	return r
}

func (r *Rover) SetInitialPosition(p *position.Plateau, x, y int, orientation state.State) error {

	if x > p.UpperBoundCoordinateX || y > p.UpperBoundCoordinateY {
		return fmt.Errorf("coordinate limit of x (%v) - y (%v) has been exceeded in either or both of the supplied inputs", p.UpperBoundCoordinateX, p.UpperBoundCoordinateY)
	}

	r.plateau = p
	r.Coordinates.X = x
	r.Coordinates.Y = y
	r.currentState = orientation
	r.Status = Active

	return nil
}

func (r *Rover) MoveForward() (s state.State) {

	r.currentState.SetPosition(r.Coordinates)

	s, e := r.currentState.MoveForward()

	if e != nil {
		log.Fatalf(e.Error())
	}

	r.Coordinates = r.currentState.GetPosition()

	r.validateLocation()

	return r.currentState
}

func (r *Rover) TurnLeft() (s state.State) {

	r.currentState.SetPosition(r.Coordinates)

	s, e := r.currentState.TurnLeft()

	if e != nil {
		log.Fatalf(e.Error())
	}

	r.SetState(s)
	return r.currentState
}

func (r *Rover) TurnRight() (s state.State) {

	r.currentState.SetPosition(r.Coordinates)

	s, e := r.currentState.TurnRight()

	if e != nil {
		log.Fatalf(e.Error())
	}

	r.SetState(s)
	return r.currentState
}

func (r *Rover) Process(commands ...string) error {

	if r.Status == Disabled {
		log.Fatalf("Please set the rover initial place on the plateau")
	}
	for _, command := range commands {
		switch command {
		case "L":
			r.TurnLeft()
		case "R":
			r.TurnRight()
		case "M":
			r.MoveForward()
		default:
			return fmt.Errorf("unknown command: %q, ignoring it", command)
		}
	}

	return nil
}

func (r *Rover) SetState(s state.State) {
	r.currentState = s
}

func (r *Rover) GetState() state.State {
	return r.currentState
}

func (r *Rover) GetPlateau() *position.Plateau {
	return r.plateau
}

func (r *Rover) DisplayRobotStats() string {
	return fmt.Sprintf("%d %d %s", r.Coordinates.X, r.Coordinates.Y, r.currentState.String())
}

func (r *Rover) validateLocation() {
	if r.Coordinates.X > r.plateau.UpperBoundCoordinateX ||
		r.Coordinates.Y > r.plateau.UpperBoundCoordinateY ||
		r.Coordinates.X < r.plateau.LowerBoundCoordinateX ||
		r.Coordinates.Y < r.plateau.LowerBoundCoordinateY {
		log.Fatalf("The Rover cannot be out of bounds of the plateau! %v", r.Coordinates)
	}
}

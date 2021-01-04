package direction_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"testing"
)

func TestNorth_TurnLeft(t *testing.T) {
	t.Parallel()

	d := direction.North{}
	newDirection, e := d.TurnLeft90Degrees()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.West{}, newDirection)
}

func TestNorth_TurnRight(t *testing.T) {
	t.Parallel()

	d := direction.North{}
	newDirection, e := d.TurnRight90Degrees()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.East{}, newDirection)
}

func TestNorth_MoveForward(t *testing.T) {
	t.Parallel()

	d := direction.North{
		Coordinates: &position.Coordinates{
			X: 0,
			Y: 0,
		},
	}

	newDirection, e := d.MoveForward()

	assert.Equal(t, nil, e)
	assert.Equal(t, 1, newDirection.GetPosition().Y)
}

func TestNorth_String(t *testing.T) {
	t.Parallel()

	d := direction.North{}

	assert.Equal(t, "N", d.String())
}

func TestNorth_SetPosition(t *testing.T) {
	t.Parallel()

	d := direction.North{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.Coordinates)
}

func TestNorth_GetPosition(t *testing.T) {
	t.Parallel()

	d := direction.North{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.GetPosition())
}

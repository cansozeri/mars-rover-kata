package direction_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"testing"
)

func TestWest_TurnLeft(t *testing.T) {
	t.Parallel()

	d := direction.West{}
	newDirection, e := d.TurnLeft90Degrees()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.South{}, newDirection)
}

func TestWest_TurnRight(t *testing.T) {
	t.Parallel()

	d := direction.West{}
	newDirection, e := d.TurnRight90Degrees()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.North{}, newDirection)
}

func TestWest_MoveForward(t *testing.T) {
	t.Parallel()

	d := direction.West{
		Coordinates: &position.Coordinates{
			X: 1,
			Y: 0,
		},
	}

	newDirection, e := d.MoveForward()

	assert.Equal(t, nil, e)
	assert.Equal(t, 0, newDirection.GetPosition().X)
}

func TestWest_String(t *testing.T) {
	t.Parallel()

	d := direction.West{}

	assert.Equal(t, "W", d.String())
}

func TestWest_SetPosition(t *testing.T) {
	t.Parallel()

	d := direction.West{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.Coordinates)
}

func TestWest_GetPosition(t *testing.T) {
	t.Parallel()

	d := direction.West{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.GetPosition())
}

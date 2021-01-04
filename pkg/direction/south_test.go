package direction_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"testing"
)

func TestSouth_TurnLeft(t *testing.T) {
	t.Parallel()

	d := direction.South{}
	newDirection, e := d.TurnLeft90Degrees()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.East{}, newDirection)
}

func TestSouth_TurnRight(t *testing.T) {
	t.Parallel()

	d := direction.South{}
	newDirection, e := d.TurnRight90Degrees()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.West{}, newDirection)
}

func TestSouth_MoveForward(t *testing.T) {
	t.Parallel()

	d := direction.South{
		Coordinates: &position.Coordinates{
			X: 0,
			Y: 1,
		},
	}

	newDirection, e := d.MoveForward()

	assert.Equal(t, nil, e)
	assert.Equal(t, 0, newDirection.GetPosition().Y)
}

func TestSouth_String(t *testing.T) {
	t.Parallel()

	d := direction.South{}

	assert.Equal(t, "S", d.String())
}

func TestSouth_SetPosition(t *testing.T) {
	t.Parallel()

	d := direction.South{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.Coordinates)
}

func TestSouth_GetPosition(t *testing.T) {
	t.Parallel()

	d := direction.South{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.GetPosition())
}

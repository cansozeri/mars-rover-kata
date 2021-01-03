package direction_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"testing"
)

func TestEast_TurnLeft(t *testing.T) {
	t.Parallel()

	d := direction.East{}
	newDirection, e := d.TurnLeft()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.North{}, newDirection)
}

func TestEast_TurnRight(t *testing.T) {
	t.Parallel()

	d := direction.East{}
	newDirection, e := d.TurnRight()

	assert.Equal(t, nil, e)
	assert.Equal(t, &direction.South{}, newDirection)
}

func TestEast_MoveForward(t *testing.T) {
	t.Parallel()

	d := direction.East{
		Coordinates: &position.Coordinates{
			X: 0,
			Y: 0,
		},
	}

	newDirection, e := d.MoveForward()

	assert.Equal(t, nil, e)
	assert.Equal(t, 1, newDirection.GetPosition().X)
}

func TestEast_String(t *testing.T) {
	t.Parallel()

	d := direction.East{}

	assert.Equal(t, "E", d.String())
}

func TestEast_SetPosition(t *testing.T) {
	t.Parallel()

	d := direction.East{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.Coordinates)
}

func TestEast_GetPosition(t *testing.T) {
	t.Parallel()

	d := direction.East{}

	p := &position.Coordinates{
		X: 2,
		Y: 1,
	}

	d.SetPosition(p)

	assert.Equal(t, p, d.GetPosition())
}

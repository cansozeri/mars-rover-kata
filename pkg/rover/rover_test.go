package rover_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/rover"
	"testing"
)

func TestNewRover(t *testing.T) {
	t.Parallel()

	r := rover.NewRover()

	assert.Equal(t, 0, r.Coordinates.X)
	assert.Equal(t, 0, r.Coordinates.Y)
	assert.Equal(t, rover.Disabled, r.Status)
}

func TestRover_SetInitialPosition(t *testing.T) {
	t.Parallel()

	plateau := position.NewPlateau(5, 5)

	r := rover.NewRover()

	assert.Nil(t, r.SetInitialPosition(plateau, 1, 2, &direction.North{}))

	assert.Equal(t, plateau, r.GetPlateau())
	assert.Equal(t, 1, r.Coordinates.X)
	assert.Equal(t, 2, r.Coordinates.Y)
	assert.Equal(t, &direction.North{}, r.GetState())
}

func TestRover_SetInitialPosition_ReturnsError_If_Coordinate_Limit_Exceeded(t *testing.T) {
	t.Parallel()

	plateau := position.NewPlateau(5, 5)

	r := rover.NewRover()

	assert.NotNil(t, r.SetInitialPosition(plateau, 1, 8, &direction.North{}))
}

func TestRover_SetState(t *testing.T) {
	t.Parallel()

	r := rover.NewRover()

	r.SetState(&direction.West{})

	assert.Equal(t, &direction.West{}, r.GetState())
}

func TestRover_Process(t *testing.T) {
	t.Parallel()

	plateau := position.NewPlateau(5, 5)

	r := rover.NewRover()

	assert.Nil(t, r.SetInitialPosition(plateau, 1, 2, &direction.North{}))

	assert.Nil(t, r.Process("L"))

	assert.Equal(t, "W", r.GetState().String())

	assert.Nil(t, r.Process("M", "R"))

	assert.Equal(t, 0, r.Coordinates.X)
	assert.Equal(t, 2, r.Coordinates.Y)
	assert.Equal(t, "N", r.GetState().String())
}

func TestRover_DisplayRobotStats(t *testing.T) {
	t.Parallel()

	plateau := position.NewPlateau(5, 5)

	r := rover.NewRover()

	assert.Nil(t, r.SetInitialPosition(plateau, 1, 2, &direction.North{}))

	assert.Equal(t, "1 2 N", r.DisplayRobotStats())
}

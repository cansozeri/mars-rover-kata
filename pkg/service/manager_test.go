package service_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/rover"
	"mars-rover-kata/pkg/service"
	"testing"
)

func TestInstructions_GetInstructions_Empty_Message(t *testing.T) {
	t.Parallel()

	i := &service.Instructions{}
	e := i.GetInstructions("")

	assert.Nil(t, e)
}

func TestInstructions_GetInstructions_Plateau(t *testing.T) {
	t.Parallel()

	i := &service.Instructions{}
	e := i.GetInstructions("5 5")

	assert.Nil(t, e)
	assert.Equal(t, i.GetPlateau().UpperBoundCoordinateX, 5)
	assert.Equal(t, i.GetPlateau().UpperBoundCoordinateY, 5)
}

func TestInstructions_GetInstructions_Rover_Coordinates(t *testing.T) {
	t.Parallel()

	i := &service.Instructions{}
	e := i.GetInstructions("5 5")
	assert.Nil(t, e)
	e = i.GetInstructions("1 2 N")

	assert.Nil(t, e)
	assert.Equal(t, i.GetRover().Coordinates.X, 1)
	assert.Equal(t, i.GetRover().Coordinates.Y, 2)
	assert.Equal(t, i.GetRover().GetState().String(), "N")
}

func TestInstructions_GetInstructions(t *testing.T) {
	t.Parallel()

	i := &service.Instructions{}

	e := i.GetInstructions("5 5")
	assert.Nil(t, e)

	e = i.GetInstructions("1 2 N")
	assert.Nil(t, e)

	e = i.GetInstructions("LMLMLMLMM")

	assert.Nil(t, e)

	assert.Equal(t, i.GetRover().Coordinates.X, 1)
	assert.Equal(t, i.GetRover().Coordinates.Y, 3)
	assert.Equal(t, i.GetRover().GetState().String(), "N")
}

func TestInstructions_GetInstructions_Invalid_Instruction(t *testing.T) {
	t.Parallel()

	i := &service.Instructions{}
	r := rover.NewRover()
	r.Status = rover.Active
	i.SetRover(r)
	e := i.GetInstructions("1 something")
	assert.NotNil(t, e)
}

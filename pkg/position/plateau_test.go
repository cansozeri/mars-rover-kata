package position_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-kata/pkg/position"
	"testing"
)

func TestNewPlateau(t *testing.T) {
	t.Parallel()

	p := position.NewPlateau(5, 5)

	assert.Equal(t, 5, p.UpperBoundCoordinateY)
	assert.Equal(t, 5, p.UpperBoundCoordinateY)
	assert.Equal(t, 0, p.LowerBoundCoordinateX)
	assert.Equal(t, 0, p.LowerBoundCoordinateY)
}

package main

import (
	"fmt"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/rover"
)

func main() {
	plateau := position.NewPlateau(5, 5)

	r := rover.NewRover()

	r.SetInitialPosition(plateau, 1, 2, &direction.North{})
	r.Process("L", "M", "L", "M", "L", "M", "L", "M", "M")
	fmt.Println(r.DisplayRobotStats())

	r.SetInitialPosition(plateau, 3, 3, &direction.East{})
	r.Process("M", "M", "R", "M", "M", "R", "M", "R", "R", "M")
	fmt.Println(r.DisplayRobotStats())
}

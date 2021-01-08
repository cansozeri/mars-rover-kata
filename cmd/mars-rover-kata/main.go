package main

import (
	"mars-rover-kata/pkg/service"
)

func main() {
	m := &service.Manager{}
	m.LoadAndExecuteInstructions()
}

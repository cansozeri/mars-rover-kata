package main

import (
	"mars-rover-kata/pkg/service"
)

func main() {
	i := &service.Instructions{}
	i.LoadAndExecuteInstructions()
}

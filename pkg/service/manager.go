package service

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"mars-rover-kata/pkg/direction"
	"mars-rover-kata/pkg/position"
	"mars-rover-kata/pkg/rover"
	"mars-rover-kata/pkg/state"
	"os"
	"strconv"
	"strings"
)

var StringToOrientation = map[string]state.State{
	"N": &direction.North{},
	"E": &direction.East{},
	"S": &direction.South{},
	"W": &direction.West{},
}

const (
	InitPlateau = iota + 2
	InitRoverPosition

	failedToParseCoordinates = "failed to parse coordinates, due to %v"
)

type Instructions struct {
	plateau *position.Plateau
	rover   *rover.Rover
}

func (i *Instructions) LoadAndExecuteInstructions() {
	filePtr := flag.String("fPath", "instructions.txt", "service path to read from")
	flag.Parse()
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := i.GetInstructions(scanner.Text()); err != nil {
			fmt.Println(err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (i *Instructions) GetInstructions(text string) error {
	message := strings.Split(strings.Replace(text, " ", "", -1), "")
	if len(message) == 0 {
		return nil
	}

	switch {
	case len(message) == InitPlateau:
		horizontalAxis, verticalAxis, err := getXAndYAxisFromText(message[0], message[1])
		if err != nil {
			return err
		}

		plateau := position.NewPlateau(horizontalAxis, verticalAxis)
		i.plateau = plateau

	case len(message) == InitRoverPosition:
		horizontalAxis, verticalAxis, err := getXAndYAxisFromText(message[0], message[1])
		if err != nil {
			return err
		}

		orientation, ok := StringToOrientation[message[2]]

		if !ok {
			return fmt.Errorf("unknown direction: %q", message[2])
		}

		i.rover = rover.NewRover()
		err = i.rover.SetInitialPosition(i.plateau, horizontalAxis, verticalAxis, orientation)

		if err != nil {
			return err
		}

	default:
		err := i.rover.Process(message...)
		if err != nil {
			return err
		}
		fmt.Println(i.rover.DisplayRobotStats())
	}

	return nil
}

func (i *Instructions) SetPlateau(p *position.Plateau) {
	i.plateau = p
}

func (i *Instructions) GetPlateau() *position.Plateau {
	return i.plateau
}

func (i *Instructions) SetRover(r *rover.Rover) {
	i.rover = r
}

func (i *Instructions) GetRover() *rover.Rover {
	return i.rover
}

func getXAndYAxisFromText(horizontal, vertical string) (int, int, error) {
	horizontalAxis, err := strconv.Atoi(horizontal)
	if err != nil {
		log.Fatalf(failedToParseCoordinates, err)
	}

	verticalAxis, err := strconv.Atoi(vertical)
	if err != nil {
		log.Fatalf(failedToParseCoordinates, err)
	}

	return horizontalAxis, verticalAxis, nil
}

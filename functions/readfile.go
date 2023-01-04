package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X    int
	Y    int
}

type Connection struct {
	Room1 string
	Room2 string
}

type AntFarm struct {
	Ants        int
	Start       Room
	End         Room
	Rooms       []Room
	Connections []Connection
}

func ReadAntFarmFile(filename string) (*AntFarm, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new AntFarm struct
	antFarm := &AntFarm{}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			// This is the line with the number of ants
			Ants, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			antFarm.Ants = Ants
		} else if strings.HasPrefix(line, "##") {
			// This is a directive (e.g. ##start, ##end)
			if line == "##start" {
				// The next line will contain the starting room
				scanner.Scan()
				startRoom, err := parseRoom(scanner.Text())
				if err != nil {
					return nil, err
				}
				antFarm.Start = *startRoom
			} else if line == "##end" {
				// The next line will contain the ending room
				scanner.Scan()
				endRoom, err := parseRoom(scanner.Text())
				if err != nil {
					return nil, err
				}
				antFarm.End = *endRoom
			}
		} else if strings.Contains(line, "-") {
			// This is a connection
			connection, err := parseConnection(line)
			if err != nil {
				return nil, err
			}
			antFarm.Connections = append(antFarm.Connections, *connection)
		} else {
			// This is a room
			room, err := parseRoom(line)
			if err != nil {
				return nil, err
			}
			antFarm.Rooms = append(antFarm.Rooms, *room)
		}
		lineNum++
	}

	return antFarm, nil
}

func parseRoom(line string) (*Room, error) {
	// Split the line on whitespace
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid room format: %s", line)
	}

	// Parse the x and y coordinates
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	return &Room{Name: parts[0], X: x, Y: y}, nil
}

func parseConnection(line string) (*Connection, error) {
	// Split the line on the '-' character
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid connection format: %s", line)
	}

	return &Connection{Room1: parts[0], Room2: parts[1]}, nil
}

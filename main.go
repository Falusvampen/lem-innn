package main

import (
	"fmt"
	"lem-in/functions"
	"os"
)

func main() {
	// Read the ant farm file
	antFarm, err := functions.ReadAntFarmFile("./examples/example04.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// // Print the AntFarm struct
	// fmt.Printf("%+v\n", antFarm)

	// Print the number of ants
	fmt.Printf("Number of ants: %d\n", antFarm.Ants)

	fmt.Println("Start room:")
	fmt.Printf("%s (%d, %d)\n", antFarm.Start.Name, antFarm.Start.X, antFarm.Start.Y)
	fmt.Println("End room:")
	fmt.Printf("%s (%d, %d)\n", antFarm.End.Name, antFarm.End.X, antFarm.End.Y)
	// Print the rooms
	fmt.Println("Rooms:")
	for _, room := range antFarm.Rooms {
		fmt.Printf("%s (%d, %d)\n", room.Name, room.X, room.Y)
	}
	// Print the connections
	fmt.Println("Connections:")
	for _, connection := range antFarm.Connections {
		fmt.Printf("%s <-> %s\n", connection.Room1, connection.Room2)
	}
}

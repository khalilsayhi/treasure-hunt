package main

import (
	"github.com/khalilsayhi/treasure-hunt/gameservice"
	"github.com/khalilsayhi/treasure-hunt/utils"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("No argument provided, make sure to provide the map path as an argument (make FILEPATH=<map_path>)")
	}

	mapPath := os.Args[1]

	board := gameservice.NewBoard(mapPath)

	board.RunSimulation()

	utils.WriteSimulationResultToFile(board.Cells, board.Width, board.Height, mapPath)
}

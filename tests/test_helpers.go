package tests

import (
	"github.com/khalilsayhi/treasure-hunt/domain/types"
	"github.com/khalilsayhi/treasure-hunt/gameservice"
)

func InitBoardWithoutConfig(height, width int) *gameservice.Board {
	board := &gameservice.Board{
		Width:  width,
		Height: height,
		Cells:  make([][]types.Cell, height),
	}
	for i := range board.Cells {
		board.Cells[i] = make([]types.Cell, width)
	}

	return board
}

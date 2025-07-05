package gameservice

import (
	"fmt"
	"github.com/khalilsayhi/treasure-hunt/domain/constants"
	"github.com/khalilsayhi/treasure-hunt/domain/enum"
	"github.com/khalilsayhi/treasure-hunt/domain/types"
	"github.com/khalilsayhi/treasure-hunt/utils"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type Board struct {
	Width          int
	Height         int
	Cells          [][]types.Cell
	AdventurerList []*types.Adventurer
}

func NewBoard(mapPath string) *Board {
	width, height, landscapeConfLines, err := utils.GetBoardInitialConfig(mapPath)
	if err != nil {
		panic(constants.ErrorBoardConfigLoading)
	}

	cells := make([][]types.Cell, height)
	for ix := range cells {
		cells[ix] = make([]types.Cell, width)
	}

	board := &Board{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
	if err = board.InitLandscapeConfig(landscapeConfLines); err != nil {
		panic(err)
	}

	return board
}

func (b *Board) AreMovesAvailable() bool {
	return len(b.AdventurerList) > 0
}

func (b *Board) RunSimulation() {
	for b.AreMovesAvailable() {
		for ix, adventurer := range b.AdventurerList {
			if len(adventurer.Path) == 0 {
				// Remove adventurer if no more moves
				b.AdventurerList = append(b.AdventurerList[:ix], b.AdventurerList[ix+1:]...)

				continue
			}

			oldX := adventurer.X
			oldY := adventurer.Y

			nextMove, err := adventurer.GetNextMove()
			if err != nil {
				log.Printf("Error getting next move for adventurer %s: %v\n", adventurer.Name, err)

				continue
			}

			//Handle Rotation
			if utils.IsRotation(nextMove) {
				// Handle rotation
				adventurer.Rotate(utils.GetNextOrientation(adventurer.Orientation, nextMove))
				b.RerenderBoardAdventurer(adventurer, oldX, oldY, oldX, oldY)

				continue
			}

			// Handle movement
			newX, newY := utils.GetNextCoordinates(adventurer.X, adventurer.Y, adventurer.Orientation)
			if !b.IsCoordinatesInBoard(newX, newY) || b.HasCellMountain(newX, newY) {
				// Adventurer cannot move out of bounds or into mountain but consumes the move
				adventurer.ConsumeMove()
				b.RerenderBoardAdventurer(adventurer, oldX, oldY, oldX, oldY)

				continue
			}
			if b.HasCellAdventurer(newX, newY) {
				// Adventurer cannot move to a cell with another adventurer and passes turn

				continue
			}

			// Adventurer can move in all other cases
			adventurer.Move(newX, newY)

			if b.HasCellTreasure(newX, newY) {
				// Adventurer moves to a cell with treasure
				b.LootTreasure(newX, newY, adventurer)
			}

			b.RerenderBoardAdventurer(adventurer, newX, newY, oldX, oldY)
		}
	}
}

// ------- Cell Content Checkers -------

func (b *Board) IsCoordinatesInBoard(x, y int) bool {
	return x >= 0 && x < b.Width && y >= 0 && y < b.Height
}
func (b *Board) HasCellMountain(x, y int) bool {
	return b.Cells[y][x].Mountain != nil
}

func (b *Board) HasCellTreasure(x, y int) bool {
	return b.Cells[y][x].Treasure != nil && b.Cells[y][x].Treasure.Amount > 0
}

func (b *Board) HasCellAdventurer(x, y int) bool {
	return b.Cells[y][x].Adventurer != nil
}

// ------ Board Rerendering ------

// RerenderBoardAdventurer updates the board's cell content for an adventurer's movement.
func (b *Board) RerenderBoardAdventurer(adventurer *types.Adventurer, newX, newY, oldX, oldY int) {
	b.Cells[oldY][oldX].Adventurer = nil
	b.Cells[newY][newX].Adventurer = adventurer
}

func (b *Board) LootTreasure(x, y int, adventurer *types.Adventurer) {
	treasure := b.Cells[y][x].Treasure
	if treasure == nil {
		log.Printf("No treasure found at (%d, %d) for adventurer %s", x, y, adventurer.Name)

		return
	}

	adventurer.TreasureCount++

	treasure.Amount--
	if treasure.Amount == 0 {
		b.Cells[y][x].Treasure = nil
	}
}

// ------ Landscape Configuration ------

func (b *Board) InitLandscapeConfig(entityConfList []string) error {
	for _, line := range entityConfList {
		if err := b.PlaceEntityIntoBoard(line); err != nil {
			return err
		}
	}

	return nil
}

func (b *Board) PlaceEntityIntoBoard(line string) error {
	flag := line[0]
	switch flag {
	case 'M': // Mountain
		var x, y int
		if _, err := fmt.Sscanf(line, constants.MountainFormat, &x, &y); err != nil {
			return errors.Wrapf(constants.ErrorConfigLandscapeLineFormat, "line: %s, error: %+v", line, err)
		}

		if err := b.SetCellContent(x, y, types.NewMountain()); err != nil {
			return err
		}
	case 'T': // Treasure
		var x, y, amount int
		if _, err := fmt.Sscanf(line, constants.TreasureFormat, &x, &y, &amount); err != nil {
			return errors.Wrapf(constants.ErrorConfigLandscapeLineFormat, "line: %s, error: %+v", line, err)
		}
		if err := b.SetCellContent(x, y, types.NewTreasure(amount)); err != nil {
			return err
		}
	case 'A': // Adventurer
		var name string
		var x, y int
		var path string
		var orientation enum.PlayerOrientation
		if _, err := fmt.Sscanf(line, constants.AdventurerFormat, &name, &x, &y, &orientation, &path); err != nil {
			return errors.Wrapf(constants.ErrorConfigLandscapeLineFormat, "line: %s, error: %+v", line, err)
		}
		adventurer := types.NewAdventurer(name, orientation, x, y, strings.Split(path, ""))
		if err := b.SetCellContent(x, y, adventurer); err != nil {
			return err
		}
		b.AdventurerList = append(b.AdventurerList, adventurer)
	default:
		return errors.Wrap(constants.ErrorUnknownConfigurationFlag, line)
	}

	return nil
}

func (b *Board) SetCellContent(x, y int, content types.CellContent) error {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return errors.Wrapf(constants.ErrorCoordinateOutOfBounds, "(%d, %d)", x, y)
	}
	switch mapObject := (content).(type) {
	case *types.Adventurer:
		b.Cells[y][x].Adventurer = mapObject
	case *types.Mountain:
		b.Cells[y][x].Mountain = mapObject
	case *types.Treasure:
		b.Cells[y][x].Treasure = mapObject
	default:
		return errors.Wrapf(constants.ErrorCellContentNotValid, "%T", content)
	}

	return nil
}

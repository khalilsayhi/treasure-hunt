package tests

import (
	"github.com/khalilsayhi/treasure-hunt/domain/constants"
	"github.com/khalilsayhi/treasure-hunt/domain/enum"
	"github.com/khalilsayhi/treasure-hunt/domain/types"
	"github.com/khalilsayhi/treasure-hunt/gameservice"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BoardTestSuite struct {
	suite.Suite
	board *gameservice.Board
}

func (suite *BoardTestSuite) SetupTest() {
	suite.board = gameservice.NewBoard("board_test_map.txt")
}

func (suite *BoardTestSuite) TestMountainBlock() {
	moves := []enum.PlayerMove{
		enum.Avancer,
	}
	initialX, initialY := 1, 1
	suite.board.Cells[initialY][initialX].Adventurer.Orientation = enum.Nord
	suite.board.Cells[initialY][initialX].Adventurer.Path = moves

	suite.board.RunSimulation()

	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.X, initialX)
	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.Y, initialY)
	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.Name, "Lara")
}

func (suite *BoardTestSuite) TestAdventurerBlock() {
	moves := []enum.PlayerMove{
		enum.Avancer,
	}

	suite.board.Cells[1][1].Adventurer.Path = moves

	initialX, initialY := 0, 1
	newAdventurer := types.NewAdventurer("John", "E", initialX, initialY, []string{"A"})
	suite.board.Cells[1][0].Adventurer = newAdventurer

	suite.board.RunSimulation()

	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.X, initialX)
	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.Y, initialY)
	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.Name, "John")
}

func (suite *BoardTestSuite) TestTreasureCollection() {
	moves := []enum.PlayerMove{
		enum.Avancer,
	}
	initialX, initialY := 1, 1
	suite.board.Cells[initialY][initialX].Adventurer.Orientation = enum.Ouest
	suite.board.Cells[initialY][initialX].Adventurer.Path = moves
	suite.board.Cells[1][0].Treasure = types.NewTreasure(2)
	assert.Equal(suite.T(), suite.board.Cells[initialY][initialX].Adventurer.TreasureCount, 0)
	assert.Equal(suite.T(), suite.board.Cells[1][0].Treasure.Amount, 2)

	suite.board.RunSimulation()

	assert.Equal(suite.T(), suite.board.Cells[initialY][0].Adventurer.X, initialX-1)
	assert.Equal(suite.T(), suite.board.Cells[initialY][0].Adventurer.TreasureCount, 1)
	assert.Equal(suite.T(), suite.board.Cells[1][0].Treasure.Amount, 1)
}

func (suite *BoardTestSuite) TestFullSimulationExample() {
	moves := []enum.PlayerMove{
		enum.Avancer,
		enum.Avancer,
		enum.Droite,
		enum.Avancer,
		enum.Droite,
		enum.Avancer,
		enum.Gauche,
		enum.Gauche,
		enum.Avancer,
	}
	lara := &types.Adventurer{
		Name:          "Lara",
		X:             1,
		Y:             1,
		Orientation:   "S",
		TreasureCount: 0,
		Path:          moves,
	}

	assert.Equal(suite.T(), suite.board.Cells[1][1].Adventurer, lara)
	assert.NotNil(suite.T(), suite.board.Cells[3][0].Treasure)
	assert.NotNil(suite.T(), suite.board.Cells[3][1].Treasure)
	assert.Equal(suite.T(), suite.board.Cells[3][1].Treasure.Amount, 3)

	suite.board.RunSimulation()

	assert.Equal(suite.T(), suite.board.Cells[3][0].Adventurer, &types.Adventurer{
		Name:          "Lara",
		X:             0,
		Y:             3,
		Orientation:   "S",
		TreasureCount: 3,
		Path:          []enum.PlayerMove{},
	})
	assert.Nil(suite.T(), suite.board.Cells[1][1].Adventurer)
	assert.Nil(suite.T(), suite.board.Cells[3][0].Treasure)
	assert.Equal(suite.T(), suite.board.Cells[3][1].Treasure.Amount, 2)
}

type BoardConfigurationTestSuite struct {
	suite.Suite
	board *gameservice.Board
}

func (suite *BoardConfigurationTestSuite) SetupTest() {
	suite.board = InitBoardWithoutConfig(3, 3)
}

func (suite *BoardConfigurationTestSuite) TestPlaceAdventurer() {
	configLine := "A - Lara - 1 - 1 - S - AA"
	err := suite.board.PlaceEntityIntoBoard(configLine)
	suite.Nil(err)
	assert.NotNil(suite.T(), suite.board.Cells[1][1].Adventurer)
}

func (suite *BoardConfigurationTestSuite) TestPlaceMountain() {
	configLine := "M - 1 - 1"
	err := suite.board.PlaceEntityIntoBoard(configLine)
	suite.Nil(err)
	assert.NotNil(suite.T(), suite.board.Cells[1][1].Mountain)
}

func (suite *BoardConfigurationTestSuite) TestPlaceTreasure() {
	configLine := "T - 1 - 1 - 2"
	err := suite.board.PlaceEntityIntoBoard(configLine)
	suite.Nil(err)
	assert.NotNil(suite.T(), suite.board.Cells[1][1].Treasure)
	assert.Equal(suite.T(), suite.board.Cells[1][1].Treasure.Amount, 2)
}

func (suite *BoardConfigurationTestSuite) TestUnknownFlag() {
	configLine := "K - 1 - 1 - 2"
	err := suite.board.PlaceEntityIntoBoard(configLine)
	suite.NotNil(err)
	assert.Equal(suite.T(), err.Error(), errors.Wrap(constants.ErrorUnknownConfigurationFlag, configLine).Error())
}

func (suite *BoardConfigurationTestSuite) TestSetCellContent() {
	err := suite.board.SetCellContent(2, 2, types.NewTreasure(3))
	suite.Nil(err)
	assert.NotNil(suite.T(), suite.board.Cells[2][2].Treasure)
}

func (suite *BoardConfigurationTestSuite) TestSetCellContentOutOfBound() {
	err := suite.board.SetCellContent(3, 3, types.NewTreasure(3))
	suite.NotNil(err)
	assert.Equal(suite.T(), err.Error(), errors.Wrapf(constants.ErrorCoordinateOutOfBounds, "(%d, %d)", 3, 3).Error())
}

type CellCheckersTestSuite struct {
	suite.Suite
	board *gameservice.Board
}

func (suite *CellCheckersTestSuite) SetupSuite() {
	suite.board = InitBoardWithoutConfig(3, 3)
}

func (suite *CellCheckersTestSuite) TestIsCoordinatesInBoard() {
	isInBoard := suite.board.IsCoordinatesInBoard(3, 3)
	assert.False(suite.T(), isInBoard)

	isInBoard = suite.board.IsCoordinatesInBoard(2, 2)
	assert.True(suite.T(), isInBoard)
}

func (suite *CellCheckersTestSuite) TestHasCellMountain() {
	hasMountain := suite.board.HasCellMountain(0, 0)
	assert.False(suite.T(), hasMountain)
	suite.board.Cells[0][0].Mountain = types.NewMountain()
	hasMountain = suite.board.HasCellMountain(0, 0)
	assert.True(suite.T(), hasMountain)
}

func (suite *CellCheckersTestSuite) TestHasCellTreasure() {
	// No treasure
	hasTreasure := suite.board.HasCellTreasure(0, 0)
	assert.False(suite.T(), hasTreasure)

	// Empty Treasure
	suite.board.Cells[0][0].Treasure = types.NewTreasure(0)
	hasTreasure = suite.board.HasCellTreasure(0, 0)
	assert.False(suite.T(), hasTreasure)

	// Treasure with amount
	suite.board.Cells[0][0].Treasure.Amount = 1
	hasTreasure = suite.board.HasCellTreasure(0, 0)
	assert.True(suite.T(), hasTreasure)
}

func (suite *CellCheckersTestSuite) TestHasCellAdventurer() {
	hasAdv := suite.board.HasCellAdventurer(0, 0)
	assert.False(suite.T(), hasAdv)
	suite.board.Cells[0][0].Adventurer = types.NewAdventurer("John", "S", 0, 0, []string{})
	hasAdv = suite.board.HasCellAdventurer(0, 0)
	assert.True(suite.T(), hasAdv)
}

func TestAll(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BoardTestSuite))
	suite.Run(t, new(BoardConfigurationTestSuite))
	suite.Run(t, new(CellCheckersTestSuite))
}

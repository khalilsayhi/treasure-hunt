package types

import (
	"github.com/khalilsayhi/treasure-hunt/domain/constants"
	"github.com/khalilsayhi/treasure-hunt/domain/enum"
	"github.com/pkg/errors"
)

type Adventurer struct {
	Name          string
	Orientation   enum.PlayerOrientation
	X             int
	Y             int
	Path          []enum.PlayerMove
	TreasureCount int
}

// NewAdventurer creates a new Adventurer instance with the given parameters.
func NewAdventurer(name string, orientation enum.PlayerOrientation, x, y int, path []string) *Adventurer {
	return &Adventurer{
		Name:        name,
		Orientation: orientation,
		X:           x,
		Y:           y,
		Path:        enum.ConvertPathListToEnum(path),
	}
}

func (a *Adventurer) isCellContent() {}

func (a *Adventurer) GetNextMove() (enum.PlayerMove, error) {
	if len(a.Path) == 0 {
		return "", errors.Wrap(constants.ErrorAdventurerHasNoPath, a.Name)
	}

	return a.Path[0], nil
}

func (a *Adventurer) Move(x, y int) {
	a.X = x
	a.Y = y
	a.ConsumeMove()
}

func (a *Adventurer) Rotate(orientation enum.PlayerOrientation) {
	a.Orientation = orientation
	a.ConsumeMove()
}

func (a *Adventurer) ConsumeMove() {
	if len(a.Path) > 0 {
		a.Path = a.Path[1:]
	}
}

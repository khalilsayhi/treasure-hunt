package constants

import (
	"errors"
)

// Config errors.
var (
	ErrorBoardConfigLoading        = errors.New("error, failed to load initial board configuration")
	ErrorConfigLandscapeLineFormat = errors.New("error failed to handle landscape configuration line")
	ErrorUnknownConfigurationFlag  = errors.New("error, unknown configuration flag in landscape configuration")
)

// Coordinate errors.
var (
	ErrorCoordinateOutOfBounds = errors.New("error, coordinate out of bounds")
)

// Cell errors.
var (
	ErrorCellContentNotValid = errors.New("error, unknown cell content type")
)

// Adventurer errors.
var (
	ErrorAdventurerHasNoPath = errors.New("error, adventurer %s has no path")
)

// File errors.
var (
	ErrorFileWrite    = errors.New("error writing to file")
	ErrorFileCreation = errors.New("error failed to create file")
)

package utils

import "github.com/khalilsayhi/treasure-hunt/domain/enum"

func IsRotation(move enum.PlayerMove) bool {
	return move == enum.Droite || move == enum.Gauche
}

func GetNextCoordinates(x, y int, orientation enum.PlayerOrientation) (int, int) {
	switch orientation {
	case enum.Nord:
		return x, y - 1
	case enum.Est:
		return x + 1, y
	case enum.Sud:
		return x, y + 1
	case enum.Ouest:
		return x - 1, y
	default:
		return x, y
	}
}

func GetNextOrientation(currentOrientation enum.PlayerOrientation, rotation enum.PlayerMove) enum.PlayerOrientation {
	switch currentOrientation {
	case enum.Nord:
		switch rotation {
		case enum.Droite:
			return enum.Est
		case enum.Gauche:
			return enum.Ouest
		}
	case enum.Est:
		switch rotation {
		case enum.Droite:
			return enum.Sud
		case enum.Gauche:
			return enum.Nord
		}
	case enum.Sud:
		switch rotation {
		case enum.Droite:
			return enum.Ouest
		case enum.Gauche:
			return enum.Est
		}
	case enum.Ouest:
		switch rotation {
		case enum.Droite:
			return enum.Nord
		case enum.Gauche:
			return enum.Sud
		}
	}

	return currentOrientation
}

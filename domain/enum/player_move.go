package enum

type PlayerMove string

const (
	Gauche  PlayerMove = "G"
	Droite  PlayerMove = "D"
	Avancer PlayerMove = "A"
)

func (pm PlayerMove) ToEnum() PlayerMove {
	switch pm {
	case Gauche:
		return Gauche
	case Droite:
		return Droite
	case Avancer:
		return Avancer
	default:
		return ""
	}
}

func ConvertPathListToEnum(path []string) []PlayerMove {
	enumPath := make([]PlayerMove, len(path))
	for i, move := range path {
		enumPath[i] = PlayerMove(move).ToEnum()
	}
	
	return enumPath
}

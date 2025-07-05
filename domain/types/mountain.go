package types

type Mountain struct {
}

// NewMountain creates a new Mountain instance.
func NewMountain() *Mountain {
	return &Mountain{}
}

func (m Mountain) isCellContent() {}

package types

type Treasure struct {
	Amount int
}

// NewTreasure creates a new Treasure instance with the specified amount.
func NewTreasure(amount int) *Treasure {
	return &Treasure{Amount: amount}
}

func (t *Treasure) isCellContent() {}

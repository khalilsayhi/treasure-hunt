package types

type Treasure struct {
	Amount int
}

func NewTreasure(amount int) *Treasure {
	return &Treasure{Amount: amount}
}

func (t *Treasure) isCellContent() {}

package types

type CellContent interface {
	isCellContent()
}

type Cell struct {
	Mountain   *Mountain
	Adventurer *Adventurer
	Treasure   *Treasure
}

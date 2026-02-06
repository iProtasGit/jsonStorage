package bins

import (
	"time"
)

type BinList *[]Bin

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, private bool, name string, binList BinList) {
	bin := Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	*binList = append(*binList, bin)
}

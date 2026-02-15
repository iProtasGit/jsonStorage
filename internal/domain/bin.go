package domain

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, private bool, name string) *Bin {
	return  &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func (b *Bin) Validate() error {
	if b.ID == "" {
		return fmt.Errorf(BIN_ID_EMPTY)
	}
	if b.Name == "" {
		return fmt.Errorf(BIN_NAME_EMPTY)
	}
	return nil
}

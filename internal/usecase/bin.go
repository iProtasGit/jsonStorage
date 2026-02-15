package usecase

import (
	"jsonStorage/internal/domain"
	"jsonStorage/internal/adapter"
)

type UseCase struct {
	storage *adapter.Storage
}

func NewUseCase(storage *adapter.Storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}

func (u *UseCase) CreateBin(id string, private bool, name string) (*domain.Bin, error) {
	bin := domain.NewBin(id, private, name)
	err := bin.Validate()
	if err != nil {
		return nil, err
	}

	u.storage.Bins = append(u.storage.Bins, *bin)
	err = u.storage.SaveStorage()
	if err != nil {
		return nil, err
	}

	return bin, nil
}

func (u *UseCase) GetBins() ([]domain.Bin, error) {
	return u.storage.Bins, nil
}
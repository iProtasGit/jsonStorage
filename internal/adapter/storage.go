package adapter

import (
	"encoding/json"
	"time"

	binsRepo "jsonStorage/internal/domain"
	"jsonStorage/pkg/files"
)

type Storage struct {
	Name string `json:"name"`
	Bins []binsRepo.Bin `json:"bins"`
	UpdateTime time.Time `json:"update_time"`
}

func NewStorage(name string, bins []binsRepo.Bin) *Storage {
	return &Storage{
		Name: name,
		Bins: bins,
		UpdateTime: time.Now(),
	}
}

func (s *Storage) SaveStorage() error {
	s.UpdateTime = time.Now()
	
	data, err := json.Marshal(&s)
	if err != nil {
		return err
	}
	
	err = files.WriteFile(data, s.Name)
	if err != nil {
		return err
	}
	
	return nil
}

func GetStorage(name string) (*Storage, error) {
	if !files.CheckFileExists(name) {
		files.CreateFile(name)
		return NewStorage(name, []binsRepo.Bin{}), nil
	}
	
	data, err := files.ReadFile(name)
	if err != nil {
		return nil, err
	}
	
	var storage Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		return nil, err
	}
	
	return &storage, nil
}

func (s *Storage) UpdateStorage(bin binsRepo.Bin) error {
	s.Bins = append(s.Bins, bin)
	
	err := s.SaveStorage()
	if err != nil {
		return err
	}
	return nil
}
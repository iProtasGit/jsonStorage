package config

import (
	"encoding/json"
	"jsonStorage/pkg/files"
)

type Config struct {
	StorageName string `json:"storage_name"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := files.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	var config Config
	err = json.Unmarshal(data, &config)
	
	if err != nil {
		return nil, err
	}
	return &config, nil
}
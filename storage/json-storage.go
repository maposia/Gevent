package storage

import (
	"os"
)

type JsonStorage struct {
	*Storage
}

func NewJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{&Storage{filename: filename}}
}

func (s *JsonStorage) Save(data []byte) error {
	err := os.WriteFile(s.getFileName(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *JsonStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(s.getFileName())
	if err != nil {
		return nil, err
	}
	return data, nil
}

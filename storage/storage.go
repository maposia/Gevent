package storage

type Store interface {
	Save(data []byte) error
	Load() ([]byte, error)
	getFileName() string
}

type Storage struct {
	filename string
}

func (s *Storage) getFileName() string {
	return s.filename
}

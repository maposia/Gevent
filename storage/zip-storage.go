package storage

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	return &ZipStorage{&Storage{filename}}
}

func (z *ZipStorage) Save(data []byte) error {
	f, err := os.Create(z.getFileName())
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	w, errZip := zw.Create("data.json")
	if errZip != nil {
		return errZip
	}

	_, err = w.Write(data)
	return err
}

func (z *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(z.getFileName())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if len(r.File) == 0 {
		return nil, errors.New("archive is empty")
	}

	file := r.File[0]
	rc, errOpen := file.Open()
	if errOpen != nil {
		return nil, errors.New("error open file")
	}

	defer rc.Close()

	return io.ReadAll(rc)
}

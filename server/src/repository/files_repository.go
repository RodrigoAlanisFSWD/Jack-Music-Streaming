package repository

import (
	"io"
	"os"
	"path/filepath"

	"github.com/Jack-Music-Streaming/server/src/models"
)

type filesRepository struct {
}

func NewFilesRepository() models.FilesRepository {
	return &filesRepository{}
}

func (filesRepository) Upload(file *models.File) error {
	src, err := file.Src.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	if _, err := os.Stat(file.Dst); err != nil {
		if err := os.Mkdir(file.Dst, 0700); err != nil {
			return err
		}
	}

	// Destination
	dst, err := os.Create(file.Dst + file.Name)

	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func (filesRepository) CreateFileName(file *models.File) {
	ext := filepath.Ext(file.Src.Filename)

	file.Name = file.Name + ext
}

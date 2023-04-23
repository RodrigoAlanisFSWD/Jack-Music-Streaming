package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type filesRepository struct {
	DB *gorm.DB
}

func NewFilesRepository(db *gorm.DB) models.FilesRepository {
	return &filesRepository{
		DB: db,
	}
}

func (f filesRepository) Upload(file *models.FileDTO) (*models.File, error) {
	src, err := file.Src.Open()

	fileModel := &models.File{
		Name: file.Name,
		Src:  file.Dst + file.Name,
		Ext:  filepath.Ext(file.Src.Filename),
	}

	if err != nil {
		return fileModel, err
	}

	defer src.Close()

	if _, err := os.Stat(file.Dst); err != nil {
		if err := os.Mkdir(file.Dst, 0700); err != nil {
			return fileModel, err
		}
	}

	// Destination
	dst, err := os.Create(file.Dst + file.Name)

	if err != nil {
		return fileModel, err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return fileModel, err
	}

	err = f.DB.Create(fileModel).Error

	if err != nil {
		return fileModel, err
	}

	return fileModel, nil
}

func (filesRepository) CreateFileName(file *models.FileDTO) {
	fmt.Println("asd")
	ext := filepath.Ext(file.Src.Filename)

	file.Name = file.Name + ext
}

func (f filesRepository) Delete(file *models.File) error {
	err := os.Remove(file.Src)

	if err != nil {
		return err
	}

	err = f.DB.Where("id = ?", file.ID).Delete(file).Error

	return err
}

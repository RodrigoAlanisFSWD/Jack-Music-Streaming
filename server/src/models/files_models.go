package models

import "mime/multipart"

type FileDTO struct {
	Name string
	Dst  string
	Src  *multipart.FileHeader
}

type File struct {
	Model
	Src  string `json:"src"`
	Ext  string `json:"ext"`
	Name string `json:"name"`
}

type FilesRepository interface {
	Upload(file *FileDTO) (*File, error)
	CreateFileName(file *FileDTO)
	Delete(file *File) error
}

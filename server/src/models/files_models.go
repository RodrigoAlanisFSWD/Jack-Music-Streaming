package models

import "mime/multipart"

type File struct {
	Name string
	Dst  string
	Src  *multipart.FileHeader
}

type FilesRepository interface {
	Upload(file *File) error
	CreateFileName(file *File)
}

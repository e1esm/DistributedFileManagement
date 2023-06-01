package service

type Service interface {
	SaveParts([]byte)
}

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

package service

import "DistributedFileManagement/server/files/repository"

type Service interface {
	SaveParts([]byte)
}

type FileService struct {
	Repositories *repository.Repositories
}

func NewFileService(repositories *repository.Repositories) *FileService {
	return &FileService{Repositories: repositories}
}

func (f *FileService) SaveParts([]byte) {

}

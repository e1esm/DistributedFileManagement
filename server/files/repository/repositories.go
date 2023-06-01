package repository

type Repositories struct {
	fileRepo   *FileRepository
	dockerRepo *DockerRepository
}

func NewRepositories() *Repositories {
	return &Repositories{dockerRepo: NewDockerRepository(), fileRepo: NewFileRepository()}
}

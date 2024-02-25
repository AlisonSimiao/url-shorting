package photo

import "time-wise/repository"

type PhotoRepository struct {
	repository.Repository
}

func NewPhotoRepository() *repository.Repository {
	pr := &PhotoRepository{}

	pr.Repository.Super("photos")
	return &pr.Repository
}


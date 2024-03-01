package click

import "url-shorting/repository"

type ClickRepository struct {
	repository.Repository
}

func NewClickRepository() *repository.Repository {
	cr := &ClickRepository{}

	cr.Repository.Super("clicks")
	return &cr.Repository
}

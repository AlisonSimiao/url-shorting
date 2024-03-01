package link

import "url-shorting/repository"

type LinkRepository struct {
	repository.Repository
}

func NewLinkRepository() *repository.Repository {
	lr := &LinkRepository{}

	lr.Repository.Super("links")
	return &lr.Repository
}
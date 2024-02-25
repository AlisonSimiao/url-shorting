package project

import "time-wise/repository"

type ProjectRepository struct {
	repository.Repository
}

func NewProjectRepository() *repository.Repository {
	pr := &ProjectRepository{}

	pr.Repository.Super("projects")
	return &pr.Repository
}

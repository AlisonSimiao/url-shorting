package user

import "url-shorting/repository"

type UserRepository struct {
	repository.Repository
}

func NewUserRepository() *repository.Repository {
	ur := &UserRepository{}

	ur.Repository.Super("users")
	return &ur.Repository
}

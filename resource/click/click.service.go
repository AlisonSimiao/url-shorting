package click

import "url-shorting/repository"

type ClickService struct {
	cr *repository.Repository
}

func NewClickService() *ClickService {
	return &ClickService{
		cr: NewClickRepository(),
	}
}

func (cs *ClickService) Create() Click {
	click := Click{
		Value: 0,
	}

	cs.cr.Create(&click)
	
	return click
}

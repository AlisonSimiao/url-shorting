package click

import (
	"url-shorting/repository"
	rest_error "url-shorting/restError"

	"gorm.io/gorm"
)

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

func (cs *ClickService) AddClick(id int64) *rest_error.Err {
	result := cs.cr.Raw().Where("id = ?", id).Update("value", gorm.Expr("value + ?", 1))
	if result.Error != nil {
		return rest_error.NewInternalError()
	}

	return nil
}

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

type object map[string]interface{}

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
func (cs *ClickService) FindOne(id int64) (Click, *rest_error.Err) {

	var click Click
	cs.cr.FindOne("id = @id", object{"id": id}, &click)
	if click.Id == 0 {
		return click, rest_error.NewInternalError()
	}

	return click, nil 
}

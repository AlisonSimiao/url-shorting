package link

import (	
	rest_error "url-shorting/restError"	
	"github.com/gin-gonic/gin"	
)

type LinkService struct {
	// ur *repository.Repository	
}

func NewLinkService() *LinkService {
	return &LinkService{
		// ur: NewLinkRepository(),		
	}
}

type object map[string]interface{}

func (u *LinkService) update(id int, c *gin.Context) *rest_error.Err {
	
	return nil
}

func (u *LinkService) create( c *gin.Context) ( *rest_error.Err) {
	
	return nil
}

func (u *LinkService) findOne(id int) (*rest_error.Err) {
	
	return nil
}

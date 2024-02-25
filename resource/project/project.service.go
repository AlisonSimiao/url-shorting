package project

import (
	"mime/multipart"
	"time-wise/repository"
	"time-wise/resource/photo"
	rest_error "time-wise/restError"

	"github.com/gin-gonic/gin"
)

type ProjectService struct {
	pr *repository.Repository
	phr *repository.Repository
	phs *photo.PhotoService
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		pr: NewProjectRepository(),
		phr: photo.NewPhotoRepository(),
		phs: photo.NewPhotoService(),
	}
}

func (ps * ProjectService) paginate(idUser int, pageInfo *repository.PageInfo) (*repository.PaginateData, *rest_error.Err){
	var projects []ProjectResponse
	
	result := ps.pr.PaginateWithJoin(`
	projects.id,
	projects.name,        
    projects.description, 
    projects.status,      
    photos.url`, 
	"left join photos on projects.id_photo = photos.id", 
	"id_user = @idUser", 
	repository.Object{"idUser": idUser}, 
	&projects, 
	pageInfo.Page, pageInfo.Limit)
	
	if result == nil {
		return nil, rest_error.NewNotFoundError("Projetos não encontrados")
	}

	return result, nil
}

func (ps * ProjectService) create(project Project, c *gin.Context) (*ProjectResponse, *rest_error.Err) {
	urlPhoto := ps.phs.GetDefaultPhoto()
	
	if file, exist := c.Get("files"); exist {
		project.IdPhoto, urlPhoto = ps.phs.SavePhoto(c, file.(map[string]*multipart.FileHeader))
	}

	ps.pr.Create(&project)
	if project.Id == 0 {
		return nil, rest_error.NewInternalError()
	}

	return &ProjectResponse{
		Id: project.Id,
		Name: project.Name,
		Description: project.Description,
		Status: project.Status,
		Url: urlPhoto,
	}, nil
}
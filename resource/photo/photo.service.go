package photo

import (
	"mime/multipart"
	"url-shorting/storage"

	"github.com/gin-gonic/gin"
)

type PhotoService struct {
}

var defaultURL = "https://res.cloudinary.com/dbwxgur7d/image/upload/v1690331509/projects/5b10e54d-2c22-43b7-9a12-552274ecb686.gif"

func NewPhotoService() *PhotoService {
	ps := &PhotoService{}

	return ps
}

func (ps *PhotoService) GetDefaultPhoto() string {
	return defaultURL
}

func (ps *PhotoService) SavePhoto(c *gin.Context, file map[string]*multipart.FileHeader) (int, string) {
	idPhoto := 1
	cloud := storage.New(c)

	src, err := file["file"].Open()
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return 1, defaultURL
	}

	err = cloud.UploadImage(src)
	if err != nil {
		return 1, defaultURL
	}

	createdPhoto := Photo{
		Url:       cloud.GetUrl(),
		Public_id: cloud.GetPublicId(),
		Path:      cloud.GetPath(),
		Type:      1,
	}

	NewPhotoRepository().Create(&createdPhoto)
	if createdPhoto.Id != 0 {
		idPhoto = createdPhoto.Id
	}

	return idPhoto, cloud.GetUrl()
}

package storage

// Import Cloudinary anmaind other necessary libraries
//===================
import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CloudinaryStorage struct {
	cld            *cloudinary.Cloudinary
	ctx            context.Context
	folder         string
	uniqueFilename bool
	overwrite      bool
	publicID       string
	url            string
}

var storage *cloudinary.Cloudinary = nil

func (cs *CloudinaryStorage) GetUrl() string {
	return cs.url
}

func (cs *CloudinaryStorage) GetPath() string {
	return cs.folder
}

func (cs *CloudinaryStorage) GetPublicId() string {
	return cs.publicID
}

func New(ctx *gin.Context) *CloudinaryStorage {
	if storage == nil {
		storage = credentials()
	}

	return &CloudinaryStorage{
		cld:            storage,
		ctx:            ctx,
		folder:         "/IMG",
		uniqueFilename: true,
		overwrite:      true,
		publicID:       uuid.New().String(),
		url:            "",
	}
}

func NewWithParams(ctx *gin.Context, folder string, publicID string) *CloudinaryStorage {
	if storage == nil {
		storage = credentials()
	}

	return &CloudinaryStorage{
		cld:            storage,
		ctx:            ctx,
		folder:         folder,
		uniqueFilename: true,
		overwrite:      true,
		publicID:       publicID,
	}
}

func credentials() *cloudinary.Cloudinary {
	CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")

	cld, _ := cloudinary.NewFromURL(CLOUDINARY_URL)
	cld.Config.URL.Secure = true
	return cld
}

func (cs *CloudinaryStorage) UploadImage(src interface{}) error {
	resp, err := cs.cld.Upload.Upload(cs.ctx, src, uploader.UploadParams{
		PublicID:       cs.publicID,
		UniqueFilename: api.Bool(cs.uniqueFilename),
		Overwrite:      api.Bool(cs.overwrite),
		Folder:         cs.folder,
	})

	if err != nil {
		return err
	}

	cs.url = resp.SecureURL

	return nil
}

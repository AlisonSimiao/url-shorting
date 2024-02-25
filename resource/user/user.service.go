package user

import (
	"mime/multipart"
	"time"
	"url-shorting/repository"
	"url-shorting/resource/photo"
	rest_error "url-shorting/restError"
	"url-shorting/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur *repository.Repository
	pr *repository.Repository
}

func NewUserService() *UserService {
	return &UserService{
		ur: NewUserRepository(),
		pr: photo.NewPhotoRepository(),
	}
}

func hashPassword(password string) string {
	bcrypt, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bcrypt)
}

type object map[string]interface{}

func (u *UserService) update(id int, body User, c *gin.Context) *rest_error.Err {
	var user UserResponse

	u.ur.FindOne("id = @id", object{"id": id}, &user)
	if user.Id == 0 {
		return rest_error.NewNotFoundError("usuario não encontrado")
	}

	if body.Email != "" && body.Email != user.Email {
		var user UserResponse
		u.ur.FindOne("email = @email and id != @id", object{"email": body.Email, "id": id}, &user)
		if user.Id != 0 {
			return rest_error.NewConflictError("Já existe um usuario com esse email")
		}
	}

	if body.Password != "" {
		body.Password = hashPassword(body.Password)
	}

	if file, exist := c.Get("files"); exist {
		body.IdPhoto, _ = photoService.SavePhoto(c, file.(map[string]*multipart.FileHeader))
	}

	u.ur.Update("id = @id", object{"id": id}, User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		IdPhoto:  body.IdPhoto,
	})

	return nil
}

func (u *UserService) create(body User, c *gin.Context) (UserResponse, *rest_error.Err) {
	var user UserResponse
	urlPhoto := photoService.GetDefaultPhoto()

	u.ur.FindOne("email = @email", object{"email": body.Email}, &user)
	if user.Id != 0 {
		return UserResponse{}, rest_error.NewConflictError("Já existe um usuario com esse email")
	}

	body.Password = hashPassword(body.Password)

	if file, exist := c.Get("files"); exist {
		body.IdPhoto, urlPhoto = photoService.SavePhoto(c, file.(map[string]*multipart.FileHeader))
	}

	u.ur.Create(&body)
	if body.Id == 0 {
		return UserResponse{}, rest_error.NewInternalError()
	}

	return UserResponse{
		Id:    body.Id,
		Name:  body.Name,
		Email: body.Email,
		Url:   urlPhoto,
	}, nil
}

func (u *UserService) login(body UserLogin) (LoginResponse, *rest_error.Err) {
	values := make(object)
	var user User

	values["email"] = body.Email
	u.ur.FindOne("email = @email", values, &user)

	if user.Email == "" {
		return LoginResponse{}, rest_error.NewUnauthorizedError("email ou senha incorretos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return LoginResponse{}, rest_error.NewUnauthorizedError("email ou senha incorretos")
	}

	t, errToken := token.CreateToken(user.Id, time.Hour*24)
	if errToken != nil {
		return LoginResponse{}, rest_error.NewInternalError()
	}

	var userPhoto photo.Photo
	photo.NewPhotoRepository().FindOne("id = @id", object{"id": user.IdPhoto}, &userPhoto)

	return LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: t,
		Url:   userPhoto.Url,
	}, nil
}

func (u *UserService) findOne(id int) (UserResponse, *rest_error.Err) {
	var user UserResponse
	u.ur.FindOneWithJoin("users.id, users.name, users.email, photos.url", "join photos on users.id_photo = photos.id", "users.id = @id", object{"id": id}, &user)

	if user.Id == 0 {
		return UserResponse{}, rest_error.NewNotFoundError("usuario não encontrado")
	}

	return user, nil
}

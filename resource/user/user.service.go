package user

import (
	//"mime/multipart"
	"time"
	"vagas-api/repository"
	"vagas-api/resource/photo"
	rest_error "vagas-api/restError"
	"vagas-api/token"

	"golang.org/x/crypto/bcrypt"
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

func (u *UserService) update(id int, body User) *rest_error.Err {
	var user UserResponse

	u.ur.FindOne("id = @id", object{"id": id}, &user)
	if user.Id == 0 {
		return rest_error.NewNotFoundError("usuário não encontrado")
	}

	if body.Email != "" && body.Email != user.Email {
		var user UserResponse
		u.ur.FindOne("email = @email and id != @id", object{"email": body.Email, "id": id}, &user)
		if user.Id != 0 {
			return rest_error.NewConflictError("Já existe um usuário com esse email")
		}
	}

	if body.Password != "" {
		body.Password = hashPassword(body.Password)
	}

	/* if file, exist := c.Get("files"); exist {
		body.IdPhoto, _ = photoService.SavePhoto(c, file.(map[string]*multipart.FileHeader))
	}
	*/
	u.ur.Update("id = @id", object{"id": id}, User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Username: body.Username,
		Status:   body.Status,
		Pro:      body.Pro,
		//IdPhoto:  body.IdPhoto,
	})

	return nil
}

	// @Summary cria novo usuario
	// @Description usuarios
	// @Accept  json
	// @Produce  json
	// @Success 201 {object} UserResponse
	// @Failure 400 {object} rest_error.Err "We need ID!!"
	// @Failure 404 {object} rest_error.Err "Can not find ID"
	// @Router /users [post]
func (u *UserService) create(body User) (UserResponse, *rest_error.Err) {
	var user UserResponse
	//urlPhoto := photoService.GetDefaultPhoto()

	u.ur.FindOne("username = @username", object{"username": body.Username}, &user)
	if user.Id != 0 {
		return UserResponse{}, rest_error.NewConflictError("Já existe um usuário com esse username")
	}

	u.ur.FindOne("email = @email", object{"email": body.Email}, &user)
	if user.Id != 0 {
		return UserResponse{}, rest_error.NewConflictError("Já existe um usuário com esse email")
	}

	body.Password = hashPassword(body.Password)

	/* if file, exist := c.Get("files"); exist {
		body.IdPhoto, urlPhoto = photoService.SavePhoto(c, file.(map[string]*multipart.FileHeader))
	} */

	u.ur.Create(&body)
	if body.Id == 0 {
		return UserResponse{}, rest_error.NewInternalError()
	}

	return UserResponse{
		Id:       body.Id,
		Name:     body.Name,
		Email:    body.Email,
		Username: body.Username,
		Status:   body.Status,
		Pro:      body.Pro,
		//Url:   urlPhoto,
	}, nil
}

// @tags User
// @Summary Get a list of books in the the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 "ok"
// @Router /books [get]
func (u *UserService) login(body UserLogin) (LoginResponse, *rest_error.Err) {
	values := make(object)
	var user User

	values["username"] = body.Username
	u.ur.FindOne("username = @username", values, &user)

	if user.Username == "" {
		return LoginResponse{}, rest_error.NewUnauthorizedError("username ou senha incorretos1")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return LoginResponse{}, rest_error.NewUnauthorizedError("username ou senha incorretos2")
	}

	t, errToken := token.CreateToken(user.Id, time.Hour*24)
	if errToken != nil {
		return LoginResponse{}, rest_error.NewInternalError()
	}

	//var userPhoto photo.Photo
	//photo.NewPhotoRepository().FindOne("id = @id", object{"id": user.IdPhoto}, &userPhoto)

	return LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: t,
		//Url:   userPhoto.Url,
	}, nil
}

func (u *UserService) findOne(id int) (UserResponse, *rest_error.Err) {
	var user UserResponse
	u.ur.FindOne("users.id = @id", object{"id": id}, &user)

	if user.Id == 0 {
		return UserResponse{}, rest_error.NewNotFoundError("usuário não encontrado")
	}

	return user, nil
}

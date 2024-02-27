package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Status   bool `json:"status"`
	Pro 	 bool `json:"pro"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Status   bool `json:"status"`
	Pro 	 bool `json:"pro"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

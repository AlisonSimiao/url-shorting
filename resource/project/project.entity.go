package project

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	IdPhoto     int    `json:"id_photo"`
	IdUser      int    `json:"id_user"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type ProjectResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Url   		string `json:"url"`
}

package repository

type PageInfo struct {
	Limit int
	Page  int
}

type Repository struct {
	Table string
}

type PaginateData struct {
	TotalEntities int
	CurrentPage   int
	TotalPages    int
	Data          any
}

type Object map[string]interface{}
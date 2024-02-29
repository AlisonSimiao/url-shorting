package rest_error

import "net/http"

type NotFoundError struct {
	Err
}

func NewNotFoundError(Mensagem string) *Err {
	return &Err{
		Mensagem: Mensagem,
		status:   http.StatusNotFound,
	}
}

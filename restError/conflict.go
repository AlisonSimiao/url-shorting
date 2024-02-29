package rest_error

import "net/http"

type ConflictError struct {
	Err
}

func NewConflictError(Mensagem string) *Err {
	return &Err{
		Mensagem: Mensagem,
		status:   http.StatusConflict,
	}
}

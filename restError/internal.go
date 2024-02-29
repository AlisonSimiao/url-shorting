package rest_error

import "net/http"

type InternalError struct {
	Err
}

func NewInternalError() *Err {
	return &Err{
		Mensagem: "Erro interno no servidor",
		status:   http.StatusInternalServerError,
	}
}

package rest_error

import "net/http"

type UnauthorizedError struct {
	Err
}

func NewUnauthorizedError(Mensagem string) *Err {
	return &Err{
		Mensagem: Mensagem,
		status:   http.StatusUnauthorized,
	}
}

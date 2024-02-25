package rest_error

type UnauthorizedError struct {
	Err
}

func NewUnauthorizedError(Mensagem string) *Err {
	return &Err{
		Mensagem: Mensagem,
		status:   401,
	}
}

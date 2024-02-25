package rest_error

type NotFoundError struct {
	Err
}

func NewNotFoundError(Mensagem string) *Err {
	return &Err{
		Mensagem: Mensagem,
		status:   404,
	}
}

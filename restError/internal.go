package rest_error

type InternalError struct {
	Err
}

func NewInternalError() *Err {
	return &Err{
		Mensagem: "Erro interno no servidor",
		status:   404,
	}
}

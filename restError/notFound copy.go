package rest_error

type ConflictError struct {
	Err
}

func NewConflictError(Mensagem string) *Err {
	return &Err{
		Mensagem: Mensagem,
		status:   409,
	}
}

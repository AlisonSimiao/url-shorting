package rest_error

type Err struct {
	Mensagem string
	status   int
}

func (e *Err) Error() string {
	return e.Mensagem
}

func (e *Err) JsonError() map[string]string {
	return map[string]string{
		"mensagem": e.Mensagem,
	}
}

func (e *Err) GetMensagem() string {
	return e.Mensagem
}
func (e *Err) GetStatus() int {
	return e.status
}

package utils

import "fmt"

func Alpha(field string) string {
	return fmt.Sprintf("O campo %s deve conter apenas letras", field)
}

func AlphaNum(field string) string {
	return fmt.Sprintf("O campo %s deve conter apenas letras e números", field)
}

func Required(field string) string {
	return fmt.Sprintf("O campo %s é obrigatário", field)
}

func Invalid(field string) string {
	return fmt.Sprintf("O campo %s é inválido", field)
}

func MaxLength(field string, length string) string {
	return fmt.Sprintf("O campo %s deve ter no máximo %s caracteres", field, length)
}

func MinLength(field string, length string) string {
	return fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres", field, length)
}

func Email() string {
	return Invalid("email")
}

func IsIn(field string, values ...string) string {
	return fmt.Sprintf("O campo %s deve ser um dos seguintes valores: %s", field, values)
}

func Numeric(field string) string {
	return fmt.Sprintf("O campo %s deve ser numero", field)
}

func Equal(field string, value string) string {
	return fmt.Sprintf("O campo %s deve ser igual a %s", field, value)
}

func Between(field string, min string, max string) string {
	return fmt.Sprintf("O campo %s deve estar entre %s e %s", field, min, max)
}

func Boolean(field string) string {
	return fmt.Sprintf("O campo %s deve ser 'true' ou 'false'", field)
}

func Url(field string) string {
	return fmt.Sprintf("O campo %s deve ser uma url válida", field)
}

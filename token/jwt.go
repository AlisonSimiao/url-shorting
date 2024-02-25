package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = "SECRET_KEY"

type Payload struct {
	IdUser int
	Exp    int64
	Iat    int64
}

func ExtractJWTToken(authHeader string) (token, error string) {
	if authHeader == "" {
		return "", "Cabeçalho de autenticação ausente"
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", "Cabeçalho de autenticação mal formatado. Deve começar com 'Bearer '"
	}

	token = strings.TrimPrefix(authHeader, "Bearer ")
	return token, ""
}

func CreateToken(id int, duration time.Duration) (string, error) {
	sk := os.Getenv(SECRET_KEY)
	payload := jwt.MapClaims{
		"idUser": id,
		"exp":    time.Now().Add(duration).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(sk))
}

func VerifyToken(tokenString string) (*Payload, error) {
	// Obtém a chave secreta do ambiente
	secretKey := os.Getenv(SECRET_KEY)
	// Faz o parsing do token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodo de assinatura inválido: %v", token.Header["alg"])
		}
		// Retorna a chave secreta para validar a assinatura
		return []byte(secretKey), nil
	})
	// Verifica se houve algum erro durante o parsing do token
	if err != nil {
		return nil, err
	}

	// Verifica se o token é válido
	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	// Converte o payload do token para a estrutura Payload
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := int(claims["idUser"].(float64))

		if !ok {
			return nil, fmt.Errorf("erro ao extrair payload do token")
		}
		payload := &Payload{
			IdUser: id,
		}
		return payload, nil
	}

	return nil, fmt.Errorf("erro ao extrair payload do token")
}

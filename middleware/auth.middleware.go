package middleware

import (
	"time-wise/token"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	t, errToken := token.ExtractJWTToken(c.GetHeader("Authorization"))
	if errToken != "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Falha na autenticação", "message": errToken})
		return
	}

	payload, err := token.VerifyToken(t)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	c.Set("idUser", payload.IdUser)
	c.Next()
}

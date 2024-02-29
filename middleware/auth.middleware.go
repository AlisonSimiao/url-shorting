package middleware

import (
	"url-shorting/token"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	auth, _ := c.Cookie("auth_token")
	
	t, errToken := token.ExtractJWTToken(auth)
	if errToken != "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Falha na autenticação", "message": errToken})
		return
	}
	
	payload, err := token.VerifyToken(t)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Falha na autenticação", "message": "Unauthorized"})
		return
	}

	c.Set("idUser", payload.IdUser)
	c.Next()
}

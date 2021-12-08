package middleware

import (
	"fmt"
	"go-gin-api/authentication"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
				
			tokenString := authHeader[len(BEARER_SCHEMA):]
			
			token, err := authentication.JWTAuthService().ValidateToken(tokenString)
			if err!= nil {
				println(err)
				c.JSON(http.StatusUnauthorized,gin.H{"err_msg":err.Error()})
			}
			if token.Valid {
				
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
			} else {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
	}else {
		
		c.AbortWithStatus(http.StatusUnauthorized)
	}
		


	}
}
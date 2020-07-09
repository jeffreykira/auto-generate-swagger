package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jeffreykira/log-management/core/api/model"
	"github.com/jeffreykira/log-management/core/api/response"
)

//AgentToken ...
const AgentToken = "access-token"

//JwtSecret ...
var JwtSecret = []byte(AgentToken)

//ValidateToken ...
func ValidateToken(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	authToken := strings.Split(auth, "Bearer ")
	var token string
	if len(authToken) == 2 {
		token = authToken[1]
	} else {
		c.JSON(http.StatusUnauthorized, response.HTTPError{
			Code:    401,
			Message: "token unauthorized",
		})
		c.Abort()
		return
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtSecret, nil
	})
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token unauthorized"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, response.HTTPError{
			Code:    401,
			Message: message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*model.Claims); ok {
		if claims.Role == "" {
			c.JSON(http.StatusUnauthorized, response.HTTPError{
				Code:    401,
				Message: "JWT token payload is improper",
			})
			c.Abort()
		} else {
			c.Set("role", claims.Role)
			c.Next()
		}
	} else {
		c.JSON(http.StatusUnauthorized, response.HTTPError{
			Code:    401,
			Message: "JWT token payload is improper",
		})
		c.Abort()
	}
}

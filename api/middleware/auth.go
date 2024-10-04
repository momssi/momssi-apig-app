package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"momssi-apig-app/api/form"
	"momssi-apig-app/internal/domain/member"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization 헤더에서 JWT 토큰 추출
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			failResponse(c, http.StatusUnauthorized, form.ErrMissingToken, form.GetCustomErr(form.ErrMissingToken))
			c.Abort()
			return
		}

		// JWT 토큰 파싱 및 검증
		claims := &member.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return member.JWTKey, nil
		})

		if err != nil || !token.Valid {
			failResponse(c, http.StatusUnauthorized, form.ErrInvalidToken, form.GetCustomErr(form.ErrInvalidToken))
			c.Abort()
			return
		}

		// 토큰이 유효한 경우, 다음 핸들러로 진행
		c.Set("username", claims.Username)
		c.Next()
	}
}

func failResponse(c *gin.Context, statusCode int, errorCode int, err error) {

	logMessage := form.GetCustomErrMessage(errorCode, err.Error())
	c.Errors = append(c.Errors, &gin.Error{
		Err:  fmt.Errorf(logMessage),
		Type: gin.ErrorTypePrivate,
	})

	c.JSON(statusCode, form.ApiResponse{
		ErrorCode: errorCode,
		Message:   form.GetCustomMessage(errorCode),
	})
}

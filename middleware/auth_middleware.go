package middleware

import (
	"net/http"
	"strings"

	"github.com/GeZaM8/laundry-be/auth"
	"github.com/GeZaM8/laundry-be/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        h := c.GetHeader("Authorization")

        if !strings.HasPrefix(h, "Bearer ") {
            c.JSON(http.StatusUnauthorized, model.Response{
                Status:  false,
                Message: "Token dibutuhkan",
            })
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(h, "Bearer ")

        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return auth.JWT_SECRET, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, model.Response{
				Status:  false,
				Message: "Token tidak valid",
			})
            c.Abort()
            return
        }

        c.Next()
    }
}

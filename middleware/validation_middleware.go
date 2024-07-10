package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Validator() gin.HandlerFunc {
	v := validator.New()

	return func(c *gin.Context) {
		if err := v.Struct(c.Request.Body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

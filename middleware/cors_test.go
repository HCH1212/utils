package middleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test(t *testing.T) {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.Use(CorsGin())

	router.Run(":8080")
}

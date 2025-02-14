package resp

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test(t *testing.T) {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.GET("/", func(c *gin.Context) {
		Success(c, "成功", nil)
	})

	router.Run(":8080")
}

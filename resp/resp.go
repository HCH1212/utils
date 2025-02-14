package resp

// 实际使用根据所选框架进行调整，这里使用gin

import "github.com/gin-gonic/gin"

const (
	OK         = 2000
	NO         = 4000
	ServerFail = 5000
)

// ResponseWithStatusAndData 确定统一返回格式
func ResponseWithStatusAndData(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// 成功的返回
func Success(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndData(c, OK, message, data)
}

// 客户端请求失败的返回
func Fail(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndData(c, NO, message, data)
}

// 服务端响应失败的返回
func FailButServer(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndData(c, ServerFail, message, data)
}

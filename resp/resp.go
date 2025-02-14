package resp

// 实际使用根据所选框架进行调整，这里使用gin

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
)

const (
	OK         = 2000
	NO         = 4000
	ServerFail = 5000
)

// ResponseWithStatusAndData 确定统一返回格式
func ResponseWithStatusAndDataGin(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// 成功的返回
func SuccessGin(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndDataGin(c, OK, message, data)
}

// 客户端请求失败的返回
func FailGin(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndDataGin(c, NO, message, data)
}

// 服务端响应失败的返回
func FailButServerGin(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndDataGin(c, ServerFail, message, data)
}

// hertz的
func ResponseWithStatusAndData(c *app.RequestContext, status int, message string, data interface{}) {
	c.JSON(consts.StatusOK, utils.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
func Success(c *app.RequestContext, message string, data interface{}) {
	ResponseWithStatusAndData(c, OK, message, data)
}
func Fail(c *app.RequestContext, message string, data interface{}) {
	ResponseWithStatusAndData(c, NO, message, data)
}
func FailButServer(c *app.RequestContext, message string, data interface{}) {
	ResponseWithStatusAndData(c, ServerFail, message, data)
}

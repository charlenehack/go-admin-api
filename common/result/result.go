package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回信息结构体
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

var res Result

// 返回成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res.Code = 200
	res.Message = "成功"
	res.Data = data
	c.JSONP(http.StatusOK, res)
}

// 返回失败
func Failed(c *gin.Context, code int, message string) {
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSONP(http.StatusOK, res)
}

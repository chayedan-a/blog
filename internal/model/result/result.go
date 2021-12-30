package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fail(c *gin.Context, code uint) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  GetErrMap()[code],
	})
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  GetErrMap()[200],
		"data": data,
	})
}

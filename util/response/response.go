package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    100,
		"message": msg,
	})
}

func OkWithData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    101,
		"message": msg,
		"data":    data,
	})
}

func OkWithToken(ctx *gin.Context, msg string, token string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    101,
		"message": msg,
		"token":   token,
	})
}

func Error(ctx *gin.Context, msg string, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": msg,
		"error":   err,
	})
}

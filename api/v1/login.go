package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录
func Login(ctx *gin.Context) {
	var data model.User
	ctx.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

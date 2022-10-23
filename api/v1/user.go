package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(ctx *gin.Context) {
	// 创建一个User容器 存放前端传过来的数据
	var data model.User
	var msg string
	// 接收前端传过来的JSON数据 并放入容器中
	_ = ctx.ShouldBindJSON(&data)
	msg, code := validator.Validate(&data)
	if code != errmsg.SUCCESS {
		ctx.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": msg,
			})
		return
	}

	// 检测用户是否存在
	code = model.CheckUser(data.Username)
	// code == 200
	if code == errmsg.SUCCESS {
		// 把data存入用户数据库
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		// 把 用户已存在的状态码 赋值给 code
		code = errmsg.ERROR_USERNAME_USED
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(ctx *gin.Context) {
	// 获取前端传过来的数据 并转换给int类型
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	// 判断数据
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(ctx *gin.Context) {
	var data model.User
	// 从前端获取路径参数、值
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	// 检查用户是否存在
	code = model.CheckUser(data.Username)
	// 如果当前用户不存在
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		// 调用 Abort 以确保这个请求的其他函数不会被调用
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {
	// 获取前端请求的路径参数 并把获取的路径参数转换为int类型
	id, _ := strconv.Atoi(ctx.Param("id"))
	// 这里可以判断下用户是否存在
	// 操作数据库 获取返回状态码结果
	code = model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

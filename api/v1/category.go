package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(ctx *gin.Context) {
	// 创建一个User容器 存放前端传过来的数据
	var data model.Category
	// 接收前端传过来的JSON数据 并放入容器中
	_ = ctx.ShouldBindJSON(&data)
	// 检测分类是否存在
	code = model.CheckCategory(data.Name)
	// code == 200
	if code == errmsg.SUCCESS {
		// 把data存入分类数据库
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		// 把 分类已存在的状态码 赋值给 code
		code = errmsg.ERROR_CATENAME_USED
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类

// GetCate 查询分类列表
func GetCate(ctx *gin.Context) {
	// 获取前端传过来的数据 并转换给int类型
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	// 判断数据
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCate 编辑分类
func EditCate(ctx *gin.Context) {
	var data model.Category
	// 从前端获取路径参数、值
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	// 检查分类是否存在
	code = model.CheckCategory(data.Name)
	// 如果当前分类不存在
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		// 调用 Abort 以确保这个请求的其他函数不会被调用
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCate 删除分类
func DeleteCate(ctx *gin.Context) {
	// 获取前端请求的路径参数 并把获取的路径参数转换为int类型
	id, _ := strconv.Atoi(ctx.Param("id"))
	// 这里可以判断下分类是否存在
	// 操作数据库 获取返回状态码结果
	code = model.DeleteCate(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类下的文章

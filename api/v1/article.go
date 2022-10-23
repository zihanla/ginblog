package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddArt 添加分类
func AddArt(ctx *gin.Context) {
	// 创建一个User容器 存放前端传过来的数据
	var data model.Article
	// 接收前端传过来的JSON数据 并放入容器中
	_ = ctx.ShouldBindJSON(&data)
	// 把数据放入数据库 并获取CreateArt 返回的状态码
	code = model.CreateArt(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfo 查询单个文章
func GetArtInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArtInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateArt 查询指定分类下的文章
func GetCateArt(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	id, _ := strconv.Atoi(ctx.Param("id"))
	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArt 查询文章列表
func GetArt(ctx *gin.Context) {
	// 获取前端传过来的数据 并转换给int类型
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	// 判断数据
	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := model.GetArt(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditArt 编辑文章
func EditArt(ctx *gin.Context) {
	var data model.Article
	// 从前端获取路径参数、值
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code = model.EditArt(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArt 删除文章
func DeleteArt(ctx *gin.Context) {
	// 获取前端请求的路径参数 并把获取的路径参数转换为int类型
	id, _ := strconv.Atoi(ctx.Param("id"))
	// 这里可以判断下分类是否存在
	// 操作数据库 获取返回状态码结果
	code = model.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

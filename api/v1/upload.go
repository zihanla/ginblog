package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpLoad 上传图片接口
func UpLoad(ctx *gin.Context) {
	// 返回文件 文件头信息
	file, fileHeader, _ := ctx.Request.FormFile("file")

	// 通过请求文件获取文件
	fileSize := fileHeader.Size

	url, code := model.UpLoadFile(file, fileSize)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}

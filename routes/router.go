package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 设置项目运行模式
	gin.SetMode(utils.AppMode)
	// 创建一个gin实例
	// 与gin.New() 不同的是 多了两个中间件
	r := gin.New()

	// 中间件使用
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	// api分组
	// auth 需要中间件鉴权
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块
		auth.POST("article/add", v1.AddArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}
	// router 公共组
	router := r.Group("api/v1")
	{
		// 登录
		router.POST("user/add", v1.AddUser)
		// 查询分类
		router.GET("category", v1.GetCate)
		// 查询用户
		router.GET("users", v1.GetUsers)
		// 查询所有文章
		router.GET("article", v1.GetArt)
		// 查询单个文章
		router.GET("article/info/:id", v1.GetArtInfo)
		// 查询指定分类所有文章
		router.GET("article/list/:id", v1.GetCateArt)
		// 登录
		router.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}

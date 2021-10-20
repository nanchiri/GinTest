/**
*@author yezhenglin
*@date 2021/10/15 16:48
 */

package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goproject1/api/v1"
	"goproject1/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		//	用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.DELETE("user/:id", v1.DeleteUser)
		router.PUT("user/:id", v1.EditUser)
		//	分类模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		router.PUT("category/:id", v1.EditCategory)

		//文章模块的路由接口
		router.DELETE("article/:id", v1.DeleteArticle)
		router.PUT("article/:id", v1.UpdateArticle)
		router.POST("article/add", v1.AddArticle)
		router.GET("article/list/:id", v1.GetCategoryAllArticle)
		router.GET("article", v1.GetArticle)
		router.GET("article/:id", v1.GetSingeArticle)
		//
	}

	//这里有默认时间 超过默认时间就会跟客户端断开连接
	r.Run(utils.HttpPort)
}

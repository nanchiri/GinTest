/**
*@author yezhenglin
*@date 2021/10/18 10:40
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"goproject1/model"
	"goproject1/utils/errmsg"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章

//查询文章列表
func GetArticle(c *gin.Context) {

}

//编辑文章
func UpdateArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{

		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询分类下所有文章

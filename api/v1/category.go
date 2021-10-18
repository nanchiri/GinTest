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

//添加用户
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//查询单个用户

//查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//如果pagesize为-1则不使用分页功能
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategory(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑分类名
func EditCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	id, _ := strconv.Atoi(c.Param("id"))
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)

	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg,
	})
}

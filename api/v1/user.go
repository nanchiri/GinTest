/**
*@author yezhenglin
*@date 2021/10/18 10:39
 */

package v1

import "C"
import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"goproject1/model"
	"goproject1/utils/errmsg"
	"net/http"
	"strconv"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//如果pagesize为-1则不使用分页功能
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	id, _ := strconv.Atoi(c.Param("id"))
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)

	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg,
	})
}

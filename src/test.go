package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jmoiron/sqlx"
	"goproject1/model"
	"goproject1/routes"
	"strings"
	"xorm.io/core"
)


type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" url:"password" xml:"password" binding:"required:"`
}
const(
	userName="root"
	password="123456"
	ip="localhost"
	port="3306"
	dbName="users"
)

type passWd struct {
	password string
}
// Xorm Struct
type Users struct {
	Username string `xorm:"VARCHAR(255)"`
	Password string `xorm:"VARCHAR(255)"`
}
type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}
var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}




func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

var engine *xorm.Engine

func initDBXorm()  {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	var err error
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	engine, err = xorm.NewEngine("mysql", path)
	if  err != nil {
		fmt.Println("创建engine失败")
		return
	}
	engine.ShowSQL(true)
	engine.SetTableMapper(core.SnakeMapper{})
	if err := engine.Ping(); err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	fmt.Println("连接数据库成功")
}

func register(c *gin.Context){

}
func main() {
	//initDBXorm()
	//router := gin.Default()
	//user:=router.Group("/user")
	//{
	//	user.GET("/login",userLogin)
	//	user.GET("/register",userRegister)
	//	user.GET("/changepassword",changePassword)
	//	user.GET("/deleteuser",deleteUsername)
	//
	//}
	////创建get请求
	//
	//router.Run("127.0.0.1:8964")
	//r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//id, err := r.LastInsertId()
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//
	//fmt.Println("insert succ:", id)
	model.InitDB()
	routes.InitRouter()

}

//func userRegister(c *gin.Context) {
//	userName := c.Request.URL.Query().Get("username")
//	passWord := c.Request.URL.Query().Get("password")
//	//查询列表
//	st2 := new(Users)
//	result,err := engine.Where("username=?", userName).Get(st2)
//	fmt.Println("查询结果为", result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if userName != st2.Username {
//		// 无此用户
//		st2.Username = userName
//		st2.Password = passWord
//		affected, err := engine.Insert(st2)
//		if err != nil{
//			fmt.Println(err)
//		}
//		if affected != 1 {
//			c.JSON(200,gin.H{
//				"success":false,
//			})
//		} else {
//			c.JSON(200,gin.H{
//				"success":true,
//				"username":userName,
//				"msg":"Register success",
//			})
//		}
//	} else {
//		// 已存在用户，注册失败
//		fmt.Println("Already has one exsit account!")
//		c.JSON(200,gin.H{
//			"code":400,
//			"success":false,
//			"msg":"用户名已被注册",
//		})
//	}
//}
//
//func deleteUsername(c *gin.Context) {
//	userName := c.Request.URL.Query().Get("username")
//	passWord := c.Request.URL.Query().Get("password")
//	//查询列表
//	st2 := new(Users)
//	result,err := engine.Where("username=?", userName).Get(st2)
//	fmt.Println("查询结果为", result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if userName != st2.Username {
//		// 无此用户
//		c.JSON(200,gin.H{
//			"success":false,
//			"code":400,
//			"msg":"无此用户",
//		})
//	} else {
//		// 密码是否匹配
//		if passWord != st2.Password{
//			fmt.Println("password error")
//			c.JSON(200,gin.H{
//				"success":false,
//				"code":400,
//				"msg":"密码错误",
//			})
//		} else {
//			//删除账号
//			mm, err := engine.Where("username=?", userName).Delete(st2)
//			if err != nil{
//				fmt.Println(err)
//				return
//			}
//			fmt.Println(mm)
//			fmt.Println("delete account success")
//			c.JSON(200,gin.H{
//				"success":true,
//				"code":200,
//				"msg":"删除成功",
//			})
//		}
//	}
//}
//
//func changePassword(c *gin.Context) {
//	userName := c.Request.URL.Query().Get("username")
//	passWord := c.Request.URL.Query().Get("password")
//	newPassWord := c.Request.URL.Query().Get("newpassword")
//	//查询列表
//	st2 := new(Users)
//	result,err := engine.Where("username=?", userName).Get(st2)
//	fmt.Println("查询结果为", result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if userName != st2.Username {
//		// 无此用户
//		c.JSON(200,gin.H{
//			"success":false,
//			"code":400,
//			"msg":"无此用户",
//		})
//	} else {
//		// 密码是否匹配
//		if passWord != st2.Password{
//			fmt.Println("password error")
//			c.JSON(200,gin.H{
//				"success":false,
//				"code":400,
//				"msg":"密码错误",
//			})
//		} else {
//			//修改密码
//			mm, err := engine.Exec("update users set password = ? where username = ?", newPassWord, userName)
//			if err != nil{
//				fmt.Println(err)
//				return
//			}
//			fmt.Println(mm)
//			fmt.Println("change password success")
//			c.JSON(200,gin.H{
//				"success":true,
//				"code":200,
//				"msg":"修改成功",
//			})
//		}
//	}
//
//}
//
//func userLogin(c *gin.Context) {
//	userName := c.Request.URL.Query().Get("username")
//	passWord := c.Request.URL.Query().Get("password")
//	//查询列表
//	st2 := new(Users)
//	result,err := engine.Where("username=?", userName).Get(st2)
//	fmt.Println("查询结果为", result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if userName != st2.Username {
//		// 无此用户
//		c.JSON(200,gin.H{
//			"success":false,
//			"code":400,
//			"msg":"无此用户",
//		})
//	} else {
//		// 密码是否匹配
//		if passWord != st2.Password{
//			c.JSON(200,gin.H{
//				"success":false,
//				"code":400,
//				"msg":"密码错误",
//			})
//		} else {
//			c.JSON(200,gin.H{
//				"success":true,
//				"code":200,
//				"msg":"登录成功",
//			})
//		}
//	}


//


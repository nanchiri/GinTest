/**
*@author yezhenglin
*@date 2021/10/15 16:14
 */



package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goproject1/utils"
	"time"
)


var db *gorm.DB
var err error

func InitDB()  {
	db,err=gorm.Open(utils.Db,fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err != nil {
		fmt.Println(err)
		fmt.Println("连接数据库失败")
	}
	//禁止表命名复数形式
	db.SingularTable(true)
	//自动迁移
	db.AutoMigrate(&User{},&Article{},&Category{})
	//设置连接池的最大限制连接数
	db.DB().SetMaxIdleConns(10)
	//设置数据库最大连接数量
	db.DB().SetMaxOpenConns(100)
	//设置连接最大可复用时间 不能超过gin框架连接的时间
	db.DB().SetConnMaxLifetime(10*time.Second)
	//defer db.Close()
	//db.Close()
}
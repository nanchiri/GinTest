/**
*@author yezhenglin
*@date 2021/10/15 16:13
 */

package model

import (
	"github.com/jinzhu/gorm"
	"goproject1/utils/errmsg"
)

type Article struct {
	gorm.Model
	Category Category //关联分类
	Title    string   `gorm:"type:varchar(100);notnull" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`  //分类id
	Desc     string   `gorm:"type:varchar(200)" json:"desc"` //描述
	Content  string   `gorm:"type:longtext" json:"content"`  //文章类
	Img      string   `gorm:"type:varchar(100)" json:"img"`  //文章图片

}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类下所有文章

//查询单个文章

//查询文章列表

//编辑文章
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArticle 	删除文章
func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

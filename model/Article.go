/**
*@author yezhenglin
*@date 2021/10/15 16:13
 */

package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Category Category     //关联分类
	Title string `gorm:"type:varchar(100);notnull" json:"title"`
	Cid int        `gorm:"type:int;not null" json:"cid"`        //分类id
	Desc string    `gorm:"type:varchar(200)" json:"desc"`        //描述
	Content string `gorm:"type:longtext" json:"content"`       //文章类
	Img string     `gorm:"type:varchar(100)" json:"img"`       //文章图片

}

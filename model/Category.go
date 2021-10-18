/**
*@author yezhenglin
*@date 2021/10/15 16:13
分类模型
 */

package model

import (
	"github.com/jinzhu/gorm"
	"goproject1/utils/errmsg"
)

type Category struct {
 	gorm.Model
 	Name string `gorm:"type:varchar(40);not null" json:"name"`

}
// CheckUser 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	//First查询第一个参数
	db.Select("id").Where("username = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS

}

// GetCategorys 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑分类信息
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id=?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//查询分类下所有文章
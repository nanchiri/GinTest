/**
*@author yezhenglin
*@date 2021/10/15 16:12
用户模型
*/

package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"goproject1/utils/errmsg"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(40);not null" json:"username"`
	Password string `gorm:"type:varchar(40);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` //角色 1是管理员 0是一般阅读者

}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	//First查询第一个参数
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户应该存指针
func CreateUser(data *User) int {
	//data.Password = ScryptPassWord(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS

}

// GetUsers 查询用户列表 返回User切片
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(user).Where("id=?", id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (u *User) BeforeSave() {
	u.Password = ScryptPassWord(u.Password)

}

// ScryptPassWord 密码加密
func ScryptPassWord(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 44, 56, 23, 44, 22, 45, 66}
	//key(password,salt []byte,N,r,p,keyLen int)([]byte,error)
	//HashPw 哈希切片
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	FinalPassWord := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPassWord
}

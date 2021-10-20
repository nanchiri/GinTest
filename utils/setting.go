/**
*@author yezhenglin
*@date 2021/10/15 16:27
 */
//变量处理

package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadServer(file)
	LoadData(file)

}
func LoadServer(file *ini.File) {
	//取不到值就用默认值
	AppMode = file.Section("Server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString(":3000")
}
func LoadData(file *ini.File) {
	Db = file.Section("Database").Key("Db").MustString("mysql")
	DbHost = file.Section("Database").Key("DbHost").MustString("itomitsuha.xyz")
	DbPort = file.Section("Database").Key("DbPort").MustString("3306")
	DbUser = file.Section("Database").Key("DbUser").MustString("root")
	DbPassword = file.Section("Database").Key("DbPassword").MustString("Nanchiri114514!")
	DbName = file.Section("Database").Key("DbName").MustString("ginblog")
}

/**
*@author yezhenglin
*@date 2021/10/18 10:26
 */

package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code 以1000...开头 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WORNG = 1007
	//code 以2000...开头 文章模块的错误
	ERROR_CATENAME_USED=2001
	//code 以3000...开头 分类模块的错误

)

//声明一个map string是抛出的错误信息
var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户已被使用",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "Token不存在",
	ERROR_TOKEN_RUNTIME:    "token已过期",
	ERROR_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WORNG: "TOKEN格式错误",
	ERROR_CATENAME_USED: "该分类已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

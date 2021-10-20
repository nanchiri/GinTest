/**
*@author yezhenglin
*@date 2021/10/20 15:33
 */

package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goproject1/utils"
	"goproject1/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func SetToken(username string, password string) (string, int) {
	//设置10小时有效期
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			//需要转换成时间戳
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	//第一个参数是签发的方法 第二个是设置好的参数
	reClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
	token, err := reClaim.SignedString(jwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证token ****
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if key, code := setToken.Claims.(*MyClaims); code && setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		//如果 header头为空 说明token不存在
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST

		}
		checkToken := strings.SplitN(tokenHeader, "", 2)
		//固定写法
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WORNG
			c.Abort()
		}
		key, checkcode := CheckToken(checkToken[1])
		if checkcode == errmsg.ERROR {

			code = errmsg.ERROR_TOKEN_TYPE_WORNG
			c.Abort()
		}
		//	判断是否token时间戳是否过期
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Set("username", key.Username)
		c.Next()

	}
}

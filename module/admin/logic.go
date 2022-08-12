/**
* @Author:Tristan
* @Date: 2022/08/11 2:49 下午
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/common"
	"github.com/shanlongpan/catgin/consts"
)

type UserInfo struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(ctx *gin.Context) {
	var userInfo UserInfo
	r := common.NewResult(ctx)
	if ctx.ShouldBindJSON(&userInfo) != nil {
		r.Error(consts.RtnError, "参数错误")
		return
	}
	if userInfo.Username != "admin" || userInfo.Password != "123456" {
		r.Error(consts.RtnError, "用户名或密码错误")
		return
	}

	r.Success(gin.H{
		"user":  "admin",
		"token": "b2nt7cBKPTWVGEPi",
	})
}

func logout(c *gin.Context) {
}

func refreshToken(c *gin.Context) {

}

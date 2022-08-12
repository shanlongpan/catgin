package middleware

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/common"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "b2nt7cBKPTWVGEPi" {
			r := common.NewResult(c)
			r.Error(http.StatusUnauthorized, "未登录，请登录")
			c.Abort()
			return
		}
		c.Next()
	}
}

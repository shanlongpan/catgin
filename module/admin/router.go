package admin

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/admin")
	//v1.Use(AuthRequired())
	v1.POST("/login", login)
	v1.POST("/refresh_token", refreshToken)
	v1.DELETE("/logout", logout)
}

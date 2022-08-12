package helloworld

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/helloworld")
	//v1.Use(middleware.CheckLogin())
	v1.GET("/call", call)
	v1.GET("/insert", insert)
	v1.GET("/get", get)

}

package helloworld

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/helloworld")
	v1.GET("/call", call)

}

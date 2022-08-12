package k8sresult

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/k8s")
	v1.GET("/get-pod", get)

}

package k8sresult

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/k8s")
	v1.GET("/get-pod", get)

}

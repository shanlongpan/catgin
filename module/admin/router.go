package admin

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/admin")
//v1.Use(AuthRequired())
	v1.GET("/admin", loginEndpoint)
	v1.GET("/submit", submitEndpoint)
	v1.GET("/read", readEndpoint)
}
func AuthRequired(c *gin.Context)  {
	return
}
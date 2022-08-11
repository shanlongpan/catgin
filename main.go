/**
* @Author:Tristan
* @Date: 2021/9/26 8:45 下午
 */

package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/shanlongpan/catgin/init"
	"github.com/shanlongpan/catgin/middleware"
	"github.com/shanlongpan/catgin/module/admin"
	"github.com/shanlongpan/catgin/module/helloworld"
	"github.com/shanlongpan/catgin/module/k8sresult"
	"github.com/shanlongpan/catgin/xlog"
	"log"
)

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 受信任的代理地址
	//err := r.SetTrustedProxies([]string{"192.168.1.2"})
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := middleware.JwtMid.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", middleware.JwtMid.LoginHandler)

	// 未知路由
	r.NoRoute(middleware.JwtMid.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", middleware.JwtMid.RefreshHandler)

	auth.Use(middleware.JwtMid.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	// 记录请求时间，响应时间
	r.Use(xlog.LoggerToFile())
	// 添加 trace_id
	r.Use(middleware.TrackingId())

	// 中间件 recover
	r.Use(middleware.Recover)

	admin.Router(r)
	helloworld.Router(r)
	k8sresult.Router(r)

	err := r.Run(":8090") // listen and serve on 0.0.0.0:8090 (for windows "localhost:8090")
	if err != nil {
		log.Fatalln(err)
	}
}

/**
* @Author:Tristan
* @Date: 2021/9/26 8:45 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shanlongpan/catgin/init"
	"github.com/shanlongpan/catgin/middleware"
	"github.com/shanlongpan/catgin/module/admin"
	"github.com/shanlongpan/catgin/module/helloworld"
	"github.com/shanlongpan/catgin/module/k8sresult"
	"github.com/shanlongpan/catgin/xlog"
	"log"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 可以设置受信任的代理地址
	//err := r.SetTrustedProxies([]string{"192.168.1.2"})
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// 记录请求时间，响应时间
	r.Use(xlog.LoggerToFile())
	// 添加 trace_id
	r.Use(middleware.TrackingId())

	// 中间件 recover
	r.Use(middleware.Recover)

	admin.Router(r)

	// 中间件校验登录，在这个方法之前的路由不会被校验登录
	r.Use(middleware.CheckLogin())
	helloworld.Router(r)
	k8sresult.Router(r)

	err := r.Run(":8090") // listen and serve on 0.0.0.0:8090 (for windows "localhost:8090")
	if err != nil {
		log.Fatalln(err)
	}
}

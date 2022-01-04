/**
* @Author:Tristan
* @Date: 2021/9/26 8:45 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/middleware"
	"github.com/shanlongpan/catgin/module/admin"
	"github.com/shanlongpan/catgin/module/helloworld"
	"github.com/shanlongpan/catgin/xlog"
	"log"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 受信任的代理地址
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
	helloworld.Router(r)
	err := r.Run(":8090") // listen and serve on 0.0.0.0:8090 (for windows "localhost:8090")
	if err != nil {
		log.Fatalln(err)
	}
}

package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func loginEndpoint(c *gin.Context) {
	//xlog.Logger.Infoln("admin")
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)
	time.Sleep(time.Duration(d) * time.Millisecond)
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop.admin",
	})
}

func submitEndpoint(c *gin.Context) {
	//xlog.Logger.Infoln("submit")
	d, _ := c.Params.Get("name")
	fmt.Println(d[1])
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop.submit",
	})
}
func readEndpoint(c *gin.Context) {
	//xlog.Logger.Infoln("read")
	go func() {
		time.Sleep(10 * time.Second)
		panic(1)
	}()
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop.read",
	})
}

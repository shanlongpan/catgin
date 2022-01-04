package helloworld

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/consts"
	"github.com/shanlongpan/catgin/lib"
	"github.com/shanlongpan/catgin/xlog"
	"github.com/shanlongpan/micro-v3-pub/MicroV3Adapter"
	"github.com/shanlongpan/micro-v3-pub/idl/grpc/microv3"
	"math/rand"
	"net/http"
	"time"
)

func call(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)
	uid := 122121

	ctx.Set(consts.BalancingHashKey, fmt.Sprintf("%d", uid))
	res, err := MicroV3Adapter.Call(lib.GetNewCtx(ctx), &microv3.CallRequest{
		Name: fmt.Sprintf("小明%d", d),
	})

	if err != nil {
		xlog.Errorln(ctx, err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res.Msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res.Msg,
		})
	}
}

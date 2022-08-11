package helloworld

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/consts"
	"github.com/shanlongpan/catgin/model/dao"
	"github.com/shanlongpan/catgin/xlog"
	"math/rand"
	"time"
)

var p dao.Product

func call(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)
	uid := 122121

	ctx.Set(consts.BalancingHashKey, fmt.Sprintf("%d", uid))
	ctx.JSON(200, d)
}

func insert(ctx *gin.Context) {
	p.Migrate(ctx)
	var insertData []*dao.Product
	insertData = append(insertData, &dao.Product{
		Code:  "1212",
		Price: 12,
	})
	p.Insert(ctx, insertData)
}

func get(ctx *gin.Context) {
	err, result := p.Select(ctx, nil)
	if err != nil {
		xlog.Errorln(ctx, err.Error())
	} else {
		ctx.JSON(200, result)
	}
}

package helloworld

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/common"
	"github.com/shanlongpan/catgin/consts"
	"github.com/shanlongpan/catgin/model/dao"
	"math/rand"
	"net/http"
	"time"
)

var p dao.Product

func call(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(http.StatusOK)
	uid := 122121

	ctx.Set(consts.BalancingHashKey, fmt.Sprintf("%d", uid))
	r := common.NewResult(ctx)
	r.Success(d)
}

func insert(ctx *gin.Context) {
	p.Migrate(ctx)
	var insertData []*dao.Product
	insertData = append(insertData, &dao.Product{
		Code:  "1212",
		Price: 12,
	})
	p.Insert(ctx, insertData)
	r := common.NewResult(ctx)
	r.Success("ok")
}

func get(ctx *gin.Context) {
	err, result := p.Select(ctx, nil)
	if err != nil {
		r := common.NewResult(ctx)
		r.Error(consts.RtnError, err.Error())
	} else {
		r := common.NewResult(ctx)
		r.Success(result)
	}
}

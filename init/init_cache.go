package init

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"github.com/patrickmn/go-cache"
	"github.com/shanlongpan/catgin/config"
	"time"
)

func init() {
	config.CacheLocal = cache.New(5*time.Minute, 10*time.Minute)
}

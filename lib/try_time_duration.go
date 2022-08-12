package lib

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"math"
	"time"
)

func Do(attempts int) time.Duration {
	if attempts > 13 {
		return 2 * time.Minute
	}
	return time.Duration(math.Pow(float64(attempts), math.E)) * time.Millisecond * 100
}

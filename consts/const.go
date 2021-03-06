/**
* @Author:Tristan
* @Date: 2021/12/1 2:49 下午
 */

package consts

import "github.com/shanlongpan/catgin/config"

var LogFileDir = config.Conf.LogDir

const DefaultTraceIdHeader = "trace_id"
const BalancingHashKey = "hash_key"
const (
	StderrPanicLogFile = "go_Stderr_panic"
	Suffix             = ".log"
)

package exceptions

import (
	"fmt"
	"github.com/spf13/cast"
	"wyu/configs"
	"wyu/modules"
)

func init() {
	txt = modules.UtilsMergeToMap(txt, configs.YuTxT)
}

/**
 * L: Log TxT
 */
var (
	txt map[string]interface{} = map[string]interface{}{
		/**
		 * Todo: Log Text
		 */
		"l^aa": "无错误提示",
		"l^ab": "请求地址",
		"l^ac": "请求方式",
		"l^ad": "请求头部",
		"l^ae": "请求状态",
		"l^af": "请求耗时",
		"l^ag": "错误内容",
	}
)

func TxT(Tag string, customize ...interface{}) (str string) {
	s, ok := txt[Tag]
	if ok {
		str = cast.ToString(s)
	}

	if len(customize) > 0 {
		str = str + ", " + fmt.Sprint(customize ...)
	}

	return
}

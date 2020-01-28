package exceptions

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"wyu/configs"
	"wyu/modules"
)

func init() {
	msg = modules.UtilsMergeToMap(msg, configs.YuErr)
}

/**
 * m^: Models Error
 * w^: Middleware Error
 */
var (
	msg map[string]interface{} = map[string]interface{}{
		"m^aa": "条件字段或条件参数错误",
		"m^ab": "请先在「dbInitialized」结构体中设置查询类型",
		"m^ac": "数据表自增字段错误",
		"m^ad": "数据表不存在",
	}
)

func Err(Tag string, customize ...interface{}) (err error) {
	var str string = "Unknown Error!"

	s, ok := msg[Tag]
	if ok {
		str = cast.ToString(s)
	}

	if len(customize) > 0 {
		str = str + ", " + fmt.Sprint(customize ...)
	}

	err = errors.New(str)
	return
}



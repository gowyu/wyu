package configs

const(
	YuSuffix string = "html"
	YuKey string = "wYuVersion"
)

var (
	Yu map[string]interface{}

	/**
	 * Todo: Initialized Routes in routes:route
	 */
	YuRoutes map[string]string = map[string]string{}

	/**
	 * Todo: Config Subscribe & Publish in console:subscribe
	 */
	YuSubscribe []interface{} = []interface{}{
		"service",
	}

	/**
	 * Todo: Error Message customized in exceptions:errors
	 */
	YuErr map[string]interface{} = map[string]interface{}{
		"s^aa": "测试自定义错误提示",
	}
)
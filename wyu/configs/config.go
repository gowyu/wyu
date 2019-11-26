package configs

const(
	YuSuffix string = "html"
	YuKey string = "wYuVersion"
)

var (
	Yu map[string]interface{}
	YuRoutes map[string]string = map[string]string{}
	YuSubscribe []string = []string{
		"test", // Todo: do something
	}
)
package configs
/**
 * Todo: System Struct
 */


/**
 * Todo: Controller Post Construct
**/


/* ==================================================================================================== */

/**
 * Todo: Models & Basic Construct
**/

type MdbInitialized struct {
	Types string `json:"types"`// Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
	Table string `json:"table"`
	Field string `json:"field"`
	Joins [][]interface{} `json:"joins"`
	Limit int `json:"limit"`
	Start []int `json:"start"`
	Query interface{} `json:"query"`
	QueryArgs []interface{} `json:"query_args"`
	Columns []string `json:"columns"`
	OrderType string `json:"order_type"`
	OrderArgs []string `json:"order_args"`
}

type JsonMsg struct {
	Code int `json:"code"`
	Status bool `json:"status"`
	Feedback interface{} `json:"feedback"`
	Callback interface{} `json:"callback"`
}

type Logs struct {
	UserId int
	Url string
}

/* ==================================================================================================== */

//type Y struct {
//	Trees []interface{}
//	HeadsTitle string
//	BreadCrumb bool
//	DataTables bool
//	Navigation []string
//	Name string
//	Mark []string
//}
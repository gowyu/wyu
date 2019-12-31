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
	Types string // Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
	Joins [][]interface{}
	Query interface{}
	QueryArgs []interface{}
	Columns []string
}

type Y struct {
	Trees []interface{}
	HeadsTitle string
	BreadCrumb bool
	DataTables bool
	Navigation []string
	Name string
	Mark []string
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

/**
 * Todo: .Yaml files
 */
//type Tpls struct {
//	Status bool `json:"status"`
//	Resources string `json:"resources"`
//	Dir string `json:"dir"`
//	DirViews string `json:"dir_views"`
//	DirLayout string `json:"dir_layout"`
//	DirShared string `json:"dir_shared"`
//	Suffix string `json:"suffix"`
//	StaticStatus bool `json:"static_status"`
//	Static string `json:"static"`
//	StaticIcon string `json:"static_icon"`
//}

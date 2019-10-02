package configs

/**
 * Controller Post Struct
**/


// ==================================================================================

/**
 * Models & Basic Struct
**/
type MdbInitialized struct {
	Query		interface{}
	QueryArgs 	[]interface{}
	Columns 	[]string
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

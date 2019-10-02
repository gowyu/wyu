package configs

const(
	CookieLoginEmail		string = "Email"
	CookieLoginUserId		string = "UserID"
	CookieLoginUserRoleId	string = "UserRoleID"

	HttpSuffix			string =  ".g"

	WYuKey		string = "WYuControlVersion"
	WYuKeyTree	string = "WYuTrees"
	WYuKeyRole	string = "WYuRoles"
)

var (
	WYu map[string]string = map[string]string{
		WYuKeyTree:"",
		WYuKeyRole:"",
	}

	WYuTrees map[int]interface{}
	WYuRoles map[int]interface{}

	WYuRouteHttp map[string]string
)

var ArrLoginCookie []string = []string{
		CookieLoginUserId,
		CookieLoginEmail,
		CookieLoginUserRoleId,
	}
package errmsg

// 定义状态码
const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1xxx 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code = 2xxx 文章模块的错误
	ERROR_ART_NOT_EXIST = 2001
	// code = 3xxx 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

// 给状态码指定错误信息
var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ART_NOT_EXIST: "该文章不存在",

	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

// GetErrMsg 根据状态码获取状态信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}

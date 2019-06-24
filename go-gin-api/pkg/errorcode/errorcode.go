package errorcode

import "fmt"

// 错误代码.
type ErrorCode struct {
	Code int
	Msg  string
	Desc string
}

// 添加错误详情信息
func (e *ErrorCode) AddDetail(format string, param ...interface{}) {
	e.Desc = e.Desc + " " + fmt.Sprintf(format, param)
}

// 返回错误时的HTTP状态码
var errorCodeToHttpStatusCode = map[string]int{}

func (e *ErrorCode) GetHttpStatusCode() int {
	if code, OK := errorCodeToHttpStatusCode[e.Msg]; OK {
		return code
	}
	return 0
}

func (e *ErrorCode) SetHttpStatusCode(code int) {
	errorCodeToHttpStatusCode[e.Msg] = code
}

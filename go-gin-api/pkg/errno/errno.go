package errno

import "net/http"

// 错误代码
type Errno struct {
	No   int
	Code string
}

// 错误码表
const (
	// 请求成功.
	OK = 0

	// 请求无效.
	INVALID = http.StatusBadRequest // 400

	// 缺少参数.
	MissingParameter = http.StatusBadRequest // 400
	// 参数无效.
	InvalidParameter = http.StatusBadRequest // 400

	// 请求未认证.
	UNAUTHORIZED = http.StatusUnauthorized // 401

	// 无权限执行该操作.
	Forbidden = http.StatusForbidden // 403
	// 账号未开通相应服务.
	OperationDenied = http.StatusForbidden // 403
	// 账号已欠费，请充值。
	OperationDeniedSuspended = http.StatusForbidden // 403

	// 请求发生错误.
	ERROR = http.StatusInternalServerError // 500

	// 后台发生未知错误，请稍后重试或联系客服解决。
	InternalError = http.StatusInternalServerError // 500
	// 服务不可用。
	ServiceUnAvailable = http.StatusServiceUnavailable // 503
)

// 错误码表详情
var Texts = map[*Errno]string{
	&Errno{OK, "OK"}:                                             "OK",
	&Errno{INVALID, "INVALID"}:                                   "请求无效",
	&Errno{MissingParameter, "MissingParameter"}:                 "缺少参数",
	&Errno{InvalidParameter, "InvalidParameter"}:                 "参数无效",
	&Errno{UNAUTHORIZED, "UNAUTHORIZED"}:                         "请求未认证",
	&Errno{Forbidden, "Forbidden"}:                               "无权限执行该操作",
	&Errno{OperationDenied, "OperationDenied"}:                   "账号未开通相应服务",
	&Errno{OperationDeniedSuspended, "OperationDeniedSuspended"}: "账号已欠费，请充值。",
	&Errno{ERROR, "ERROR"}:                                       "请求发生错误",
	&Errno{InternalError, "InternalError"}:                       "后台发生未知错误，请稍后重试或联系客服解决。",
	&Errno{ServiceUnAvailable, "ServiceUnAvailable"}:             "服务不可用。",
}

// 获取错误信息
func (e *Errno) Error() string {
	return Texts[e]
}

package errorcode

import "net/http"

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

func Common(code int) ErrorCode {
	e := ErrorCode{Code: code}

	switch code {
	case OK:
		e.Msg, e.Desc = "OK", "请求成功"
		e.SetHttpStatusCode(http.StatusOK)
		break
	case INVALID:
		e.Msg, e.Desc = "INVALID", "请求无效"
		e.SetHttpStatusCode(http.StatusBadRequest)
		break
	}

	return e
}

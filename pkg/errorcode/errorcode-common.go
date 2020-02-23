package errorcode

import "net/http"

// 错误码表.
const (
	// 请求成功.
	OK = 0

	// 请求无效.
	INVALID = 400

	// 缺少参数.
	MissingParameter = 100401
	// 参数无效.
	InvalidParameter = 100402

	// 请求未认证通过.
	UNAUTHORIZED = 401

	// 无权限执行该操作.
	Forbidden = 403
	// 账号未开通相应服务.
	OperationDenied = 100403
	// 账号已欠费，请充值.
	OperationDeniedSuspended = 100404

	// 请求发生错误.
	ERROR = 500

	// 后台发生未知错误，请稍后重试或联系客服解决.
	InternalError = 100500
	// 服务不可用.
	ServiceUnAvailable = 503
)

// Common 生成错误信息.
func Common(code int) ErrorCode {
	e := ErrorCode{Code: code}

	switch code {
	case OK:
		e.Msg, e.Desc = "OK", "请求成功."
		e.SetHttpStatusCode(http.StatusOK)
		break
	case INVALID:
		e.Msg, e.Desc = "INVALID", "请求无效."
		e.SetHttpStatusCode(http.StatusBadRequest)
		break
	case MissingParameter:
		e.Msg, e.Desc = "MissingParameter", "缺少参数."
		e.SetHttpStatusCode(http.StatusBadRequest)
		break
	case InvalidParameter:
		e.Msg, e.Desc = "InvalidParameter", "参数无效."
		e.SetHttpStatusCode(http.StatusBadRequest)
		break
	case UNAUTHORIZED:
		e.Msg, e.Desc = "UNAUTHORIZED", "请求未认证通过."
		e.SetHttpStatusCode(http.StatusUnauthorized)
		break
	case Forbidden:
		e.Msg, e.Desc = "Forbidden", "无权限执行该操作."
		e.SetHttpStatusCode(http.StatusForbidden)
		break
	case OperationDenied:
		e.Msg, e.Desc = "OperationDenied", "账号未开通相应服务."
		e.SetHttpStatusCode(http.StatusForbidden)
		break
	case OperationDeniedSuspended:
		e.Msg, e.Desc = "OperationDeniedSuspended", "账号已欠费，请充值."
		e.SetHttpStatusCode(http.StatusForbidden)
		break
	case ERROR:
		e.Msg, e.Desc = "ERROR", "请求发生错误."
		e.SetHttpStatusCode(http.StatusInternalServerError)
		break
	case InternalError:
		e.Msg, e.Desc = "InternalError", "后台发生未知错误，请稍后重试或联系客服解决."
		e.SetHttpStatusCode(http.StatusInternalServerError)
		break
	case ServiceUnAvailable:
		e.Msg, e.Desc = "ServiceUnAvailable", "服务不可用."
		e.SetHttpStatusCode(http.StatusServiceUnavailable)
		break
	}

	return e
}

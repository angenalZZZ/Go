package model

import "github.com/angenalZZZ/gofunc/http/errorcode"

type Response struct {
	errorcode.ErrorCode
	Data interface{} `json:"data"`
}

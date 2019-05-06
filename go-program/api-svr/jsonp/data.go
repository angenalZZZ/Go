package jsonp

import (
	"encoding/json"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// 数据输出
type Data map[string]interface{}

func Success(d Data) Data {
	var b = Data{
		"code": 0,
		"msg":  "success",
	}
	for k, v := range d {
		b[k] = v
	}

	return b
}

func Error(e interface{}) Data {
	var msg = ""
	if e != nil {
		if h, ok := e.(error); ok {
			msg = h.Error()
		} else if h, ok := e.(fmt.Stringer); ok {
			msg = h.String()
		} else {
			msg = fmt.Sprint(e)
		}
	}

	return Data{
		"code": 1,
		"msg":  msg,
	}
}

// response ok
func (d Data) OK(w http.ResponseWriter, r *http.Request) {
	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//set http status code
	w.WriteHeader(http.StatusOK)
	//output data
	json.NewEncoder(w).Encode(d)
}

// response error
func (d Data) Error(w http.ResponseWriter, r *http.Request) {
	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//set http status code 202
	w.WriteHeader(http.StatusAccepted)
	//output data
	json.NewEncoder(w).Encode(d)
}

// 数据转换
func Marshal(v interface{}) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
}

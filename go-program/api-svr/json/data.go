package json

import (
	"encoding/json"
	"net/http"
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

func Error(e error) Data {
	var msg = "error"
	if e != nil {
		msg = e.Error()
	}

	return Data{
		"code": 1,
		"msg":  msg,
	}
}

// response ok
func (d Data) ResponseSuccess(w http.ResponseWriter, r *http.Request) {
	//set json response
	setHeader(w)
	w.WriteHeader(http.StatusOK)
	//output data
	json.NewEncoder(w).Encode(d)
}

// response error
func (d Data) ResponseError(w http.ResponseWriter, r *http.Request) {
	//set json response
	setHeader(w)
	w.WriteHeader(http.StatusAccepted)
	//output data
	json.NewEncoder(w).Encode(d)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

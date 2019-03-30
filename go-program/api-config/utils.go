package api_config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 加载输入参数|文件路径
func LoadArgInput(s string) ([]byte, error) {
	if s == "" {
		return nil, fmt.Errorf("no input")
	} else if s == "+" {
		return []byte("{}"), nil
	}
	var r io.Reader
	if s == "-" {
		r = os.Stdin
	} else {
		if f, e := os.Open(s); e != nil {
			return nil, e
		} else {
			defer f.Close() // end
			r = f
		}
	}
	return ioutil.ReadAll(r)
}

// "+" > {}
func JsonParse(s string) (v interface{}, e error) {
	var data []byte
	if data, e = LoadArgInput(s); e == nil {
		e = json.Unmarshal(data, v)
	}
	return
}

// {} > "+"
func JsonStringify(v interface{}, indent bool) (s []byte, e error) {
	if indent == false {
		s, e = json.MarshalIndent(v, "", "    ")
	} else {
		s, e = json.Marshal(v)
	}
	return
}

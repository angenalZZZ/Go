package cors

import (
	"log"
	"net/http"
)

// cors request
func Cors(w http.ResponseWriter, r *http.Request, method []string) (ok bool) {
	ok = true

	// 跟踪请求
	log.Printf(" http %s %s\n", r.Method, r.URL)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	for _, v := range method {
		if v == r.Method {
			ok = false
		}
	}
	if ok == true {
		w.WriteHeader(http.StatusForbidden)
	}
	return
}

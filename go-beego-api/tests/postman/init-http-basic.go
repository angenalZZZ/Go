package postman

import (
	"io"
	"net/http"
	"time"
)

//const url = "https://postman-echo.com"

func httpBasicCurrentUTCTime(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, time.Now().String())
}

func InitHttpBasic() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/time/now", httpBasicCurrentUTCTime)
	return mux
}

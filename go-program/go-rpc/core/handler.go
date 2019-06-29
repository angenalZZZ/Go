package core

import "errors"

// HandlerName provider the name of the only
// method that `core` exposes via the RPC
// interface.
//
// This could be replaced by the use of the reflect
// package (e.g, `reflect.ValueOf(func).Pointer()).Name()`).
const HandlerName = "Handler.Execute"

type Handler struct {
	Actions map[string]func(req *Request, res *Response) (err error)
}

// Execute is the exported method that a RPC client can
// make use of by calling the RPC server using `HandlerName`
// as the endpoint.
//
// It takes a Request and produces a Response if no error
func (h *Handler) Execute(req *Request, res *Response) (err error) {
	if req.Action == "" {
		err = errors.New("A action must be specified")
		return
	}

	if handler, OK := h.Actions[req.Action]; OK {
		err = handler(req, res)
		return
	}

	err = errors.New("Action[" + req.Action + "] is not found")
	return
}

// Request Action + Query
type Request struct {
	Action string
	Query  string
}

// Response Code + Result
type Response struct {
	Code   int
	Result string
}

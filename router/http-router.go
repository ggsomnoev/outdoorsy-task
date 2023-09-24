package router

import (
	"log/slog"
	"net/http"
)

type HttpRouter interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type httpRouter struct{}

func NewHttpRouter() HttpRouter {
	return &httpRouter{}
}

func (*httpRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(uri, f)
}

func (*httpRouter) SERVE(port string) {
	slog.Info("Http server running on...", slog.String("port", port))
	http.ListenAndServe(port, nil)
}
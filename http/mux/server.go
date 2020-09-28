package mux

import (
	"fmt"
	goHttp "net/http"

	"github.com/gorilla/mux"

	"github.com/liampulles/go-driver-common/http"
)

// Listener wraps the elements of the net/http.Server we care about
type Listener interface {
	ListenAndServe() error
}

// GoServer wraps the Go HTTP Server
type GoServer struct {
	Original Listener
}

var _ http.Server = &GoServer{}

// Start implements the http.Server interface
func (g *GoServer) Start() error {
	if err := g.Original.ListenAndServe(); err != nil {
		return fmt.Errorf("go http server error: %w", err)
	}
	return nil
}

// NewServer creates a new Mux server
func NewServer(mappings []http.HandlerMapping, port int) *GoServer {
	router := registerRouter(mappings)
	return createServer(router, port)
}

func createServer(router *mux.Router, port int) *GoServer {
	return &GoServer{
		Original: &goHttp.Server{
			Handler: router,
			Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		},
	}
}

func registerRouter(mappings []http.HandlerMapping) *mux.Router {
	r := mux.NewRouter()
	for _, mapping := range mappings {
		r.HandleFunc(mapping.PathPattern, wrapHandler(mapping.Handler)).
			Methods(mapping.Method)
	}
	return r
}

func wrapHandler(handler http.Handler) func(res goHttp.ResponseWriter, req *goHttp.Request) {
	return func(res goHttp.ResponseWriter, req *goHttp.Request) {
		adapReq := GoRequest{
			Original: *req,
		}
		adapRes := handler(&adapReq)
		adaptResponse(adapRes, res)
	}
}

func adaptResponse(in *http.Response, out goHttp.ResponseWriter) {
	_, err := out.Write(in.Body)
	if err != nil {
		out.WriteHeader(500)
		return
	}
	out.WriteHeader(in.StatusCode)

	for k, v := range in.Headers {
		for _, vi := range v {
			out.Header().Add(k, vi)
		}
	}
}

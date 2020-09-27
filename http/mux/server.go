package mux

import (
	"fmt"
	goHttp "net/http"

	"github.com/gorilla/mux"
	"github.com/liampulles/go-driver-common/http"
)

// GoServer wraps the Go HTTP Server
type GoServer struct {
	original goHttp.Server
}

var _ http.Server = &GoServer{}

// Start implements the http.Server interface
func (g *GoServer) Start() error {
	return g.original.ListenAndServe()
}

// NewServer creates a new Mux server
func NewServer(mappings []http.HandlerMapping, port int) http.Server {
	router := registerRouter(mappings)
	return createServer(router, port)
}

func createServer(router *mux.Router, port int) *GoServer {
	return &GoServer{
		original: goHttp.Server{
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
		adapReq := goRequest{
			original: *req,
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
		out.Header().Add(k, v)
	}
}

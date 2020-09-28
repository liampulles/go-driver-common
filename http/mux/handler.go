package mux

import (
	"fmt"
	"io/ioutil"
	goHttp "net/http"

	"github.com/gorilla/mux"

	"github.com/liampulles/go-driver-common/http"
)

// GoRequest wraps net/http.Request
type GoRequest struct {
	Original goHttp.Request
}

// Check we implement the interface
var _ http.Request = &GoRequest{}

// Headers implements the http.Request interface
func (r *GoRequest) Headers() map[string][]string {
	return r.Original.Header
}

// Body implements the http.Request interface
func (r *GoRequest) Body() ([]byte, error) {
	body, err := ioutil.ReadAll(r.Original.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil failed: %w", err)
	}
	return body, nil
}

// PathParams implements the http.Request interface
func (r *GoRequest) PathParams() map[string]string {
	return mux.Vars(&r.Original)
}

// QueryParams implements the http.Request interface
func (r *GoRequest) QueryParams() map[string][]string {
	return r.Original.URL.Query()
}

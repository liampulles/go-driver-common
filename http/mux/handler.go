package mux

import (
	"fmt"
	"io/ioutil"
	goHttp "net/http"

	"github.com/gorilla/mux"

	"github.com/liampulles/go-driver-common/http"
)

type goRequest struct {
	original goHttp.Request
}

// Check we implement the interface
var _ http.Request = &goRequest{}

func (r *goRequest) Headers() map[string]string {
	return r.Headers()
}

func (r *goRequest) Body() ([]byte, error) {
	body, err := ioutil.ReadAll(r.original.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil failed: %w", err)
	}
	return body, nil
}

func (r *goRequest) PathParams() map[string]string {
	return mux.Vars(&r.original)
}

func (r *goRequest) QueryParams() map[string][]string {
	return r.original.URL.Query()
}

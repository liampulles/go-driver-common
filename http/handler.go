package http

// Request abstracts a HTTP request from its implementation
type Request interface {
	Headers() map[string][]string
	Body() ([]byte, error)
	PathParams() map[string]string
	QueryParams() map[string][]string
}

// Response models everything you might like to return to the
// client in response to a HTTP request.
type Response struct {
	Headers    map[string][]string
	StatusCode int
	Body       []byte
}

// Handler is a function which accepts a HTTP request, processes it,
// and returns a HTTP response.
type Handler func(Request) *Response

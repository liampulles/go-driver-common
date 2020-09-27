package http

// Server serves HTTP requests
type Server interface {
	Start() error
}

// HandlerMapping describes a routing to a handler.
// Keys should be the path, and you may template them.
// For example:
//     /api/v1/users/{id}
// will provide path param id.
type HandlerMapping struct {
	Method      string
	PathPattern string
	Handler     Handler
}

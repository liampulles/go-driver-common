package http

// Server serves HTTP requests
type Server interface {
	Start()
}

// HandlerMappings describes path mappings to handlers.
// Keys should be the path, and you may template them.
// For example:
//     /api/v1/users/{id}
// will provide path param id.
type HandlerMappings map[string]Handler
